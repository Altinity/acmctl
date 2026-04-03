package api

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/models"
)

func (c *Client) ListClusterSettings(clusterID string) ([]models.ClusterSetting, error) {
	var settings []models.ClusterSetting
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/settings", clusterID), nil, &settings)
	return settings, err
}

func (c *Client) CreateClusterSetting(clusterID string, params map[string]string) (*models.ClusterSetting, error) {
	var setting models.ClusterSetting
	err := c.Do("POST", fmt.Sprintf("/cluster/%s/settings", clusterID), params, &setting)
	if err != nil {
		return nil, err
	}
	return &setting, nil
}

func (c *Client) UpdateClusterSetting(id string, params map[string]string) (*models.ClusterSetting, error) {
	var setting models.ClusterSetting
	err := c.Do("POST", fmt.Sprintf("/cluster-setting/%s", id), params, &setting)
	if err != nil {
		return nil, err
	}
	return &setting, nil
}

func (c *Client) DeleteClusterSetting(id string) error {
	return c.Do("DELETE", fmt.Sprintf("/cluster-setting/%s", id), nil, nil)
}

func (c *Client) ListClusterEnvSettings(clusterID string) ([]models.ClusterEnvSetting, error) {
	var settings []models.ClusterEnvSetting
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/env-settings", clusterID), nil, &settings)
	return settings, err
}

func (c *Client) ListClusterProfiles(clusterID string) ([]models.Profile, error) {
	var profiles []models.Profile
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/profiles", clusterID), nil, &profiles)
	return profiles, err
}
