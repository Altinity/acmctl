package api

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/models"
)

func (c *Client) AddZookeeperCluster(envID string, params map[string]string) (*models.ZookeeperCluster, error) {
	var zk models.ZookeeperCluster
	err := c.Do("POST", fmt.Sprintf("/environment/%s/zookeepers", envID), params, &zk)
	if err != nil {
		return nil, err
	}
	return &zk, nil
}

func (c *Client) LaunchZookeeperCluster(envID string, params map[string]string) (*models.ZookeeperCluster, error) {
	var zk models.ZookeeperCluster
	err := c.Do("POST", fmt.Sprintf("/environment/%s/zookeepers/launch", envID), params, &zk)
	if err != nil {
		return nil, err
	}
	return &zk, nil
}

func (c *Client) UpdateZookeeperCluster(id string, params map[string]string) (*models.ZookeeperCluster, error) {
	var zk models.ZookeeperCluster
	err := c.Do("POST", fmt.Sprintf("/zookeeper/%s", id), params, &zk)
	if err != nil {
		return nil, err
	}
	return &zk, nil
}

func (c *Client) DeleteZookeeperCluster(id string) error {
	return c.Do("DELETE", fmt.Sprintf("/zookeeper/%s", id), nil, nil)
}

func (c *Client) PushZookeeperCluster(id string) error {
	return c.Do("PUT", fmt.Sprintf("/zookeeper/%s/push", id), nil, nil)
}

func (c *Client) RescaleZookeeperCluster(id string, params map[string]string) error {
	return c.Do("PUT", fmt.Sprintf("/zookeeper/%s/rescale", id), params, nil)
}

func (c *Client) GetZookeeperStatus(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/zookeeper/%s/status", id), nil, &result)
	return result, err
}
