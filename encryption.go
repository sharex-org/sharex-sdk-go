package sharex

import (
	"crypto/ecdsa"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	eciesgo "github.com/ecies/go/v2"
)

// Encryptor uses the ShareX-provided secp256k1 public key to ECIES-encrypt
// request bodies using eciesjs-compatible format.
type Encryptor struct {
	pub *eciesgo.PublicKey
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

	// Parse as ECDSA public key first
	ecdsaPubKey, err := crypto.UnmarshalPubkey(decoded)
	if err != nil {
		return nil, fmt.Errorf("unmarshal public key: %w", err)
	}

	// Convert to eciesgo public key
	eciesPubKey, err := eciesgo.NewPublicKeyFromBytes(crypto.FromECDSAPub(ecdsaPubKey))
	if err != nil {
		return nil, fmt.Errorf("create ecies public key: %w", err)
	}

	return &Encryptor{pub: eciesPubKey}, nil
}

// Encrypt runs ECIES over the given bytes and returns a base64 string.
// This uses the eciesjs-compatible format.
func (e *Encryptor) Encrypt(plaintext []byte) (string, error) {
	if e == nil || e.pub == nil {
		return "", errors.New("encryptor is not initialized")
	}
	cipher, err := eciesgo.Encrypt(e.pub, plaintext)
	if err != nil {
		return "", fmt.Errorf("encrypt payload: %w", err)
	}
	return base64.StdEncoding.EncodeToString(cipher), nil
}

// decryptForTests is only used by unit tests to decrypt a base64 payload with a
// private key. This uses the eciesjs-compatible format.
func decryptForTests(privateKey *ecdsa.PrivateKey, payload string) ([]byte, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}

	// Convert ECDSA private key to eciesgo private key
	privKeyBytes := crypto.FromECDSA(privateKey)
	eciesPrivKey := eciesgo.NewPrivateKeyFromBytes(privKeyBytes)

	return eciesgo.Decrypt(eciesPrivKey, cipherBytes)
}
