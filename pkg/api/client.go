package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/altinity/acmctl/pkg/models"
)

type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
	Verbose    bool
}

func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseURL: strings.TrimRight(baseURL, "/"),
		Token:   token,
		HTTPClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

// Do executes an API request. Params are sent as query parameters for all methods.
// If result is non-nil, the response "data" field is unmarshaled into it.
func (c *Client) Do(method, path string, params map[string]string, result interface{}) error {
	u, err := url.Parse(c.BaseURL + "/" + strings.TrimLeft(path, "/"))
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	q := u.Query()
	for k, v := range params {
		if v != "" {
			q.Set(k, v)
		}
	}
	u.RawQuery = q.Encode()

	if c.Verbose {
		fmt.Printf(">> %s %s\n", method, u.String())
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	if c.Token != "" {
		req.Header.Set("X-Auth-Token", c.Token)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if c.Verbose {
		fmt.Printf("<< HTTP %d (%d bytes)\n", resp.StatusCode, len(body))
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return &APIError{StatusCode: 401, Message: "authentication required — run 'acmctl login' or set --token"}
	}

	// Handle empty response body
	trimmed := strings.TrimSpace(string(body))
	if trimmed == "" || trimmed == "null" {
		if resp.StatusCode >= 400 {
			return &APIError{StatusCode: resp.StatusCode, Message: fmt.Sprintf("HTTP %d (empty response)", resp.StatusCode)}
		}
		return nil
	}

	var envelope models.APIResponse
	if err := json.Unmarshal(body, &envelope); err != nil {
		// Response is not in envelope format — try direct unmarshal or return error
		if resp.StatusCode >= 400 {
			return &APIError{StatusCode: resp.StatusCode, Message: trimmed}
		}
		// Try to unmarshal directly into result for non-envelope responses
		if result != nil {
			if directErr := json.Unmarshal(body, result); directErr == nil {
				return nil
			}
		}
		return fmt.Errorf("failed to parse response: %w", err)
	}

	if envelope.Error != nil && *envelope.Error != "" {
		return &APIError{StatusCode: resp.StatusCode, Message: *envelope.Error}
	}

	if resp.StatusCode >= 400 {
		return &APIError{StatusCode: resp.StatusCode, Message: trimmed}
	}

	if result != nil && envelope.Data != nil {
		if err := json.Unmarshal(envelope.Data, result); err != nil {
			return fmt.Errorf("failed to parse response data: %w", err)
		}
	}

	return nil
}
