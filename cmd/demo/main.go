package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	sharex "github.com/sharex-org/sharex-sdk-go"

	eciesgo "github.com/ecies/go/v2"
)

func main() {
	ctx := context.Background()

	// Simulate the Indexer's private key used to decrypt device uploads.
	serverWallet, err := sharex.GenerateWallet()
	if err != nil {
		log.Fatalf("generate indexer wallet: %v", err)
	}

	serverKey, err := crypto.HexToECDSA(strings.TrimPrefix(serverWallet.PrivateKeyHex, "0x"))
	if err != nil {
		log.Fatalf("parse server private key: %v", err)
	}

	server := startMockIndexer(serverKey)
	defer server.Close()

	fmt.Printf("Mock Indexer running at %s\n", server.URL)

	client, err := sharex.NewClient(sharex.Config{
		IndexerBaseURL:         server.URL,
		EncryptionPublicKeyHex: serverWallet.PublicKeyHex,
	})
	if err != nil {
		log.Fatalf("new client: %v", err)
	}

	// Generate a wallet on the device side.
	deviceWallet, err := sharex.GenerateWallet()
	if err != nil {
		log.Fatalf("generate wallet: %v", err)
	}

	fmt.Printf("Device wallet address: %s\n", deviceWallet.Address)

	keyPath, err := saveWalletKey(deviceWallet.PrivateKeyHex)
	if err != nil {
		log.Fatalf("save private key: %v", err)
	}

	fmt.Printf("Private key exported to %s (0600)\n", keyPath)

	restoredWallet, err := loadWalletFromFile(keyPath)
	if err != nil {
		log.Fatalf("reload wallet: %v", err)
	}

	if restoredWallet.Address != deviceWallet.Address {
		log.Fatalf("restored wallet mismatch: got %s want %s", restoredWallet.Address, deviceWallet.Address)
	}

	fmt.Println("Reloaded wallet from disk; public key:", restoredWallet.PublicKeyHex)

	// Register the device
	regResp, err := client.RegisterDevice(ctx, sharex.RegisterDeviceRequest{
		DeviceID:    "DEVICE-DEMO-001",
		DeviceType:  "Terminal",
		PartnerCode: "PARTNER-01",
		MerchantID:  "MERCH-01",
		WalletAddr:  deviceWallet.Address,
	})
	if err != nil {
		log.Fatalf("register device: %v", err)
	}

	fmt.Printf("RegisterDevice response: %+v\n", regResp)

	// Upload a demo batch using the simplified high-level API
	// The SDK automatically:
	// 1. Parses transaction data to derive orderCount and totalAmount
	// 2. Builds and signs the on-chain transaction
	// 3. Submits the batch to the indexer
	batchResp, err := client.UploadTransactionBatch(ctx, sharex.UploadTransactionBatchParams{
		DeviceID:            "DEVICE-DEMO-001",
		DateComparable:      time.Now().UTC().Format("20060102"),
		TransactionDataJSON: `{"transactions":[{"id":1,"factOvertimeMoney":"99.99","cdb":"tawdajbntawdajbqtnwp6jhrt2zpekxq","deviceTerminal":"DEVICE-DEMO-001"}]}`,
		PrivateKeyHex:       deviceWallet.PrivateKeyHex,
		Nonce:               0, // Replace with pending nonce from your RPC when running against a live chain.
		// ChainID & ContractAddress use SDK defaults (opBNB mainnet + production Deshare).
	})
	if err != nil {
		log.Fatalf("upload batch: %v", err)
	}

	fmt.Printf("UploadTransactionBatch response: %+v\n", batchResp)
}

func startMockIndexer(priv *ecdsa.PrivateKey) *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/sdk/register-device", func(w http.ResponseWriter, r *http.Request) {
		req, err := decryptDeviceRequest(r, priv)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("[Indexer] Registered device %s wallet=%s\n", req.DeviceID, req.WalletAddr)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(sharex.RegisterDeviceResponse{
			Success:                true,
			TransactionHash:        "0xdevicehash",
			RoleTransactionHash:    "0xrolehash",
			FundingTransactionHash: "0xfundhash",
			DeviceID:               req.DeviceID,
			WalletAddress:          req.WalletAddr,
			Message:                "device registered",
		})
	})

	mux.HandleFunc("/sdk/upload", func(w http.ResponseWriter, r *http.Request) {
		req, err := decryptBatchRequest(r, priv)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("[Indexer] Received %d signed tx(s) from %s\n", len(req.SignedTransactions), req.DeviceID)

		json.NewEncoder(w).Encode(sharex.BatchResponse{
			Success:           true,
			TransactionHashes: []string{"0xdeadbeef", "0xfee"},
			BroadcastCount:    len(req.SignedTransactions),
			Message:           "batch accepted",
		})
	})

	return httptest.NewServer(mux)
}

func decryptDeviceRequest(r *http.Request, priv *ecdsa.PrivateKey) (sharex.RegisterDeviceRequest, error) {
	var env struct {
		Payload string `json:"payload"`
	}
	if err := json.NewDecoder(r.Body).Decode(&env); err != nil {
		return sharex.RegisterDeviceRequest{}, fmt.Errorf("decode envelope: %w", err)
	}
	data, err := decrypt(env.Payload, priv)
	if err != nil {
		return sharex.RegisterDeviceRequest{}, err
	}
	var req sharex.RegisterDeviceRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return sharex.RegisterDeviceRequest{}, fmt.Errorf("decode payload: %w", err)
	}
	return req, nil
}

func decryptBatchRequest(r *http.Request, priv *ecdsa.PrivateKey) (sharex.BatchRequest, error) {
	var env struct {
		Payload string `json:"payload"`
	}
	if err := json.NewDecoder(r.Body).Decode(&env); err != nil {
		return sharex.BatchRequest{}, fmt.Errorf("decode envelope: %w", err)
	}
	data, err := decrypt(env.Payload, priv)
	if err != nil {
		return sharex.BatchRequest{}, err
	}
	var req sharex.BatchRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return sharex.BatchRequest{}, fmt.Errorf("decode payload: %w", err)
	}
	return req, nil
}

func decrypt(payload string, priv *ecdsa.PrivateKey) ([]byte, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, fmt.Errorf("decode base64: %w", err)
	}

	// Convert ECDSA private key to eciesgo private key
	privKeyBytes := crypto.FromECDSA(priv)
	eciesPrivKey := eciesgo.NewPrivateKeyFromBytes(privKeyBytes)

	plain, err := eciesgo.Decrypt(eciesPrivKey, cipherBytes)
	if err != nil {
		return nil, fmt.Errorf("decrypt payload: %w", err)
	}
	return plain, nil
}

func saveWalletKey(privateKeyHex string) (string, error) {
	file, err := os.CreateTemp("", "sharex-wallet-*.key")
	if err != nil {
		return "", fmt.Errorf("create wallet file: %w", err)
	}
	defer file.Close()
	if err := file.Chmod(0o600); err != nil {
		os.Remove(file.Name())
		return "", fmt.Errorf("chmod wallet file: %w", err)
	}
	if _, err := fmt.Fprintln(file, privateKeyHex); err != nil {
		os.Remove(file.Name())
		return "", fmt.Errorf("write wallet file: %w", err)
	}
	return file.Name(), nil
}

func loadWalletFromFile(path string) (*sharex.Wallet, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read wallet file: %w", err)
	}
	key := strings.TrimSpace(string(data))
	wallet, err := sharex.WalletFromPrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("wallet from private key: %w", err)
	}
	return wallet, nil
}
