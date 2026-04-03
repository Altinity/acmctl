package api

import (
	"encoding/json"
	"fmt"

	"github.com/altinity/acmctl/pkg/models"
)

func (c *Client) GetNode(id string) (*models.Node, error) {
	var node models.Node
	err := c.Do("GET", fmt.Sprintf("/node/%s", id), nil, &node)
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func (c *Client) RestartNode(id string, hard bool) error {
	params := map[string]string{}
	if hard {
		params["hard"] = "true"
	}
	return c.Do("PUT", fmt.Sprintf("/node/%s/restart", id), params, nil)
}

func (c *Client) GetNodeStatus(id string) (json.RawMessage, error) {
	var result json.RawMessage
	err := c.Do("GET", fmt.Sprintf("/node/%s/status", id), nil, &result)
	return result, err
}

func (c *Client) GetNodeMetrics(id string, detailed bool) (json.RawMessage, error) {
	params := map[string]string{}
	if detailed {
		params["detailed"] = "true"
	}
	var result json.RawMessage
	err := c.Do("GET", fmt.Sprintf("/node/%s/metrics", id), params, &result)
	return result, err
}
