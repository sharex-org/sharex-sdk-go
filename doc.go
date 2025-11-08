// Package sharex is the official ShareX terminal SDK. It generates secp256k1
// wallets locally, encrypts device-registration and batch-transaction requests
// using the Indexer's ECIES public key, and invokes `/devices` or
// `/transactions/batch` before anything touches the blockchain.
//
// Core capabilities:
//  1. `GenerateWallet`/`WalletFromPrivateKey`: wrap wallet generation, import,
//     and address derivation.
//  2. `SignTransaction`/`SignTransactions`: sign EVM transactions into bytes
//     that are ready for `eth_sendRawTransaction`.
//  3. `NewClient` plus `RegisterDevice`/`SubmitTransactionBatch`: automatically
//     encrypt payloads, send HTTP requests, and unwrap errors.
//  4. `Encryptor` and verification helpers that can be reused in custom flows
//     or security audits.
//
// The typical flow is “generate wallet → register device → sign transactions →
// periodically upload batches.” Every outbound request is ECIES encrypted, and
// `APIError` surfaces readable failures to keep resource-constrained devices
// safe across the entire interaction.
package sharex
