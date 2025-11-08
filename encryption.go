package sharex

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

// Encryptor uses the ShareX-provided secp256k1 public key to ECIES-encrypt
// request bodies.
type Encryptor struct {
	pub *ecies.PublicKey
}

// NewEncryptor parses an uncompressed 0x-prefixed public key and returns an
// encryptor instance.
func NewEncryptor(publicKeyHex string) (*Encryptor, error) {
	cleaned := strings.TrimSpace(publicKeyHex)
	if cleaned == "" {
		return nil, errors.New("encryption public key is required")
	}

	decoded, err := hexutil.Decode(cleaned)
	if err != nil {
		return nil, fmt.Errorf("decode public key: %w", err)
	}

	pubKey, err := crypto.UnmarshalPubkey(decoded)
	if err != nil {
		return nil, fmt.Errorf("unmarshal public key: %w", err)
	}

	return &Encryptor{pub: ecies.ImportECDSAPublic(pubKey)}, nil
}

// Encrypt runs ECIES over the given bytes and returns a base64 string.
func (e *Encryptor) Encrypt(plaintext []byte) (string, error) {
	if e == nil || e.pub == nil {
		return "", errors.New("encryptor is not initialized")
	}
	cipher, err := ecies.Encrypt(rand.Reader, e.pub, plaintext, nil, nil)
	if err != nil {
		return "", fmt.Errorf("encrypt payload: %w", err)
	}
	return base64.StdEncoding.EncodeToString(cipher), nil
}

// decryptForTests is only used by unit tests to decrypt a base64 payload with a
// private key.
func decryptForTests(privateKey *ecdsa.PrivateKey, payload string) ([]byte, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}
	return ecies.ImportECDSA(privateKey).Decrypt(cipherBytes, nil, nil)
}
