package sharex

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// SignTransaction signs a single Ethereum transaction and returns a 0x-prefixed
// raw payload ready for RPC submission.
func SignTransaction(privateKeyHex string, tx *types.Transaction) (string, error) {
	if tx == nil {
		return "", errors.New("transaction is nil")
	}
	chainID := tx.ChainId()
	if chainID == nil {
		return "", errors.New("transaction missing chain id")
	}
	if strings.TrimSpace(privateKeyHex) == "" {
		return "", errors.New("private key is required")
	}

	key, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return "", fmt.Errorf("parse private key: %w", err)
	}

	signer := types.LatestSignerForChainID(chainID)
	signedTx, err := types.SignTx(tx, signer, key)
	if err != nil {
		return "", fmt.Errorf("sign transaction: %w", err)
	}

	raw, err := signedTx.MarshalBinary()
	if err != nil {
		return "", fmt.Errorf("marshal signed transaction: %w", err)
	}

	return hexutil.Encode(raw), nil
}

// SignTransactions signs multiple transactions in order and returns their raw
// hex payloads.
func SignTransactions(privateKeyHex string, txs []*types.Transaction) ([]string, error) {
	if len(txs) == 0 {
		return nil, errors.New("transactions list is empty")
	}

	result := make([]string, len(txs))
	for i, tx := range txs {
		payload, err := SignTransaction(privateKeyHex, tx)
		if err != nil {
			return nil, fmt.Errorf("sign transaction %d: %w", i, err)
		}
		result[i] = payload
	}
	return result, nil
}
