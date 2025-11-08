package sharex

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultHTTPTimeout     = 15 * time.Second
	defaultRegisterPath    = "/devices"
	defaultTransactionPath = "/transactions/batch"
)

// Config aggregates every setting required to talk to the ShareX Indexer, including
// the base URL, an optional custom HTTP client, and the remote ECIES public key.
type Config struct {
	IndexerBaseURL         string
	HTTPClient             *http.Client
	EncryptionPublicKeyHex string
	Routes                 Routes
}

// Routes lets callers override the default `/devices` and `/transactions/batch`
// paths for local or pre-production environments.
type Routes struct {
	RegisterDevice   string
	TransactionBatch string
}

// Client wraps every interaction with the ShareX Indexer. It is safe for reuse
// across goroutines and encrypts each outbound request.
type Client struct {
	httpClient *http.Client
	baseURL    string
	routes     Routes
	encryptor  *Encryptor
}

// RegisterDeviceRequest mirrors the plaintext payload expected by the device
// registration endpoint so it can be JSON-serialized and encrypted directly.
type RegisterDeviceRequest struct {
	DeviceID    string `json:"deviceId"`
	DeviceType  string `json:"deviceType"`
	PartnerCode string `json:"partnerCode"`
	MerchantID  string `json:"merchantId"`
	WalletAddr  string `json:"walletAddress"`
}

// RegisterDeviceResponse describes the device registration response, including
// the transaction hash and optional server message.
type RegisterDeviceResponse struct {
	Success         bool   `json:"success"`
	TransactionHash string `json:"transactionHash"`
	Message         string `json:"message"`
	Error           string `json:"error,omitempty"`
}

// BatchRequest contains the plaintext data required to submit a batch of
// transactions. SignedTransactions holds fully signed raw transaction hex
// strings.
type BatchRequest struct {
	DeviceID           string   `json:"deviceId"`
	DateComparable     string   `json:"dateComparable"`
	OrderCount         int      `json:"orderCount"`
	TotalAmount        string   `json:"totalAmount"`
	SignedTransactions []string `json:"signedTransactions"`
}

// BatchResponse represents the outcome of a batch submission and may include
// optional compression metrics.
type BatchResponse struct {
	Success         bool       `json:"success"`
	TransactionHash string     `json:"transactionHash"`
	Message         string     `json:"message"`
	Error           string     `json:"error,omitempty"`
	BatchInfo       *BatchInfo `json:"batchInfo,omitempty"`
}

// BatchInfo exposes compression ratios and sizing metrics to help troubleshoot
// network and storage costs for batch uploads.
type BatchInfo struct {
	DeviceID         string  `json:"deviceId"`
	WalletAddress    string  `json:"walletAddress"`
	OrderCount       int     `json:"orderCount"`
	TotalAmount      string  `json:"totalAmount"`
	DateComparable   string  `json:"dateComparable"`
	CompressionRatio float64 `json:"compressionRatio,omitempty"`
	OriginalSize     int     `json:"originalSize"`
	CompressedSize   int     `json:"compressedSize"`
	SignedCount      int     `json:"signedCount,omitempty"`
}

// secureEnvelope is the canonical encrypted payload sent to the server.
type secureEnvelope struct {
	Payload string `json:"payload"`
}

// APIError captures readable errors returned by the Indexer and can be
// unwrapped via errors.As to inspect the status code or body.
type APIError struct {
	StatusCode int
	Message    string
	Body       []byte
}

func (e *APIError) Error() string {
	if e == nil {
		return "<nil>"
	}
	msg := strings.TrimSpace(e.Message)
	if msg == "" {
		msg = strings.TrimSpace(string(e.Body))
	}
	if msg == "" {
		msg = http.StatusText(e.StatusCode)
	}
	return fmt.Sprintf("sharex api error %d: %s", e.StatusCode, msg)
}

// NewClient builds a ShareX Indexer client: it validates the URL, fills default
// routes, initializes the HTTP client, and parses the encryption public key.
func NewClient(cfg Config) (*Client, error) {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{Timeout: defaultHTTPTimeout}
	}
	baseURL, err := normalizeBaseURL(cfg.IndexerBaseURL)
	if err != nil {
		return nil, fmt.Errorf("indexer url: %w", err)
	}

	routes := cfg.Routes.withDefaults()

	encryptor, err := NewEncryptor(cfg.EncryptionPublicKeyHex)
	if err != nil {
		return nil, err
	}

	return &Client{
		httpClient: cfg.HTTPClient,
		baseURL:    baseURL,
		routes:     routes,
		encryptor:  encryptor,
	}, nil
}

