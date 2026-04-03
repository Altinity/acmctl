package api

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/models"
)

func (c *Client) ListAccounts() ([]models.User, error) {
	var users []models.User
	err := c.Do("GET", "/accounts", nil, &users)
	return users, err
}

func (c *Client) GetAccount() (*models.User, error) {
	var user models.User
	err := c.Do("GET", "/account", nil, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) CreateAccount(params map[string]string) (*models.User, error) {
	var user models.User
	err := c.Do("POST", "/accounts", params, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) UpdateAccount(id string, params map[string]string) (*models.User, error) {
	var user models.User
	err := c.Do("POST", fmt.Sprintf("/account/%s", id), params, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) DeleteAccount(id string) error {
	return c.Do("DELETE", fmt.Sprintf("/account/%s", id), nil, nil)
}
