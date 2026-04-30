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

func (c *Client) UpdateOwnAccount(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/account", params, &result)
	return result, err
}

func (c *Client) GetAccessRights() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/access-rights", nil, &result)
	return result, err
}

func (c *Client) ListAccountRoles() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/account-roles", nil, &result)
	return result, err
}

func (c *Client) CreateAccountRole(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/account-roles", params, &result)
	return result, err
}

func (c *Client) UpdateAccountRole(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/account-role/%s", id), params, &result)
	return result, err
}

func (c *Client) DeleteAccountRole(id string) error {
	return c.Do("DELETE", fmt.Sprintf("/account-role/%s", id), nil, nil)
}

func (c *Client) GenerateAccountToken() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/account/token", nil, &result)
	return result, err
}

func (c *Client) GenerateAnywhereToken() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/account/anywhere-token", nil, &result)
	return result, err
}

func (c *Client) GetAccountLog(user string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/account/%s/log", user), params, &result)
	return result, err
}

func (c *Client) SetAccountAccess(id string, params map[string]string) error {
	return c.Do("POST", fmt.Sprintf("/account/%s/access", id), params, nil)
}
