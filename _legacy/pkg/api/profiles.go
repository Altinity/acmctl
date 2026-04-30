package api

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/models"
)

func (c *Client) ListProfileChOptions(profileID string) ([]models.ProfileSetting, error) {
	var settings []models.ProfileSetting
	err := c.Do("GET", fmt.Sprintf("/profile/%s/ch-options", profileID), nil, &settings)
	return settings, err
}

func (c *Client) ListProfileSettings(profileID string) ([]models.ProfileSetting, error) {
	var settings []models.ProfileSetting
	err := c.Do("GET", fmt.Sprintf("/profile/%s/settings", profileID), nil, &settings)
	return settings, err
}

func (c *Client) CreateProfileSetting(profileID string, params map[string]string) (*models.ProfileSetting, error) {
	var setting models.ProfileSetting
	err := c.DoForm("POST", fmt.Sprintf("/profile/%s/settings", profileID), params, &setting)
	if err != nil {
		return nil, err
	}
	return &setting, nil
}

func (c *Client) UpdateProfileSetting(id string, params map[string]string) (*models.ProfileSetting, error) {
	var setting models.ProfileSetting
	err := c.DoForm("POST", fmt.Sprintf("/setting/%s", id), params, &setting)
	if err != nil {
		return nil, err
	}
	return &setting, nil
}

func (c *Client) DeleteProfileSetting(id string) error {
	return c.Do("DELETE", fmt.Sprintf("/setting/%s", id), nil, nil)
}

func (c *Client) CreateClusterProfile(clusterID string, params map[string]string) (*models.Profile, error) {
	var profile models.Profile
	err := c.Do("POST", fmt.Sprintf("/cluster/%s/profiles", clusterID), params, &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (c *Client) UpdateProfile(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/profile/%s", id), params, &result)
	return result, err
}

func (c *Client) DeleteProfile(id string) error {
	return c.Do("DELETE", fmt.Sprintf("/profile/%s", id), nil, nil)
}
