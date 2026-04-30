package api

import "fmt"

func (c *Client) GetCloudOptions(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/cloud/options", params, &result)
	return result, err
}

func (c *Client) GetEnvironmentCloudOptions(envID string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cloud/%s/options", envID), params, &result)
	return result, err
}

func (c *Client) CloudQuery(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("POST", "/cloud/query", params, &result)
	return result, err
}
