package sharex

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/sharex-org/sharex-sdk-go/contracts/deshare"
)

const (
	// DefaultDeshareContractAddress points to the production Deshare contract.
	DefaultDeshareContractAddress = "0x28e3889A3bc57D4421a5041E85Df8b516Ab683F8"
	// DefaultOpBNBChainID is the opBNB mainnet chain id.
	DefaultOpBNBChainID int64 = 204
)

// UploadBatchTxParams collects every field needed to build and sign a single
// uploadTransactionBatch call to the Deshare contract. The SDK will:
// 1) parse transactionDataJSON and derive orderCount & totalAmount,
// 2) compress the JSON using FastLZ (compatible with LibZip.flzDecompress),
// 3) ABI-pack the call, build an EIP-1559 transaction, and sign it.
type UploadBatchTxParams struct {
	PrivateKeyHex       string
	ContractAddress     string
	ChainID             *big.Int
	Nonce               uint64
	GasTipCap           *big.Int // optional; defaults to 1.5 gwei when nil
	GasFeeCap           *big.Int // optional; defaults to 2.5 gwei when nil
	GasLimit            uint64   // optional; defaults to 500_000 when zero
	DeviceID            string
	DateComparable      string
	TransactionDataJSON string
}

// BuildSignedUploadBatchTx builds and signs a raw transaction hex that calls
// deshare.uploadTransactionBatch with properly compressed transactionData.
func BuildSignedUploadBatchTx(p UploadBatchTxParams) (string, error) {
	if strings.TrimSpace(p.PrivateKeyHex) == "" {
		return "", errors.New("privateKeyHex is required")
	}
	if p.ChainID == nil {
		p.ChainID = big.NewInt(DefaultOpBNBChainID)
	}
	if p.ContractAddress == "" {
		p.ContractAddress = DefaultDeshareContractAddress
	}
	if p.DeviceID == "" || p.DateComparable == "" {
		return "", errors.New("deviceId and dateComparable are required")
	}
	if strings.TrimSpace(p.TransactionDataJSON) == "" {
		return "", errors.New("transactionDataJSON is required")
	}

	orderCount, totalAmount, err := deriveBatchTotals(p.TransactionDataJSON)
	if err != nil {
		return "", err
	}

	// FastLZ compress the JSON payload – must match on-chain LibZip.flzDecompress.
	compressed := FlzCompress([]byte(p.TransactionDataJSON))

	abiDef, err := deshare.DeshareMetaData.GetAbi()
	if err != nil {
		return "", fmt.Errorf("load deshare ABI: %w", err)
	}

	callData, err := abiDef.Pack("uploadTransactionBatch", deshare.UploadBatchParams{
		DeviceId:        p.DeviceID,
		DateComparable:  p.DateComparable,
		OrderCount:      uint32(orderCount),
		TotalAmount:     totalAmount,
		TransactionData: compressed,
	})
	if err != nil {
		return "", fmt.Errorf("pack uploadTransactionBatch: %w", err)
	}

	gasTip := p.GasTipCap
	if gasTip == nil {
		gasTip = big.NewInt(1_500_000_000) // 1.5 gwei
	}
	gasFee := p.GasFeeCap
	if gasFee == nil {
		gasFee = big.NewInt(2_500_000_000) // 2.5 gwei
	}
	gasLimit := p.GasLimit
	if gasLimit == 0 {
		gasLimit = 500_000
	}

	to := common.HexToAddress(p.ContractAddress)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   p.ChainID,
		Nonce:     p.Nonce,
		GasTipCap: gasTip,
		GasFeeCap: gasFee,
		Gas:       gasLimit,
		To:        &to,
		Value:     big.NewInt(0),
		Data:      callData,
	})

	return SignTransaction(p.PrivateKeyHex, tx)
}

// deriveBatchTotals parses transactionDataJSON to compute orderCount and totalAmount.
// It expects a top-level "transactions" array. Amount is read from
// factOvertimeMoney, amount, or money fields, in that order.
func deriveBatchTotals(rawJSON string) (int, string, error) {
	decoder := json.NewDecoder(strings.NewReader(rawJSON))
	decoder.UseNumber()

	var payload struct {
		Transactions []map[string]interface{} `json:"transactions"`
	}
	if err := decoder.Decode(&payload); err != nil {
		return 0, "", fmt.Errorf("parse transactionData JSON: %w", err)
	}
	if len(payload.Transactions) == 0 {
		return 0, "", errors.New(`transactionData must contain a non-empty "transactions" array`)
	}

	total := new(big.Rat)
	for idx, tx := range payload.Transactions {
		amt, ok := firstAmount(tx, "factOvertimeMoney", "amount", "money")
		if !ok {
			return 0, "", fmt.Errorf("transactions[%d] missing amount (factOvertimeMoney/amount/money)", idx)
		}
		total.Add(total, amt)
	}

	// Format with 2 decimal places for consistency with existing API.
	totalAmount := formatRat(total, 2)
	return len(payload.Transactions), totalAmount, nil
}

func firstAmount(m map[string]interface{}, keys ...string) (*big.Rat, bool) {
	for _, k := range keys {
		if v, ok := m[k]; ok {
			if amt := toRat(v); amt != nil {
				return amt, true
			}
		}
	}
	return nil, false
}

func toRat(v interface{}) *big.Rat {
	switch t := v.(type) {
	case json.Number:
		if r, ok := new(big.Rat).SetString(t.String()); ok {
			return r
		}
		if f, err := t.Float64(); err == nil {
			return new(big.Rat).SetFloat64(f)
		}
	case string:
		if r, ok := new(big.Rat).SetString(strings.TrimSpace(t)); ok {
			return r
		}
		if f, err := strconv.ParseFloat(strings.TrimSpace(t), 64); err == nil {
			return new(big.Rat).SetFloat64(f)
		}
	case float64:
		return new(big.Rat).SetFloat64(t)
	}
	return nil
}

func formatRat(r *big.Rat, decimals int) string {
	if r == nil {
		return "0"
	}
	// Scale by 10^decimals, round to nearest integer, then format.
	scale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	scaled := new(big.Rat).Mul(r, new(big.Rat).SetInt(scale))
	n, _ := scaled.Float64()
	rounded := int64(n + 0.5)
	result := new(big.Rat).SetInt(big.NewInt(rounded))
	result.Quo(result, new(big.Rat).SetInt(scale))
	return result.FloatString(decimals)
}