// RegisterDevice validates the request, encrypts it, sends it to `/devices`, and
// returns the resulting on-chain transaction hash on success.
func (c *Client) RegisterDevice(ctx context.Context, req RegisterDeviceRequest) (*RegisterDeviceResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	var result RegisterDeviceResponse
	if err := c.postEncrypted(ctx, c.routes.RegisterDevice, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SubmitTransactionBatch validates a batch request and posts it to
// `/transactions/batch`, typically for periodic offline uploads.
func (c *Client) SubmitTransactionBatch(ctx context.Context, req BatchRequest) (*BatchResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	var result BatchResponse
	if err := c.postEncrypted(ctx, c.routes.TransactionBatch, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) postEncrypted(ctx context.Context, path string, payload interface{}, out interface{}) error {
	if c == nil || c.encryptor == nil {
		return errors.New("client encryptor not configured")
	}
	plaintext, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("encode payload: %w", err)
	}
	cipherText, err := c.encryptor.Encrypt(plaintext)
	if err != nil {
		return fmt.Errorf("encrypt payload: %w", err)
	}

	endpoint := c.baseURL + path
	_, body, err := c.sendJSON(ctx, http.MethodPost, endpoint, secureEnvelope{Payload: cipherText})
	if err != nil {
		return err
	}

	if out != nil {
		if err := json.Unmarshal(body, out); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}
	return nil
}

func (c *Client) sendJSON(ctx context.Context, method, url string, payload interface{}) (*http.Response, []byte, error) {
	var body io.Reader
	if payload != nil {
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(payload); err != nil {
			return nil, nil, fmt.Errorf("encode payload: %w", err)
		}
		body = buf
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, nil, fmt.Errorf("build request: %w", err)
	}

	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return resp, data, parseAPIError(resp.StatusCode, data)
	}

	return resp, data, nil
}

func (r Routes) withDefaults() Routes {
	if r.RegisterDevice == "" {
		r.RegisterDevice = defaultRegisterPath
	}
	if r.TransactionBatch == "" {
		r.TransactionBatch = defaultTransactionPath
	}
	r.RegisterDevice = ensureLeadingSlash(r.RegisterDevice)
	r.TransactionBatch = ensureLeadingSlash(r.TransactionBatch)
	return r
}

func ensureLeadingSlash(p string) string {
	if p == "" {
		return p
	}
	if !strings.HasPrefix(p, "/") {
		return "/" + p
	}
	return p
}

func normalizeBaseURL(raw string) (string, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return "", errors.New("base URL is required")
	}
	parsed, err := url.Parse(trimmed)
	if err != nil {
		return "", fmt.Errorf("parse url: %w", err)
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return "", fmt.Errorf("invalid url: %s", raw)
	}
	parsed.Path = strings.TrimRight(parsed.Path, "/")
	parsed.RawQuery = ""
	parsed.Fragment = ""
	return strings.TrimRight(parsed.String(), "/"), nil
}

func parseAPIError(status int, body []byte) *APIError {
	apiErr := &APIError{StatusCode: status, Body: body}
	var payload struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}
	if err := json.Unmarshal(body, &payload); err == nil {
		if payload.Message != "" {
			apiErr.Message = payload.Message
		} else if payload.Error != "" {
			apiErr.Message = payload.Error
		}
	}
	return apiErr
}

// Validate ensures the registration payload contains every required field before
// sending it to the Indexer.
func (r RegisterDeviceRequest) Validate() error {
	if r.DeviceID == "" || r.DeviceType == "" || r.PartnerCode == "" || r.MerchantID == "" || r.WalletAddr == "" {
		return errors.New("deviceId, deviceType, partnerCode, merchantId, walletAddress are required")
	}
	return nil
}

// Validate ensures a batch request includes device metadata, ordering window,
// aggregate amounts, and non-empty signed transaction data.
func (r BatchRequest) Validate() error {
	if r.DeviceID == "" {
		return errors.New("deviceId is required")
	}
	if r.DateComparable == "" {
		return errors.New("dateComparable is required")
	}
	if r.OrderCount <= 0 {
		return errors.New("orderCount must be greater than zero")
	}
	if r.TotalAmount == "" {
		return errors.New("totalAmount is required")
	}
	if len(r.SignedTransactions) == 0 {
		return errors.New("signedTransactions is required")
	}
	for i, tx := range r.SignedTransactions {
		if strings.TrimSpace(tx) == "" {
			return fmt.Errorf("signedTransactions[%d] is empty", i)
		}
		if !strings.HasPrefix(tx, "0x") {
			return fmt.Errorf("signedTransactions[%d] must be hex encoded payload", i)
		}
	}
	return nil
}
