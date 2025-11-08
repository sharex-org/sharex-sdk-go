package sharex

import "testing"

func TestGenerateWallet(t *testing.T) {
	w, err := GenerateWallet()
	if err != nil {
		t.Fatalf("generate wallet: %v", err)
	}
	if len(w.PrivateKeyHex) == 0 || len(w.PublicKeyHex) == 0 || len(w.Address) == 0 {
		t.Fatalf("wallet fields should not be empty: %+v", w)
	}
	if w.PrivateKeyHex[:2] != "0x" || w.PublicKeyHex[:2] != "0x" {
		t.Fatalf("wallet hex fields should be prefixed with 0x: %+v", w)
	}
}

func TestWalletFromPrivateKey(t *testing.T) {
	seed, err := GenerateWallet()
	if err != nil {
		t.Fatalf("seed wallet: %v", err)
	}

	restored, err := WalletFromPrivateKey(seed.PrivateKeyHex)
	if err != nil {
		t.Fatalf("restore wallet: %v", err)
	}

	if restored.Address != seed.Address {
		t.Fatalf("addresses mismatch: got %s want %s", restored.Address, seed.Address)
	}
}
