package sharex

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestEncryptor(t *testing.T) {
	wallet, err := GenerateWallet()
	if err != nil {
		t.Fatalf("wallet: %v", err)
	}

	enc, err := NewEncryptor(wallet.PublicKeyHex)
	if err != nil {
		t.Fatalf("new encryptor: %v", err)
	}

	cipher, err := enc.Encrypt([]byte("hello"))
	if err != nil {
		t.Fatalf("encrypt: %v", err)
	}

	priv, err := crypto.HexToECDSA(strings.TrimPrefix(wallet.PrivateKeyHex, "0x"))
	if err != nil {
		t.Fatalf("priv: %v", err)
	}

	plain, err := decryptForTests(priv, cipher)
	if err != nil {
		t.Fatalf("decrypt: %v", err)
	}

	if string(plain) != "hello" {
		t.Fatalf("unexpected plaintext: %s", string(plain))
	}
}
