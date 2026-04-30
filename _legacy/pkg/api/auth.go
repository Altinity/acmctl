package api

import "github.com/altinity/acmctl/pkg/models"

// Login authenticates with email/password and returns user data including token.
func (c *Client) Login(login, password string) (*models.User, error) {
	var user models.User
	err := c.Do("POST", "/login", map[string]string{
		"login":    login,
		"password": password,
	}, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// LoginVerify completes 2FA verification.
func (c *Client) LoginVerify(code, userID string) (*models.User, error) {
	var user models.User
	err := c.Do("POST", "/login/verify", map[string]string{
		"code": code,
		"user": userID,
	}, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Logout invalidates the current token.
func (c *Client) Logout() error {
	return c.Do("GET", "/logout", nil, nil)
}

// Probe is a lightweight health-check.
func (c *Client) Probe() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/probe", nil, &result)
	return result, err
}

// Auth0Connections lists available Auth0 connections.
func (c *Client) Auth0Connections() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/auth0/connections", nil, &result)
	return result, err
}

// LoginRecover sends a password reset email.
func (c *Client) LoginRecover(login string) error {
	return c.Do("POST", "/login/recover", map[string]string{"login": login}, nil)
}

// CheckResetToken verifies whether a password reset code is still active.
func (c *Client) CheckResetToken(code string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/login/reset", map[string]string{"code": code}, &result)
	return result, err
}

// ResetPassword finalizes a password reset using the emailed code.
func (c *Client) ResetPassword(code, password string) error {
	return c.Do("POST", "/login/reset", map[string]string{"code": code, "password": password}, nil)
}

// Signup creates a trial account.
func (c *Client) Signup(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/signup", params, &result)
	return result, err
}

// SignupEmail creates a trial account using only an email address.
func (c *Client) SignupEmail(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/signup-email", params, &result)
	return result, err
}

// CheckSignupToken verifies whether a signup confirmation code is still active.
func (c *Client) CheckSignupToken(code string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/signup/confirm", map[string]string{"code": code}, &result)
	return result, err
}

// SignupConfirm finishes a signup using the emailed code.
func (c *Client) SignupConfirm(code, password string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/signup/confirm", map[string]string{"code": code, "password": password}, &result)
	return result, err
}

// SingleAuthURL returns a target URL for Auth0 single-sign-on.
func (c *Client) SingleAuthURL(authType string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/singleauth", map[string]string{"type": authType}, &result)
	return result, err
}

// SingleAuth completes Auth0 single-sign-on with an authorization code.
func (c *Client) SingleAuth(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/singleauth", params, &result)
	return result, err
}

// AWSMarketplaceGateway handles landing for AWS Marketplace.
func (c *Client) AWSMarketplaceGateway(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/aws-marketplace-gateway", params, &result)
	return result, err
}

// AWSMarketplaceSub handles AWS Marketplace SNS subscriptions.
func (c *Client) AWSMarketplaceSub(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/aws-marketplace-sub", params, &result)
	return result, err
}

// GCPMarketplaceGateway handles landing for Google Marketplace.
func (c *Client) GCPMarketplaceGateway(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/gcp-marketplace-gateway", params, &result)
	return result, err
}
