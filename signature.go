package sharex

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// DeviceRegistrationSignaturePayload captures the fields that need to be signed
// when asking ShareX to bind a device ID to a wallet. Adding the endpoint path
// prevents replaying the signature against multiple routes.
type DeviceRegistrationSignaturePayload struct {
	DeviceID      string
	WalletAddress string
	Timestamp     int64
	Endpoint      string
}

// Encode converts the payload into a deterministic byte slice for signing.
func (p DeviceRegistrationSignaturePayload) Encode() []byte {
	builder := strings.Builder{}
	builder.WriteString(strings.TrimSpace(p.DeviceID))
	builder.WriteString("|")
	builder.WriteString(strings.ToLower(strings.TrimSpace(p.WalletAddress)))
	builder.WriteString("|")
	builder.WriteString(fmt.Sprintf("%d", p.Timestamp))
	builder.WriteString("|")
	builder.WriteString(strings.TrimSpace(p.Endpoint))
	return []byte(builder.String())
}

// NewDeviceRegistrationPayload creates a payload anchored to the current time.
func NewDeviceRegistrationPayload(deviceID, walletAddress, endpoint string) DeviceRegistrationSignaturePayload {
	return DeviceRegistrationSignaturePayload{
		DeviceID:      deviceID,
		WalletAddress: walletAddress,
		Timestamp:     time.Now().Unix(),
		Endpoint:      endpoint,
	}
}

// SignDeviceRegistration produces an ECDSA signature (secp256k1) for the payload.
func SignDeviceRegistration(privateKeyHex string, payload DeviceRegistrationSignaturePayload) (string, error) {
	return SignMessage(privateKeyHex, payload.Encode())
}

// SignMessage keccak-hashes the message and signs it using secp256k1.
func SignMessage(privateKeyHex string, message []byte) (string, error) {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return "", fmt.Errorf("parse private key: %w", err)
	}

	hash := crypto.Keccak256(message)
	sig, err := crypto.Sign(hash, key)
	if err != nil {
		return "", fmt.Errorf("sign message: %w", err)
	}

	return hexutil.Encode(sig), nil
}

// SignatureVerifier validates signatures returned by the ShareX server.
type SignatureVerifier interface {
	Verify(message []byte, signatureHex string) error
}

// ECDSAVerifier verifies keccak256 hashed payloads against a known public key.
type ECDSAVerifier struct {
	key *ecdsa.PublicKey
}

// NewECDSAVerifier builds a verifier from an uncompressed secp256k1 public key hex string.
func NewECDSAVerifier(publicKeyHex string) (*ECDSAVerifier, error) {
	decoded, err := hexutil.Decode(strings.TrimSpace(publicKeyHex))
	if err != nil {
		return nil, fmt.Errorf("decode public key: %w", err)
	}

	pubKey, err := crypto.UnmarshalPubkey(decoded)
	if err != nil {
		return nil, fmt.Errorf("unmarshal public key: %w", err)
	}

	return &ECDSAVerifier{key: pubKey}, nil
}

// Verify ensures the signature matches the provided payload.
func (v *ECDSAVerifier) Verify(message []byte, signatureHex string) error {
	if v == nil || v.key == nil {
		return errors.New("verifier not configured")
	}

	sigBytes, err := hexutil.Decode(signatureHex)
	if err != nil {
		return fmt.Errorf("decode signature: %w", err)
	}

	if len(sigBytes) < 64 {
		return fmt.Errorf("signature must be at least 64 bytes, got %d", len(sigBytes))
	}

	hash := crypto.Keccak256(message)

	// Drop recovery ID if present (65 byte signature)
	if len(sigBytes) == 65 {
		sigBytes = sigBytes[:64]
	}

	r := new(big.Int).SetBytes(sigBytes[:32])
	s := new(big.Int).SetBytes(sigBytes[32:])

	if !ecdsa.Verify(v.key, hash, r, s) {
		return errors.New("invalid signature")
	}

	return nil
}
