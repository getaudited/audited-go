package audited

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const Version = "v0.0.1"

var ErrUnauthorized = errors.New("unauthorized")
var ErrEventTypeNotFound = errors.New("event type not found")

type Client struct {
	baseAPI    string
	apiToken   string
	httpClient *http.Client
}

type Config struct {
	BaseAPI    string
	APIToken   string
	HttpClient *http.Client
	Timeout    time.Duration
}

func NewClient(cfg Config) (*Client, error) {
	if strings.TrimSpace(cfg.BaseAPI) == "" {
		return nil, errors.New("audited-go: BaseAPI cannot be empty")
	}

	if strings.TrimSpace(cfg.APIToken) == "" {
		return nil, errors.New("audited-go: APIToken cannot be empty")
	}

	timeout := time.Second * 5
	if cfg.Timeout != 0 {
		timeout = cfg.Timeout
	}

	httpClient := &http.Client{
		Timeout: timeout,
	}
	if cfg.HttpClient != nil {
		httpClient = cfg.HttpClient
	}

	return &Client{
		baseAPI:    cfg.BaseAPI,
		apiToken:   cfg.APIToken,
		httpClient: httpClient,
	}, nil
}

func (c *Client) CreateEvent(ctx context.Context, event Event) error {
	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("audited-go: error marshalling event: %v", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/events", c.baseAPI),
		bytes.NewReader(body),
	)
	if err != nil {
		return fmt.Errorf("audited-go: error creating request: %w", err)
	}

	req.Header.Add("x-token", c.apiToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", fmt.Sprintf("audited-go/%s", Version))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode == http.StatusUnauthorized {
		return ErrUnauthorized
	}

	if resp.StatusCode == http.StatusNotFound {
		return ErrEventTypeNotFound
	}

	return nil
}
