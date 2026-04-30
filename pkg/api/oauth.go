package api

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/altinity/acmctl/pkg/models"
)

// OAuth flow constants. Keep in sync with the Auth0 application
// "ACM" (client_id below) on the altinity.auth0.com tenant.
//
// IMPORTANT — Auth0 setup required: each port in callbackPorts must
// be registered in the Auth0 client's "Allowed Callback URLs" list
// as `http://localhost:<port>/cb`. Auth0 doesn't honor RFC 8252's
// wildcard-port suggestion; redirect_uri matching is exact.
const (
	auth0Domain    = "altinity.auth0.com"
	// Native Auth0 application dedicated to acmctl (PKCE-only,
	// no client_secret, no use of refresh tokens for now).
	// Distinct from the web UI's "Regular Web Application" client.
	acmClientID    = "BjVCqpdbU4Z6zYpxHPXG1zU4N6QCIlGX"
	callbackPath   = "/cb"
	defaultScope   = "openid profile email"
	flowTimeout    = 5 * time.Minute
	exchangeAPIURL = "/singleauth"
)

// callbackPorts are the loopback ports the CLI tries in order.
// Multiple are listed so a port-conflict on one machine doesn't
// require an Auth0 dashboard change. All must be pre-registered
// in Auth0 as `http://localhost:<port>/cb`.
var callbackPorts = []int{49152, 49153, 49154}

// OAuthResult is what OAuthLogin returns on success.
type OAuthResult struct {
	// SessionToken is the ACM session token issued by /singleauth.
	// It can be saved to ~/.acmctl.yaml and used as X-Auth-Token on
	// subsequent requests.
	SessionToken string
	// User describes the authenticated principal, when ACM returns
	// it. Useful for the CLI's confirmation message.
	User *models.User
}

// OAuthLogin runs the PKCE-with-loopback OAuth flow against
// Auth0/ACM and returns an ACM session token. Blocks until the
// user completes the flow in the browser, or until flowTimeout
// elapses, or until ctx is cancelled.
//
// openBrowser=true tries to spawn the user's default browser; if
// the spawn fails (or is disabled), the CLI prints the URL for
// manual paste.
func (c *Client) OAuthLogin(ctx context.Context, openBrowser bool) (*OAuthResult, error) {
	// Bind one of the registered ports.
	ln, port, err := bindCallbackPort()
	if err != nil {
		return nil, err
	}
	defer ln.Close()
	redirectURI := fmt.Sprintf("http://localhost:%d%s", port, callbackPath)

	// PKCE state.
	verifier, err := generateCodeVerifier()
	if err != nil {
		return nil, fmt.Errorf("generate code_verifier: %w", err)
	}
	challenge := codeChallengeS256(verifier)
	state, err := randomBase64URL(24)
	if err != nil {
		return nil, fmt.Errorf("generate state: %w", err)
	}

	authURL := buildAuthorizeURL(redirectURI, challenge, state)

	// One-shot HTTP server: handles a single /cb request, then
	// signals via the channel.
	type cbResult struct {
		code, state string
		err         error
	}
	resultCh := make(chan cbResult, 1)
	server := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != callbackPath {
				http.NotFound(w, r)
				return
			}
			q := r.URL.Query()

			// Auth0 may return error=...&error_description=... on cancel.
			if errStr := q.Get("error"); errStr != "" {
				desc := q.Get("error_description")
				html(w, "Login failed", fmt.Sprintf("<p><b>%s</b></p><p>%s</p>", htmlEscape(errStr), htmlEscape(desc)))
				resultCh <- cbResult{err: fmt.Errorf("auth0 error: %s: %s", errStr, desc)}
				return
			}

			gotState := q.Get("state")
			if gotState != state {
				html(w, "Login failed", "<p>State parameter mismatch — possible CSRF.</p>")
				resultCh <- cbResult{err: fmt.Errorf("state mismatch")}
				return
			}
			code := q.Get("code")
			if code == "" {
				html(w, "Login failed", "<p>No authorization code in callback.</p>")
				resultCh <- cbResult{err: fmt.Errorf("missing code in callback")}
				return
			}

			html(w, "Logged in", "<p>You can close this tab and return to the terminal.</p>")
			resultCh <- cbResult{code: code, state: gotState}
		}),
	}
	go func() { _ = server.Serve(ln) }()
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_ = server.Shutdown(shutdownCtx)
	}()

	// Open browser (best effort) and tell the user where to go.
	fmt.Fprintf(os.Stderr, "Opening browser for ACM login. If nothing opens, visit:\n  %s\n", authURL)
	if openBrowser {
		_ = openInBrowser(authURL)
	}

	// Wait for the callback or timeout.
	var cb cbResult
	select {
	case cb = <-resultCh:
	case <-time.After(flowTimeout):
		return nil, fmt.Errorf("timed out waiting for browser callback after %s", flowTimeout)
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	if cb.err != nil {
		return nil, cb.err
	}

	// Exchange the authorization code for an Auth0 id_token.
	idToken, err := c.exchangeAuthCode(ctx, cb.code, verifier, redirectURI)
	if err != nil {
		return nil, fmt.Errorf("exchange auth code with Auth0: %w", err)
	}

	// Trade the id_token for an ACM session token.
	return c.exchangeIDToken(ctx, idToken, cb.code, cb.state)
}

