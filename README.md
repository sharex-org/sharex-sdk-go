# ShareX Go SDK

This SDK serves ShareX terminals: devices generate wallets locally, talk only to
the ShareX Indexer API, and rely on **public-key encryption** to upload
device-registration details and batched transaction data. The server decrypts
with its private key and forwards payloads to the Deshare contracts.

## Feature Overview

- secp256k1 wallet generation/import (`GenerateWallet`, `WalletFromPrivateKey`).
- ECIES (secp256k1) public-key encryption so every request body becomes
  ciphertext via the Indexer key, keeping device identifiers and wallet
  addresses off the wire.
- Indexer integrations: `/sdk/register-device` (register device) and
  `/sdk/upload` (batch submission) both run through the encrypted channel.
- Unified error wrapper `APIError` to quickly inspect HTTP status codes and
  response bodies.

## Installation

```bash
go get github.com/sharex-org/sharex-sdk-go
```

## Quick Start

```go
package main

import (
    "context"
    "log"

    "github.com/sharex-org/sharex-sdk-go"
)

func main() {
    client, err := sharex.NewClient(sharex.Config{
        IndexerBaseURL:         "https://your-indexer-url.com",
        EncryptionPublicKeyHex: "0x04...", // Get from your indexer configuration
    })
    if err != nil {
        log.Fatal(err)
    }

    // 1. Generate a wallet on the device
    wallet, err := sharex.GenerateWallet()
    if err != nil {
        log.Fatal(err)
    }

    // 2. Register the device through an encrypted request
    if _, err := client.RegisterDevice(context.Background(), sharex.RegisterDeviceRequest{
        DeviceID:    "DEVICE001",
        DeviceType:  "Terminal",
        PartnerCode: "TECH001",
        MerchantID:  "MERCH001",
        WalletAddr:  wallet.Address,
    }); err != nil {
        log.Fatal(err)
    }

    // 3. Upload transaction batch
    //    SDK automatically handles all technical details:
    //    - Querying nonce from blockchain
    //    - Estimating gas price from network
    //    - Calculating orderCount and totalAmount
    //    - Building and signing transaction
    //    - Encrypting and submitting to indexer
    resp, err := client.UploadTransactionBatch(context.Background(), sharex.UploadTransactionBatchParams{
        DeviceID:            "DEVICE001",
        DateComparable:      "20250131",
        TransactionDataJSON: `{"transactions":[{"amount":"99.99","order_no":"ORDER001"}]}`,
        PrivateKeyHex:       wallet.PrivateKeyHex,
    })
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Success! Broadcast %d transactions\n", resp.BroadcastCount)
}
```

## Key Features

### Automatic Calculation
- **Nonce**: Automatically queried from blockchain
- **Gas Price**: Automatically estimated from network  
- **OrderCount**: Automatically calculated from transaction array
- **TotalAmount**: Automatically summed from transaction amounts

### Transaction Amount Fields
The SDK supports multiple amount field names (priority order):
1. `factOvertimeMoney`
2. `amount`
3. `money`

## Configuration

### Default Values
- Chain: opBNB Mainnet (chainId=204)
- Contract: Production Deshare contract
- Gas: Automatically estimated from network
- Nonce: Automatically queried from blockchain

All defaults can be overridden via optional parameters.

## Error Handling

```go
resp, err := client.UploadTransactionBatch(ctx, params)
if err != nil {
    var apiErr *sharex.APIError
    if errors.As(err, &apiErr) {
        log.Printf("API error %d: %s", apiErr.StatusCode, apiErr.Message)
    }
    return err
}
```

## Testing

```bash
go test ./...
```

## Demo

```bash
go run ./cmd/demo
```

Demonstrates wallet generation, device registration, and transaction upload with a mock indexer.
