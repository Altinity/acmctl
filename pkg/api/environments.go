package api

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/models"
)

func (c *Client) ListEnvironments() ([]models.Environment, error) {
	var envs []models.Environment
	err := c.Do("GET", "/environments", nil, &envs)
	return envs, err
}

func (c *Client) GetEnvironment(id string) (*models.Environment, error) {
	var env models.Environment
	err := c.Do("GET", fmt.Sprintf("/environment/%s", id), nil, &env)
	if err != nil {
		return nil, err
	}
	return &env, nil
}

func (c *Client) UpdateEnvironment(id string, params map[string]string) (*models.Environment, error) {
	var env models.Environment
	err := c.Do("POST", fmt.Sprintf("/environment/%s", id), params, &env)
	if err != nil {
		return nil, err
	}
	return &env, nil
}

func (c *Client) DeleteEnvironment(id string) error {
	return c.Do("DELETE", fmt.Sprintf("/environment/%s", id), nil, nil)
}

func (c *Client) ListEnvironmentClusters(envID string) ([]models.Cluster, error) {
	var clusters []models.Cluster
	err := c.Do("GET", fmt.Sprintf("/environment/%s/clusters", envID), nil, &clusters)
	return clusters, err
}

func (c *Client) ListEnvironmentNodeTypes(envID string) ([]models.NodeType, error) {
	var nodeTypes []models.NodeType
	err := c.Do("GET", fmt.Sprintf("/environment/%s/nodetypes", envID), nil, &nodeTypes)
	return nodeTypes, err
}

func (c *Client) ListEnvironmentZookeepers(envID string) ([]models.ZookeeperCluster, error) {
	var zks []models.ZookeeperCluster
	err := c.Do("GET", fmt.Sprintf("/environment/%s/zookeepers", envID), nil, &zks)
	return zks, err
}