// bindCallbackPort tries each port in callbackPorts until one is
// available, returning the listener and the bound port.
func bindCallbackPort() (net.Listener, int, error) {
	var lastErr error
	for _, port := range callbackPorts {
		ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
		if err == nil {
			return ln, port, nil
		}
		lastErr = err
	}
	return nil, 0, fmt.Errorf("none of the registered callback ports (%v) are free: %w", callbackPorts, lastErr)
}

func generateCodeVerifier() (string, error) {
	// 32 bytes of randomness → 43 base64url chars; meets RFC 7636
	// length recommendation (43–128 chars, no padding).
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func codeChallengeS256(verifier string) string {
	h := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(h[:])
}

func randomBase64URL(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func buildAuthorizeURL(redirectURI, challenge, state string) string {
	q := url.Values{}
	q.Set("response_type", "code")
	q.Set("client_id", acmClientID)
	q.Set("redirect_uri", redirectURI)
	q.Set("scope", defaultScope)
	q.Set("state", state)
	q.Set("code_challenge", challenge)
	q.Set("code_challenge_method", "S256")
	return fmt.Sprintf("https://%s/authorize?%s", auth0Domain, q.Encode())
}

// exchangeAuthCode does the PKCE token exchange directly with
// Auth0 (the ACM backend isn't involved at this step). Returns
// the id_token, which carries the authenticated user's identity
// signed by Auth0.
func (c *Client) exchangeAuthCode(ctx context.Context, code, verifier, redirectURI string) (string, error) {
	body := url.Values{}
	body.Set("grant_type", "authorization_code")
	body.Set("client_id", acmClientID)
	body.Set("code", code)
	body.Set("redirect_uri", redirectURI)
	body.Set("code_verifier", verifier)

	endpoint := fmt.Sprintf("https://%s/oauth/token", auth0Domain)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(body.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	if c.Verbose {
		fmt.Fprintf(os.Stderr, ">> POST %s (PKCE token exchange)\n", endpoint)
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var parsed struct {
		IDToken     string `json:"id_token"`
		AccessToken string `json:"access_token"`
		Error       string `json:"error"`
		ErrorDesc   string `json:"error_description"`
	}
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return "", fmt.Errorf("decode token response (HTTP %d): %w", resp.StatusCode, err)
	}
	if parsed.Error != "" {
		return "", fmt.Errorf("auth0 %s: %s", parsed.Error, parsed.ErrorDesc)
	}
	if parsed.IDToken == "" {
		return "", errors.New("auth0 returned no id_token")
	}
	return parsed.IDToken, nil
}

// exchangeIDToken posts the Auth0 id_token to ACM's /singleauth
// endpoint to mint an ACM session token. Tries variants in order:
//
//  1. {token: <id_token>}                     — token-only path
//  2. {token: <id_token>, code, state}        — full PKCE shape
//
// /singleauth was originally designed for ACM's web flow (where ACM
// initiates the OAuth and remembers state server-side). For a
// CLI-initiated PKCE flow with a Native client, the token-only path
// is the only one that can work without backend cooperation. We try
// it first, fall back to the full shape, and surface whichever
// error is more informative.
func (c *Client) exchangeIDToken(ctx context.Context, idToken, code, state string) (*OAuthResult, error) {
	attempts := []map[string]string{
		{"token": idToken},
		{"token": idToken, "code": code, "state": state},
	}
	var lastErr error
	for i, form := range attempts {
		var user models.User
		err := c.DoForm(http.MethodPost, exchangeAPIURL, form, &user)
		if err == nil && user.Token != "" {
			return &OAuthResult{SessionToken: user.Token, User: &user}, nil
		}
		if err == nil {
			err = fmt.Errorf("/singleauth attempt %d returned no token", i+1)
		}
		lastErr = err
	}
	return nil, fmt.Errorf("ACM /singleauth: %w", lastErr)
}

// openInBrowser launches the platform's default URL handler.
// Best-effort; errors are non-fatal — the URL is also printed.
func openInBrowser(rawURL string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", rawURL).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", rawURL).Start()
	default:
		return exec.Command("xdg-open", rawURL).Start()
	}
}

// html writes a small landing page back to the browser after the
// OAuth dance. Title is the <title> + heading; body is raw HTML
// (caller must escape user-controlled input via htmlEscape).
func html(w http.ResponseWriter, title, body string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<!doctype html><html><head><meta charset="utf-8"><title>%s — acmctl</title>
<style>body{font:14px/1.5 -apple-system,system-ui,sans-serif;max-width:480px;margin:48px auto;padding:0 16px;color:#111}h1{font-size:22px}p{color:#444}</style>
</head><body><h1>%s</h1>%s</body></html>`, htmlEscape(title), htmlEscape(title), body)
}

// htmlEscape is a minimal escaper for the small landing page; not a
// general-purpose tool. Five chars cover the relevant injection vectors
// for the strings we control here.
func htmlEscape(s string) string {
	r := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		`"`, "&quot;",
		"'", "&#39;",
	)
	return r.Replace(s)
}
