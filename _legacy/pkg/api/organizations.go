package api

import "fmt"

func (c *Client) ListOrganizations(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/organizations", params, &result)
	return result, err
}

func (c *Client) CreateOrganization(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/organizations", params, &result)
	return result, err
}

func (c *Client) GetOrganization(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/organization/%s", id), nil, &result)
	return result, err
}

func (c *Client) UpdateOrganization(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/organization/%s", id), params, &result)
	return result, err
}

func (c *Client) DeleteOrganization(id string, terminate bool) error {
	params := map[string]string{}
	if terminate {
		params["terminate"] = "1"
	}
	return c.Do("DELETE", fmt.Sprintf("/organization/%s", id), params, nil)
}

func (c *Client) UpdateOrganizationLogins(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/organization/%s/logins", id), params, &result)
	return result, err
}
