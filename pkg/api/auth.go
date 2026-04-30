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
