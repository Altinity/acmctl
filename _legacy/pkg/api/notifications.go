package api

import "fmt"

func (c *Client) ListUserNotifications(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/n", params, &result)
	return result, err
}

func (c *Client) ListAdminNotifications(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/notifications", params, &result)
	return result, err
}

func (c *Client) CreateAdminNotification(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("POST", "/notifications", params, &result)
	return result, err
}

func (c *Client) GetNotification(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/notification/%s", id), nil, &result)
	return result, err
}

func (c *Client) UpdateNotification(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("POST", fmt.Sprintf("/notification/%s", id), params, &result)
	return result, err
}

func (c *Client) DeleteNotification(id string) error {
	return c.Do("DELETE", fmt.Sprintf("/notification/%s", id), nil, nil)
}

func (c *Client) AckNotification(id string) error {
	return c.Do("PUT", fmt.Sprintf("/notification/%s/ack", id), nil, nil)
}
