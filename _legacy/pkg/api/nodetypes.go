package api

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/models"
)

func (c *Client) CreateNodeType(envID string, params map[string]string) (*models.NodeType, error) {
	var nt models.NodeType
	err := c.Do("POST", fmt.Sprintf("/environment/%s/nodetypes", envID), params, &nt)
	if err != nil {
		return nil, err
	}
	return &nt, nil
}

func (c *Client) UpdateNodeType(id string, params map[string]string) (*models.NodeType, error) {
	var nt models.NodeType
	err := c.Do("POST", fmt.Sprintf("/nodetype/%s", id), params, &nt)
	if err != nil {
		return nil, err
	}
	return &nt, nil
}

func (c *Client) DeleteNodeType(id string) error {
	return c.Do("DELETE", fmt.Sprintf("/nodetype/%s", id), nil, nil)
}
