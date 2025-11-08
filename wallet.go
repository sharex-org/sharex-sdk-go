package sharex

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Wallet bundles a secp256k1 keypair and its derived address.
type Wallet struct {
	PrivateKeyHex string
	PublicKeyHex  string
	Address       string
}

// GenerateWallet returns a new wallet backed by a random secp256k1 keypair.
func GenerateWallet() (*Wallet, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("generate key: %w", err)
	}

	return walletFromKey(key), nil
}

// WalletFromPrivateKey creates a wallet from an existing private key hex string.
func WalletFromPrivateKey(privateKeyHex string) (*Wallet, error) {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("parse private key: %w", err)
	}

	return walletFromKey(key), nil
}

func walletFromKey(key *ecdsa.PrivateKey) *Wallet {
	privBytes := crypto.FromECDSA(key)
	pubBytes := crypto.FromECDSAPub(&key.PublicKey)
	addr := crypto.PubkeyToAddress(key.PublicKey)

	return &Wallet{
		PrivateKeyHex: hexutil.Encode(privBytes),
		PublicKeyHex:  hexutil.Encode(pubBytes),
		Address:       strings.ToLower(addr.Hex()),
	}
}

// MustAddressToBytes decodes a checksum or lowercased address into bytes.
// It panics if the address is invalid and is intended for internal helpers only.
func MustAddressToBytes(address string) []byte {
	addr := common.HexToAddress(address)
	return addr.Bytes()
}

// NormalizeHex strips 0x prefixes and lowercases the payload.
func NormalizeHex(input string) (string, error) {
	cleaned := strings.TrimPrefix(strings.TrimSpace(input), "0x")
	if _, err := hex.DecodeString(cleaned); err != nil {
		return "", fmt.Errorf("invalid hex: %w", err)
	}
	return strings.ToLower(cleaned), nil
}
