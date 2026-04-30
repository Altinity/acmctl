package api

import "fmt"

func (c *Client) GetAuditReport(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/audit-report/%s", id), nil, &result)
	return result, err
}

func (c *Client) ListClusterAuditReports(clusterID string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/audit-reports", clusterID), nil, &result)
	return result, err
}

func (c *Client) CreateClusterAuditReport(clusterID string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/cluster/%s/audit-reports", clusterID), nil, &result)
	return result, err
}
