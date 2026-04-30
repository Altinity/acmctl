package api

import "github.com/altinity/acmctl/pkg/models"

// Login authenticates with email/password and returns user data
// including token.
//
// Credentials are sent as application/x-www-form-urlencoded body —
// never on the URL. Putting them on the URL (via Do's params) would
// leak the password into server access logs, reverse-proxy logs,
// referer headers, and verbose-mode stdout.
func (c *Client) Login(login, password string) (*models.User, error) {
	var user models.User
	err := c.DoForm("POST", "/login", map[string]string{
		"login":    login,
		"password": password,
	}, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
