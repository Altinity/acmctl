package api

import (
	"encoding/json"
	"fmt"

	"github.com/altinity/acmctl/pkg/models"
)

func (c *Client) ListClusters() ([]models.Cluster, error) {
	var clusters []models.Cluster
	err := c.Do("GET", "/clusters", nil, &clusters)
	return clusters, err
}

func (c *Client) GetCluster(id string) (*models.Cluster, error) {
	var cluster models.Cluster
	err := c.Do("GET", fmt.Sprintf("/cluster/%s", id), nil, &cluster)
	if err != nil {
		return nil, err
	}
	return &cluster, nil
}

func (c *Client) UpdateCluster(id string, params map[string]string) (*models.Cluster, error) {
	var cluster models.Cluster
	err := c.Do("POST", fmt.Sprintf("/cluster/%s", id), params, &cluster)
	if err != nil {
		return nil, err
	}
	return &cluster, nil
}

func (c *Client) DeleteCluster(id string, terminate bool) error {
	t := "0"
	if terminate {
		t = "1"
	}
	return c.Do("DELETE", fmt.Sprintf("/cluster/%s/%s", id, t), nil, nil)
}

func (c *Client) LaunchCluster(envID string, params map[string]string) (*models.Cluster, error) {
	var cluster models.Cluster
	err := c.Do("POST", fmt.Sprintf("/environment/%s/clusters/launch", envID), params, &cluster)
	if err != nil {
		return nil, err
	}
	return &cluster, nil
}

func (c *Client) RestartCluster(id string, method string) error {
	params := map[string]string{}
	if method != "" {
		params["method"] = method
	}
	return c.Do("PUT", fmt.Sprintf("/cluster/%s/restart", id), params, nil)
}

func (c *Client) StopCluster(id string) error {
	return c.Do("PUT", fmt.Sprintf("/cluster/%s/stop", id), nil, nil)
}

func (c *Client) ResumeCluster(id string, params map[string]string) error {
	return c.Do("PUT", fmt.Sprintf("/cluster/%s/resume", id), params, nil)
}

func (c *Client) UpgradeCluster(id string, params map[string]string) error {
	return c.Do("PUT", fmt.Sprintf("/cluster/%s/upgrade", id), params, nil)
}

func (c *Client) RescaleCluster(id string, params map[string]string) error {
	return c.Do("PUT", fmt.Sprintf("/cluster/%s/rescale", id), params, nil)
}

func (c *Client) BackupCluster(id string) error {
	return c.Do("POST", fmt.Sprintf("/cluster/%s/backup", id), nil, nil)
}

func (c *Client) RestoreCluster(id string, params map[string]string) error {
	return c.Do("POST", fmt.Sprintf("/cluster/%s/restore", id), params, nil)
}

func (c *Client) QueryCluster(id string, params map[string]string) (json.RawMessage, error) {
	var result json.RawMessage
	err := c.Do("POST", fmt.Sprintf("/cluster/%s/query", id), params, &result)
	return result, err
}

func (c *Client) ListClusterNodes(id string) ([]models.Node, error) {
	var nodes []models.Node
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/nodes", id), nil, &nodes)
	return nodes, err
}

func (c *Client) GetClusterStatus(id string) (json.RawMessage, error) {
	var result json.RawMessage
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/status", id), nil, &result)
	return result, err
}

func (c *Client) ListClusterBackups(id string) (json.RawMessage, error) {
	var result json.RawMessage
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/backups", id), nil, &result)
	return result, err
}

func (c *Client) GetClusterLogs(clusterID string, params map[string]string) (json.RawMessage, error) {
	var result json.RawMessage
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/logs", clusterID), params, &result)
	return result, err
}

func (c *Client) PushCluster(id string) error {
	return c.Do("POST", fmt.Sprintf("/cluster/%s/push", id), nil, nil)
}
