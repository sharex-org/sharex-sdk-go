package sharex

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestNormalizeBaseURL(t *testing.T) {
	got, err := normalizeBaseURL(" https://api.sharex.network/ ")
	if err != nil {
		t.Fatalf("normalize: %v", err)
	}
	if want := "https://api.sharex.network"; got != want {
		t.Fatalf("unexpected normalized url. got %s want %s", got, want)
	}

	if _, err := normalizeBaseURL("::bad::"); err == nil {
		t.Fatal("expected error for invalid url")
	}
}

func TestRegisterDeviceEncryption(t *testing.T) {
	serverWallet, err := GenerateWallet()
	if err != nil {
		t.Fatalf("wallet: %v", err)
	}
	priv, err := crypto.HexToECDSA(strings.TrimPrefix(serverWallet.PrivateKeyHex, "0x"))
	if err != nil {
		t.Fatalf("priv: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/sdk/register-device", func(w http.ResponseWriter, r *http.Request) {
		var env secureEnvelope
		if err := json.NewDecoder(r.Body).Decode(&env); err != nil {
			t.Fatalf("decode envelope: %v", err)
		}
		payload, err := decryptForTests(priv, env.Payload)
		if err != nil {
			t.Fatalf("decrypt request: %v", err)
		}
		var req RegisterDeviceRequest
		if err := json.Unmarshal(payload, &req); err != nil {
			t.Fatalf("decode payload: %v", err)
		}
		if req.DeviceID != "DEVICE1" || req.WalletAddr == "" {
			t.Fatalf("unexpected payload: %+v", req)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(RegisterDeviceResponse{
			Success:                true,
			TransactionHash:        "0xabc",
			RoleTransactionHash:    "0xdef",
			FundingTransactionHash: "0x123",
			DeviceID:               req.DeviceID,
			WalletAddress:          req.WalletAddr,
			Message:                "ok",
		})
	})

	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	client, err := NewClient(Config{
		IndexerBaseURL:         server.URL,
		EncryptionPublicKeyHex: serverWallet.PublicKeyHex,
	})
	if err != nil {
		t.Fatalf("client: %v", err)
	}

	wallet, err := GenerateWallet()
	if err != nil {
		t.Fatalf("device wallet: %v", err)
	}

	res, err := client.RegisterDevice(context.Background(), RegisterDeviceRequest{
		DeviceID:    "DEVICE1",
		DeviceType:  "terminal",
		PartnerCode: "P",
		MerchantID:  "M",
		WalletAddr:  wallet.Address,
	})
	if err != nil {
		t.Fatalf("register device: %v", err)
	}
	if !res.Success || res.TransactionHash != "0xabc" || res.RoleTransactionHash != "0xdef" {
		t.Fatalf("unexpected response: %+v", res)
	}
}

func TestUploadTransactionBatch(t *testing.T) {
	serverWallet, err := GenerateWallet()
	if err != nil {
		t.Fatalf("generate server wallet: %v", err)
	}

	serverKey, err := crypto.HexToECDSA(strings.TrimPrefix(serverWallet.PrivateKeyHex, "0x"))
	if err != nil {
		t.Fatalf("parse server private key: %v", err)
	}

	deviceWallet, err := GenerateWallet()
	if err != nil {
		t.Fatalf("generate device wallet: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/sdk/upload", func(w http.ResponseWriter, r *http.Request) {
		var env secureEnvelope
		if err := json.NewDecoder(r.Body).Decode(&env); err != nil {
			http.Error(w, "decode envelope", http.StatusBadRequest)
			return
		}

		plaintext, err := decryptForTests(serverKey, env.Payload)
		if err != nil {
			http.Error(w, fmt.Sprintf("decrypt: %v", err), http.StatusBadRequest)
			return
		}

		var req BatchRequest
		if err := json.Unmarshal(plaintext, &req); err != nil {
			http.Error(w, "decode request", http.StatusBadRequest)
			return
		}

		// Verify the batch request fields
		if req.DeviceID != "TEST-DEVICE" {
			http.Error(w, "wrong deviceId", http.StatusBadRequest)
			return
		}
		if req.OrderCount != 1 {
			http.Error(w, "wrong orderCount", http.StatusBadRequest)
			return
		}
		if req.TotalAmount != "99.99" {
			http.Error(w, "wrong totalAmount", http.StatusBadRequest)
			return
		}
		if len(req.SignedTransactions) != 1 {
			http.Error(w, "missing signedTransactions", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(BatchResponse{
			Success:           true,
			TransactionHashes: []string{"0xabc123"},
			BroadcastCount:    1,
			Message:           "batch uploaded",
		})
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	client, err := NewClient(Config{
		IndexerBaseURL:         srv.URL,
		EncryptionPublicKeyHex: serverWallet.PublicKeyHex,
	})
	if err != nil {
		t.Fatalf("new client: %v", err)
	}

	// Use the simplified high-level API
	resp, err := client.UploadTransactionBatch(context.Background(), UploadTransactionBatchParams{
		DeviceID:            "TEST-DEVICE",
		DateComparable:      "20241128",
		TransactionDataJSON: `{"transactions":[{"factOvertimeMoney":"99.99","order_no":"TEST001"}]}`,
		PrivateKeyHex:       deviceWallet.PrivateKeyHex,
		Nonce:               0,
	})
	if err != nil {
		t.Fatalf("upload batch: %v", err)
	}

	if !resp.Success {
		t.Errorf("expected success=true, got false")
	}
	if resp.BroadcastCount != 1 {
		t.Errorf("expected broadcastCount=1, got %d", resp.BroadcastCount)
	}
}

func TestSubmitTransactionBatchEncryption(t *testing.T) {
	serverWallet, err := GenerateWallet()
	if err != nil {
		t.Fatalf("wallet: %v", err)
	}
	priv, err := crypto.HexToECDSA(strings.TrimPrefix(serverWallet.PrivateKeyHex, "0x"))
	if err != nil {
		t.Fatalf("priv: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/sdk/upload", func(w http.ResponseWriter, r *http.Request) {
		var env secureEnvelope
		if err := json.NewDecoder(r.Body).Decode(&env); err != nil {
			t.Fatalf("decode envelope: %v", err)
		}
		payload, err := decryptForTests(priv, env.Payload)
		if err != nil {
			t.Fatalf("decrypt request: %v", err)
		}
		var req BatchRequest
		if err := json.Unmarshal(payload, &req); err != nil {
			t.Fatalf("decode payload: %v", err)
		}
		if len(req.SignedTransactions) != 2 || req.SignedTransactions[0] == "" {
			t.Fatalf("unexpected payload: %+v", req)
		}
		json.NewEncoder(w).Encode(BatchResponse{
			Success:           true,
			TransactionHashes: []string{"0x111", "0x222"},
			BroadcastCount:    len(req.SignedTransactions),
			Message:           "submitted",
		})
	})

	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	client, err := NewClient(Config{
		IndexerBaseURL:         server.URL,
		EncryptionPublicKeyHex: serverWallet.PublicKeyHex,
	})
	if err != nil {
		t.Fatalf("client: %v", err)
	}

	deviceWallet, err := GenerateWallet()
	if err != nil {
		t.Fatalf("device wallet: %v", err)
	}
	signed := mustSignDemoTx(t, deviceWallet.PrivateKeyHex)
	res, err := client.SubmitTransactionBatch(context.Background(), BatchRequest{
		DeviceID:           "DEVICE1",
		DateComparable:     "20241108",
		OrderCount:         2,
		TotalAmount:        "100",
		SignedTransactions: []string{signed, signed},
	})
	if err != nil {
		t.Fatalf("submit batch: %v", err)
	}
	if !res.Success || res.BroadcastCount != len(res.TransactionHashes) {
		t.Fatalf("unexpected response: %+v", res)
	}
}

func mustSignDemoTx(t *testing.T, privateKeyHex string) string {
	t.Helper()
	to := common.HexToAddress("0x0000000000000000000000000000000000000000")
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(11155111),
		Nonce:     0,
		GasTipCap: big.NewInt(1_500_000_000),
		GasFeeCap: big.NewInt(2_000_000_000),
		Gas:       21_000,
		To:        &to,
		Value:     big.NewInt(0),
	})
	signed, err := SignTransaction(privateKeyHex, tx)
	if err != nil {
		t.Fatalf("sign tx: %v", err)
	}
	return signed
}
