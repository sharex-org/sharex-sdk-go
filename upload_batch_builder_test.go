package sharex

import (
	"testing"
)

func TestDeriveBatchTotals(t *testing.T) {
	tests := []struct {
		name            string
		json            string
		wantOrderCount  int
		wantTotalAmount string
		wantErr         bool
		errContains     string
	}{
		{
			name:            "single transaction with amount",
			json:            `{"transactions":[{"amount":"10.50","order_no":"ORDER001"}]}`,
			wantOrderCount:  1,
			wantTotalAmount: "10.50",
			wantErr:         false,
		},
		{
			name:            "multiple transactions with amount",
			json:            `{"transactions":[{"amount":"10.00"},{"amount":"20.50"},{"amount":"5.25"}]}`,
			wantOrderCount:  3,
			wantTotalAmount: "35.75",
			wantErr:         false,
		},
		{
			name:            "transaction with factOvertimeMoney",
			json:            `{"transactions":[{"factOvertimeMoney":"99.99","cdb":"DEV001"}]}`,
			wantOrderCount:  1,
			wantTotalAmount: "99.99",
			wantErr:         false,
		},
		{
			name:            "transaction with money field",
			json:            `{"transactions":[{"money":"25.00","order_no":"ORDER002"}]}`,
			wantOrderCount:  1,
			wantTotalAmount: "25.00",
			wantErr:         false,
		},
		{
			name: "mixed amount fields (factOvertimeMoney priority)",
			json: `{"transactions":[
				{"factOvertimeMoney":"10.00","amount":"5.00"},
				{"amount":"15.50"},
				{"money":"20.00"}
			]}`,
			wantOrderCount:  3,
			wantTotalAmount: "45.50",
			wantErr:         false,
		},
		{
			name:        "empty transactions array",
			json:        `{"transactions":[]}`,
			wantErr:     true,
			errContains: "non-empty",
		},
		{
			name:        "missing transactions key",
			json:        `{"data":[{"amount":"10.00"}]}`,
			wantErr:     true,
			errContains: "non-empty",
		},
		{
			name:        "invalid JSON",
			json:        `{"transactions":`,
			wantErr:     true,
			errContains: "parse",
		},
		{
			name:        "empty string",
			json:        "",
			wantErr:     true,
			errContains: "empty",
		},
		{
			name:        "transaction missing amount",
			json:        `{"transactions":[{"order_no":"ORDER001"}]}`,
			wantErr:     true,
			errContains: "missing amount",
		},
		{
			name:            "string amounts",
			json:            `{"transactions":[{"amount":"10.50"},{"amount":"20.25"}]}`,
			wantOrderCount:  2,
			wantTotalAmount: "30.75",
			wantErr:         false,
		},
		{
			name:            "decimal precision",
			json:            `{"transactions":[{"amount":"0.01"},{"amount":"0.02"}]}`,
			wantOrderCount:  2,
			wantTotalAmount: "0.03",
			wantErr:         false,
		},
		{
			name:            "real-world example from error log",
			json:            `{"transactions":[{"amount":"0.01","order_no":"VI251128113013368010"}]}`,
			wantOrderCount:  1,
			wantTotalAmount: "0.01",
			wantErr:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderCount, totalAmount, err := DeriveBatchTotals(tt.json)

			if tt.wantErr {
				if err == nil {
					t.Errorf("DeriveBatchTotals() expected error containing %q, got nil", tt.errContains)
					return
				}
				if tt.errContains != "" && !contains(err.Error(), tt.errContains) {
					t.Errorf("DeriveBatchTotals() error = %v, want error containing %q", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("DeriveBatchTotals() unexpected error = %v", err)
				return
			}

			if orderCount != tt.wantOrderCount {
				t.Errorf("DeriveBatchTotals() orderCount = %d, want %d", orderCount, tt.wantOrderCount)
			}

			if totalAmount != tt.wantTotalAmount {
				t.Errorf("DeriveBatchTotals() totalAmount = %s, want %s", totalAmount, tt.wantTotalAmount)
			}
		})
	}
}

func TestBuildSignedUploadBatchTx_AutoCalculatesOrderCount(t *testing.T) {
	// Generate a test wallet
	wallet, err := GenerateWallet()
	if err != nil {
		t.Fatalf("generate wallet: %v", err)
	}

	// Build transaction with multiple transactions
	transactionJSON := `{
		"transactions": [
			{"amount": "10.00", "order_no": "ORDER001"},
			{"amount": "20.50", "order_no": "ORDER002"},
			{"amount": "5.25", "order_no": "ORDER003"}
		]
	}`

	signedTx, err := BuildSignedUploadBatchTx(UploadBatchTxParams{
		PrivateKeyHex:       wallet.PrivateKeyHex,
		DeviceID:            "TEST-DEVICE",
		DateComparable:      "20241128",
		TransactionDataJSON: transactionJSON,
		Nonce:               0,
	})

	if err != nil {
		t.Fatalf("BuildSignedUploadBatchTx() error = %v", err)
	}

	if signedTx == "" {
		t.Error("BuildSignedUploadBatchTx() returned empty transaction")
	}

	// Verify the orderCount was calculated correctly
	orderCount, totalAmount, err := DeriveBatchTotals(transactionJSON)
	if err != nil {
		t.Fatalf("DeriveBatchTotals() error = %v", err)
	}

	if orderCount != 3 {
		t.Errorf("Expected orderCount = 3, got %d", orderCount)
	}

	if totalAmount != "35.75" {
		t.Errorf("Expected totalAmount = 35.75, got %s", totalAmount)
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
