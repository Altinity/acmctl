package api

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/models"
)

func (c *Client) ListDbUsers(clusterID string) ([]models.DbUser, error) {
	var users []models.DbUser
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/users", clusterID), nil, &users)
	return users, err
}

func (c *Client) CreateDbUser(clusterID string, params map[string]string) (*models.DbUser, error) {
	var user models.DbUser
	err := c.Do("POST", fmt.Sprintf("/cluster/%s/users", clusterID), params, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) UpdateDbUser(clusterID, id string, params map[string]string) (*models.DbUser, error) {
	var user models.DbUser
	err := c.Do("POST", fmt.Sprintf("/cluster/%s/user/%s", clusterID, id), params, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) DeleteDbUser(clusterID, id string) error {
	return c.Do("DELETE", fmt.Sprintf("/cluster/%s/user/%s", clusterID, id), nil, nil)
}

// UpdateDbUserByID modifies a cluster user via the cluster-less /user/{id} path.
func (c *Client) UpdateDbUserByID(id string, params map[string]string) (*models.DbUser, error) {
	var user models.DbUser
	err := c.Do("POST", fmt.Sprintf("/user/%s", id), params, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteDbUserByID removes a cluster user via the cluster-less /user/{id} path.
func (c *Client) DeleteDbUserByID(id string) error {
	return c.Do("DELETE", fmt.Sprintf("/user/%s", id), nil, nil)
}
