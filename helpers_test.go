package sharex

import (
	"testing"
)

func TestGetNonce(t *testing.T) {
	// 使用一个已知的地址测试（Binance hot wallet）
	address := "0x28C6c06298d514Db089934071355E5743bf21d60"

	nonce, err := GetNonce(address, OpBNBMainnetRPC)
	if err != nil {
		t.Skipf("Skip test (network issue): %v", err)
		return
	}

	// Nonce 应该是非负数
	if nonce < 0 {
		t.Errorf("GetNonce() returned negative nonce: %d", nonce)
	}

	t.Logf("Address %s has nonce: %d", address, nonce)
}

func TestGetBalance(t *testing.T) {
	// 使用一个已知的地址测试
	address := "0x28C6c06298d514Db089934071355E5743bf21d60"

	balance, err := GetBalance(address, OpBNBMainnetRPC)
	if err != nil {
		t.Skipf("Skip test (network issue): %v", err)
		return
	}

	// Balance 不应该为 nil
	if balance == nil {
		t.Error("GetBalance() returned nil balance")
	}

	t.Logf("Address %s has balance: %s wei", address, balance.String())
}

func TestRPCConstants(t *testing.T) {
	tests := []struct {
		name string
		url  string
	}{
		{"OpBNB Mainnet", OpBNBMainnetRPC},
		{"OpBNB Testnet", OpBNBTestnetRPC},
		{"BSC Mainnet", BSCMainnetRPC},
		{"BSC Testnet", BSCTestnetRPC},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.url == "" {
				t.Errorf("%s RPC URL is empty", tt.name)
			}
			if len(tt.url) < 10 {
				t.Errorf("%s RPC URL seems invalid: %s", tt.name, tt.url)
			}
		})
	}
}
