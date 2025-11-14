package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	sharex "github.com/sharex/sharex-sdk-go"
)

func main() {
	ctx := context.Background()

	// Simulate the Indexer's private key used to decrypt device uploads.
	serverKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("generate indexer key: %v", err)
	}

	server := startMockIndexer(serverKey)
	defer server.Close()

	fmt.Printf("Mock Indexer running at %s\n", server.URL)

	client, err := sharex.NewClient(sharex.Config{
		IndexerBaseURL:         server.URL,
		EncryptionPublicKeyHex: hexutil.Encode(crypto.FromECDSAPub(&serverKey.PublicKey)),
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

	signedTxs, err := buildSignedDemoTransactions(deviceWallet.PrivateKeyHex)
	if err != nil {
		log.Fatalf("sign transactions: %v", err)
	}

	// Upload a demo batch
	batchResp, err := client.SubmitTransactionBatch(ctx, sharex.BatchRequest{
		DeviceID:           "DEVICE-DEMO-001",
		DateComparable:     time.Now().UTC().Format("20060102"),
		OrderCount:         len(signedTxs),
		TotalAmount:        "19999",
		SignedTransactions: signedTxs,
	})
	if err != nil {
		log.Fatalf("submit batch: %v", err)
	}

	fmt.Printf("SubmitTransactionBatch response: %+v\n", batchResp)
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

func buildSignedDemoTransactions(privateKeyHex string) ([]string, error) {
	chainID := big.NewInt(11155111) // Sepolia
	to := common.HexToAddress("0x0000000000000000000000000000000000000000")
	first := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     0,
		GasTipCap: big.NewInt(1_500_000_000),
		GasFeeCap: big.NewInt(2_500_000_000),
		Gas:       21_000,
		To:        &to,
		Value:     big.NewInt(0),
	})
	second := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     1,
		GasTipCap: big.NewInt(1_500_000_000),
		GasFeeCap: big.NewInt(2_500_000_000),
		Gas:       21_000,
		To:        &to,
		Value:     big.NewInt(0),
	})

	return sharex.SignTransactions(privateKeyHex, []*types.Transaction{first, second})
}

func decrypt(payload string, priv *ecdsa.PrivateKey) ([]byte, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, fmt.Errorf("decode base64: %w", err)
	}
	plain, err := ecies.ImportECDSA(priv).Decrypt(cipherBytes, nil, nil)
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
