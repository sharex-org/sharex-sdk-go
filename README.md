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
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/sharex-org/sharex-sdk-go"
)

func main() {
    client, err := sharex.NewClient(sharex.Config{
        IndexerBaseURL:         "https://indexer-api.sharex.network",
        EncryptionPublicKeyHex: "<sharex-indexer-public-key>", // Uncompressed secp256k1 public key (0x04...)
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

    // 3. Sign transactions and upload them as an encrypted batch
    tx := types.NewTx(&types.DynamicFeeTx{
        ChainID:   big.NewInt(11155111), // Sepolia testnet example
        Nonce:     0,
        GasTipCap: big.NewInt(1_500_000_000),
        GasFeeCap: big.NewInt(2_000_000_000),
        Gas:       21_000,
        To:        common.HexToAddress("0x0000000000000000000000000000000000000000"),
        Value:     big.NewInt(0),
    })
    signedTx, err := sharex.SignTransaction(wallet.PrivateKeyHex, tx)
    if err != nil {
        log.Fatal(err)
    }

    batch := sharex.BatchRequest{
        DeviceID:           "DEVICE001",
        DateComparable:     "20241108",
        OrderCount:         1,
        TotalAmount:        "0",
        SignedTransactions: []string{signedTx},
    }
    if _, err := client.SubmitTransactionBatch(context.Background(), batch); err != nil {
        log.Fatal(err)
    }
}
```

Device requests rely solely on the ECIES public key for encryption and implicit
auth, so no `x-api-key` header is required. Keep
`EncryptionPublicKeyHex` in sync with the Indexer configuration.

## Interaction Flow

```mermaid
sequenceDiagram
    autonumber
    participant Terminal as Terminal Device
    participant SDK as ShareX SDK
    participant Indexer as ShareX Indexer API
    participant Chain as opBNB Chain

    Terminal->>SDK: Initialize client + generate wallet
    Terminal->>SDK: RegisterDeviceRequest
    SDK->>Indexer: ECIES encrypted POST /sdk/register-device
    Indexer-->>Terminal: Registration result / tx hash

    Terminal->>SDK: Build outbound transactions
    Terminal->>SDK: SignTransactions(privateKey, txs)
    SDK-->>Terminal: Return 0x raw transactions

    Terminal->>SDK: BatchRequest (SignedTransactions)
    SDK->>Indexer: ECIES encrypted POST /sdk/upload
    Indexer->>Chain: Relay signed transactions to opBNB RPC
    Chain-->>Indexer: Transaction hashes / status
    Indexer-->>Terminal: BatchResponse (transactionHashes/broadcastCount)
```

## Encryption Pipeline

1. `RegisterDeviceRequest` or `BatchRequest` is serialized to JSON.
2. The SDK reads `EncryptionPublicKeyHex` and runs ECIES (secp256k1) with a
   random session key to encrypt the JSON.
3. The payload sent to the Indexer always looks like:
   ```json
   {
     "payload": "<base64-cipher-text>"
   }
   ```
4. The Indexer decrypts with its private key before inserting the device or
   writing the batch on-chain.

Therefore, proxies outside of the Indexer or its gateway cannot capture device
IDs, wallet addresses, or transaction details.

## Request Fields

### Device Registration `RegisterDeviceRequest`

| Field           | Description                                                    |
| --------------- | -------------------------------------------------------------- |
| `deviceId`      | Unique device identifier (hardware serial, secure element ID). |
| `deviceType`    | Device class (Terminal/Mobile/...).                            |
| `partnerCode`   | Partner identifier.                                            |
| `merchantId`    | Merchant identifier.                                           |
| `walletAddress` | Device wallet address (EVM).                                   |

### Batch Upload `BatchRequest`

| Field                | Description                                                               |
| -------------------- | ------------------------------------------------------------------------- |
| `deviceId`           | Previously registered device ID.                                          |
| `dateComparable`     | `YYYYMMDD` string used for server-side windowing.                         |
| `orderCount`         | Number of records inside the batch (>0).                                  |
| `totalAmount`        | Aggregate amount encoded as a string to avoid floating-point drift.       |
| `signedTransactions` | Array of signed raw transactions (0x+hex) that the Indexer relays to RPC. |

## Error Handling

All 4xx/5xx responses are converted into `*sharex.APIError` (with
`StatusCode`, `Message`, and `Body`). Use `errors.As(err, *sharex.APIError)` to
inspect the status and implement retries or key renegotiation strategies.

## Testing

```bash
cd sharex-sdk-go && go test ./...
```

Tests cover:
- ECIES encryption/decryption correctness.
- Whether the register/batch flows send ciphertext that the mock server can
  restore.
- Core utilities such as URL normalization and validation helpers.

Follow the repository guidelines in `AGENTS.md`: keep `gofmt`, ship tests with
new capabilities, and extend `Routes` if you need multiple environments.

## Demo: Wallet Export + End-to-End Validation

```bash
go run ./cmd/demo
```

This single binary now demonstrates the full lifecycle:

- Generates a wallet, saves the private key to a temp `sharex-wallet-*.key` file with `0600` permissions, and reloads it via `WalletFromPrivateKey` to prove deterministic recovery.
- Spins up a mock Indexer, registers the device with ECIES-encrypted payloads, and submits a batch of signed Sepolia transactions through the encrypted channel.

Use the printed wallet path if you want to inspect or back up the key; delete it once you're done.

`cmd/demo` ships a minimal example that starts a mock Indexer locally:

1. The server generates a key pair, exposes the public key, and keeps the
   private key for decryption.
2. The SDK device generates a wallet, calls `RegisterDevice`, and the server
   decrypts/logs the device data.
3. The SDK constructs a batch, calls `SubmitTransactionBatch`, and the server
   decrypts/logs the batch statistics.

Run it with:

```bash
cd sharex-sdk-go
go run ./cmd/demo
```

The terminal prints the device wallet, registration result, batch outcome, and
mock server logs, confirming the “wallet → encrypted request → decrypted
processing” loop.
