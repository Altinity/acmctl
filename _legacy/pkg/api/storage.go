package api

import "fmt"

func (c *Client) AddObjectStorage(clusterID string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/cluster/%s/object-storages", clusterID), params, &result)
	return result, err
}

func (c *Client) RemoveObjectStorage(clusterID, name string) error {
	return c.Do("DELETE", fmt.Sprintf("/cluster/%s/object-storages/%s", clusterID, name), nil, nil)
}

func (c *Client) ListClusterVolumes(clusterID string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/storage/cluster/%s/volumes", clusterID), nil, &result)
	return result, err
}

func (c *Client) ModifyClusterVolumes(clusterID string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/storage/cluster/%s/volumes", clusterID), params, &result)
	return result, err
}

func (c *Client) UpdateClusterVolume(volumeID string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("PATCH", fmt.Sprintf("/storage/cluster-volume/%s", volumeID), params, &result)
	return result, err
}

func (c *Client) DeleteClusterVolume(volumeID string) error {
	return c.Do("DELETE", fmt.Sprintf("/storage/cluster-volume/%s", volumeID), nil, nil)
}

func (c *Client) CordonClusterVolume(volumeID string, cordon string) error {
	params := map[string]string{}
	if cordon != "" {
		params["cordon"] = cordon
	}
	return c.Do("POST", fmt.Sprintf("/storage/cluster-volume/%s/cordon", volumeID), params, nil)
}

func (c *Client) FreeClusterVolume(volumeID string) error {
	return c.Do("POST", fmt.Sprintf("/storage/cluster-volume/%s/free", volumeID), nil, nil)
}

func (c *Client) ValidateClusterVolumeModification(volumeID string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/storage/cluster-volume/%s/validate-modification-pvc", volumeID), nil, &result)
	return result, err
}

func (c *Client) InterruptFreeVolumes(clusterID string) error {
	return c.Do("POST", fmt.Sprintf("/storage/cluster/%s/interrupt-free", clusterID), nil, nil)
}
