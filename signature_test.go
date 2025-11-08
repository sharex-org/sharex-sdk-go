package sharex

import "testing"

func TestSignAndVerifyDeviceRegistration(t *testing.T) {
	wallet, err := GenerateWallet()
	if err != nil {
		t.Fatalf("generate wallet: %v", err)
	}

	payload := DeviceRegistrationSignaturePayload{
		DeviceID:      "DEVICE123",
		WalletAddress: wallet.Address,
		Timestamp:     1731000000,
		Endpoint:      "/api/v1/devices/register",
	}

	sig, err := SignDeviceRegistration(wallet.PrivateKeyHex, payload)
	if err != nil {
		t.Fatalf("sign payload: %v", err)
	}

	verifier, err := NewECDSAVerifier(wallet.PublicKeyHex)
	if err != nil {
		t.Fatalf("new verifier: %v", err)
	}

	if err := verifier.Verify(payload.Encode(), sig); err != nil {
		t.Fatalf("verify signature: %v", err)
	}

	payload.DeviceID = "DEVICE999"
	if err := verifier.Verify(payload.Encode(), sig); err == nil {
		t.Fatal("expected verification to fail on tampered payload")
	}
}
