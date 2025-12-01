package sharex

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GetNonce queries the current nonce for a wallet address from the blockchain.
// This is essential for submitting transactions as the nonce must be correct.
//
// Example usage:
//
//	nonce, err := sharex.GetNonce(wallet.Address, "https://opbnb-mainnet-rpc.bnbchain.org")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	resp, err := client.UploadTransactionBatch(ctx, sharex.UploadTransactionBatchParams{
//	    // ...
//	    Nonce: nonce,  // Use the queried nonce
//	})
func GetNonce(walletAddress string, rpcURL string) (uint64, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return 0, fmt.Errorf("connect to RPC: %w", err)
	}
	defer client.Close()

	address := common.HexToAddress(walletAddress)
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, fmt.Errorf("query nonce: %w", err)
	}

	return nonce, nil
}

// GetBalance queries the current balance for a wallet address from the blockchain.
// Returns balance in wei.
//
// Example usage:
//
//	balance, err := sharex.GetBalance(wallet.Address, "https://opbnb-mainnet-rpc.bnbchain.org")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	balanceInBNB := float64(balance.Int64()) / 1e18
//	fmt.Printf("Balance: %.6f BNB\n", balanceInBNB)
func GetBalance(walletAddress string, rpcURL string) (*big.Int, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("connect to RPC: %w", err)
	}
	defer client.Close()

	address := common.HexToAddress(walletAddress)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, fmt.Errorf("query balance: %w", err)
	}

	return balance, nil
}

// GetGasPrice queries the current gas price from the blockchain.
// Returns the suggested gas price in wei.
//
// Example usage:
//
//	gasPrice, err := sharex.GetGasPrice("https://opbnb-mainnet-rpc.bnbchain.org")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	fmt.Printf("Current gas price: %s gwei\n", gasPrice.String())
func GetGasPrice(rpcURL string) (*big.Int, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("connect to RPC: %w", err)
	}
	defer client.Close()

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("query gas price: %w", err)
	}

	return gasPrice, nil
}

// GasEstimate contains estimated gas parameters for a transaction.
type GasEstimate struct {
	GasTipCap *big.Int // Suggested tip (priority fee)
	GasFeeCap *big.Int // Suggested max fee
	GasLimit  uint64   // Estimated gas limit
	GasPrice  *big.Int // Legacy gas price (for non-EIP-1559)
}

// EstimateGas queries the blockchain and returns suggested gas parameters.
// This is more accurate than using fixed values.
//
// Example usage:
//
//	estimate, err := sharex.EstimateGas("https://opbnb-mainnet-rpc.bnbchain.org")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Use in transaction
//	params := sharex.UploadTransactionBatchParams{
//	    GasTipCap: estimate.GasTipCap,
//	    GasFeeCap: estimate.GasFeeCap,
//	    // ...
//	}
func EstimateGas(rpcURL string) (*GasEstimate, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("connect to RPC: %w", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Get current gas price
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("query gas price: %w", err)
	}

	// Try to get EIP-1559 gas parameters
	gasTipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		// Fallback to legacy gas price if EIP-1559 not supported
		gasTipCap = gasPrice
	}

	// GasFeeCap = 2 × baseFee + GasTipCap (following EIP-1559 recommendation)
	// For simplicity, use gasPrice + gasTipCap
	gasFeeCap := new(big.Int).Add(gasPrice, gasTipCap)

	// Default gas limit for uploadTransactionBatch
	// This is a safe estimate; actual usage is typically lower
	const defaultGasLimit = 300_000

	return &GasEstimate{
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		GasLimit:  defaultGasLimit,
		GasPrice:  gasPrice,
	}, nil
}

// Common RPC URLs
const (
	// OpBNBMainnetRPC is the public RPC endpoint for opBNB mainnet
	OpBNBMainnetRPC = "https://opbnb-mainnet-rpc.bnbchain.org"

	// OpBNBTestnetRPC is the public RPC endpoint for opBNB testnet
	OpBNBTestnetRPC = "https://opbnb-testnet-rpc.bnbchain.org"

	// BSCMainnetRPC is the public RPC endpoint for BSC mainnet
	BSCMainnetRPC = "https://bsc-dataseed.binance.org"

	// BSCTestnetRPC is the public RPC endpoint for BSC testnet
	BSCTestnetRPC = "https://data-seed-prebsc-1-s1.binance.org:8545"
)
