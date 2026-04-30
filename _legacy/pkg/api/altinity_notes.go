package api

import "fmt"

func (c *Client) ListAltinityNotes(envID string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/altinity-notes", envID), nil, &result)
	return result, err
}

func (c *Client) CreateAltinityNote(envID string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("POST", fmt.Sprintf("/environment/%s/altinity-notes", envID), params, &result)
	return result, err
}

func (c *Client) UpdateAltinityNote(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("POST", fmt.Sprintf("/altinity-note/%s", id), params, &result)
	return result, err
}

func (c *Client) DeleteAltinityNote(id string) error {
	return c.Do("DELETE", fmt.Sprintf("/altinity-note/%s", id), nil, nil)
}
