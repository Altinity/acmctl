package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
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

// Do executes a request with optional query params. If result is non-nil, the
// response "data" envelope is unmarshaled into it.
func (c *Client) Do(method, path string, params map[string]string, result interface{}) error {
	u, err := c.url(path, params)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, u, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	c.setAuth(req)
	return c.send(req, result)
}

// DoForm sends params as application/x-www-form-urlencoded.
func (c *Client) DoForm(method, path string, body map[string]string, result interface{}) error {
	u, err := c.url(path, nil)
	if err != nil {
		return err
	}
	form := url.Values{}
	for k, v := range body {
		form.Set(k, v)
	}
	req, err := http.NewRequest(method, u, strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	c.setAuth(req)
	return c.send(req, result)
}

// DoJSON sends the body as application/json. body may be nil for no payload.
func (c *Client) DoJSON(method, path string, body []byte, result interface{}) error {
	u, err := c.url(path, nil)
	if err != nil {
		return err
	}
	var r io.Reader
	if len(body) > 0 {
		r = bytes.NewReader(body)
	}
	req, err := http.NewRequest(method, u, r)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	if len(body) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	c.setAuth(req)
	return c.send(req, result)
}

func (c *Client) url(path string, params map[string]string) (string, error) {
	u, err := url.Parse(c.BaseURL + "/" + strings.TrimLeft(path, "/"))
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}
	if len(params) > 0 {
		q := u.Query()
		for k, v := range params {
			if v != "" {
				q.Set(k, v)
			}
		}
		u.RawQuery = q.Encode()
	}
	return u.String(), nil
}

func (c *Client) setAuth(req *http.Request) {
	if c.Token != "" {
		req.Header.Set("X-Auth-Token", c.Token)
	}
}

func (c *Client) send(req *http.Request, result interface{}) error {
	if c.Verbose {
		fmt.Printf(">> %s %s\n", req.Method, req.URL.String())
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if c.Verbose {
		fmt.Printf("<< HTTP %d (%d bytes)\n", resp.StatusCode, len(body))
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return &APIError{StatusCode: 401, Message: "authentication required — run 'acmctl login' or set ACMCTL_TOKEN"}
	}

	trimmed := strings.TrimSpace(string(body))
	if trimmed == "" || trimmed == "null" {
		if resp.StatusCode >= 400 {
			return &APIError{StatusCode: resp.StatusCode, Message: fmt.Sprintf("HTTP %d (empty response)", resp.StatusCode)}
		}
		return nil
	}

	// Try envelope format first: {"data": ..., "error": ...}
	var envelope struct {
		Data  json.RawMessage `json:"data"`
		Error *string         `json:"error"`
	}
	if err := json.Unmarshal(body, &envelope); err == nil {
		if envelope.Error != nil && *envelope.Error != "" {
			return &APIError{StatusCode: resp.StatusCode, Message: *envelope.Error}
		}
		if resp.StatusCode >= 400 {
			return &APIError{StatusCode: resp.StatusCode, Message: trimmed}
		}
		if result != nil && envelope.Data != nil {
			return json.Unmarshal(envelope.Data, result)
		}
		return nil
	}

	// Non-envelope response
	if resp.StatusCode >= 400 {
		return &APIError{StatusCode: resp.StatusCode, Message: trimmed}
	}
	if result != nil {
		return json.Unmarshal(body, result)
	}
	return nil
}
