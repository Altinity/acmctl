package api

import "fmt"

// Environment lifecycle / connection management

func (c *Client) ApproveEnvironment(id, reason string) error {
	params := map[string]string{}
	if reason != "" {
		params["reason"] = reason
	}
	return c.Do("POST", fmt.Sprintf("/environment/%s/approve", id), params, nil)
}

func (c *Client) AccCheck(id string, noWait bool) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if noWait {
		params["noWait"] = "true"
	}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/acc-check", id), params, &result)
	return result, err
}

func (c *Client) AccConnect(id, resources string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if resources != "" {
		params["resources"] = resources
	}
	err := c.Do("POST", fmt.Sprintf("/environment/%s/acc-connect", id), params, &result)
	return result, err
}

func (c *Client) GetAccToken(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/environment/%s/get-acc-token", id), nil, &result)
	return result, err
}

func (c *Client) ResetEnvironment(id string) error {
	return c.Do("PUT", fmt.Sprintf("/environment/%s/reset", id), nil, nil)
}

func (c *Client) RefreshKubeEnvironment(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/environment/%s/kube-update", id), nil, &result)
	return result, err
}

// Inspection

func (c *Client) ListEnvironmentAlerts(id, resolved string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if resolved != "" {
		params["resolved"] = resolved
	}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/alerts", id), params, &result)
	return result, err
}

func (c *Client) ListEnvironmentBuckets(id, bucketType string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if bucketType != "" {
		params["type"] = bucketType
	}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/buckets", id), params, &result)
	return result, err
}

func (c *Client) GetClusterLaunchValidity(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/cluster-launch-validity", id), nil, &result)
	return result, err
}

func (c *Client) ExportEnvironment(id, format string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if format != "" {
		params["format"] = format
	}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/export", id), params, &result)
	return result, err
}

func (c *Client) GetEnvironmentLog(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/log", id), params, &result)
	return result, err
}

func (c *Client) GetIcebergSettings(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/iceberg", id), nil, &result)
	return result, err
}

func (c *Client) GetInviteDetails(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/invite-details", id), nil, &result)
	return result, err
}

func (c *Client) InviteUser(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/environment/%s/invite", id), params, &result)
	return result, err
}

func (c *Client) GetEnvironmentResources(id, limits string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if limits != "" {
		params["limits"] = limits
	}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/resources", id), params, &result)
	return result, err
}

func (c *Client) GetEnvironmentResource(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/resource", id), params, &result)
	return result, err
}

func (c *Client) GetEnvironmentUsage(id, skipClusterID string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if skipClusterID != "" {
		params["skipClusterId"] = skipClusterID
	}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/usage", id), params, &result)
	return result, err
}

// CHOP / Kubernetes management

func (c *Client) ListChopConfigurations(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/chop-configurations", id), nil, &result)
	return result, err
}

func (c *Client) AddChopConfiguration(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("POST", fmt.Sprintf("/environment/%s/chop-configurations", id), params, &result)
	return result, err
}

func (c *Client) PatchChopConfiguration(id, name string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("PUT", fmt.Sprintf("/environment/%s/chop-configuration/%s", id, name), params, &result)
	return result, err
}

func (c *Client) DeleteChopConfiguration(id, name string) error {
	return c.Do("DELETE", fmt.Sprintf("/environment/%s/chop-configuration/%s", id, name), nil, nil)
}

func (c *Client) ApplyChopConfiguration(id string) error {
	return c.Do("PUT", fmt.Sprintf("/environment/%s/chop-configuration-apply", id), nil, nil)
}

func (c *Client) ApplyKubeConfig(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/environment/%s/config-apply", id), nil, &result)
	return result, err
}

func (c *Client) ListConfigurationTemplates(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/configuration-templates", id), nil, &result)
	return result, err
}

func (c *Client) PatchConfigurationTemplate(id, name string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("PUT", fmt.Sprintf("/environment/%s/configuration-template/%s", id, name), params, &result)
	return result, err
}

func (c *Client) ListKubeConfigmaps(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/kube-configmaps", id), nil, &result)
	return result, err
}

func (c *Client) PatchKubeConfigmap(id, name string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("PUT", fmt.Sprintf("/environment/%s/kube-configmap/%s", id, name), params, &result)
	return result, err
}

func (c *Client) GetKubeMap(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/kube-map", id), nil, &result)
	return result, err
}

func (c *Client) HandleKubeCho(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/environment/%s/kube-cho", id), params, &result)
	return result, err
}

// Discovery

func (c *Client) DiscoverEnvironment(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/environment/%s/discover", id), nil, &result)
	return result, err
}

func (c *Client) ConfirmDiscovery(id, clusters string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if clusters != "" {
		params["clusters"] = clusters
	}
	err := c.Do("POST", fmt.Sprintf("/environment/%s/discover", id), params, &result)
	return result, err
}

// Backups

func (c *Client) ListExternalBackups(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/backups", params, &result)
	return result, err
}

func (c *Client) ListEnvironmentBackups(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/environment/%s/backups", id), params, &result)
	return result, err
}

// Provisioning / connect / import

func (c *Client) ConnectEnvironment(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("POST", "/environments/connect", params, &result)
	return result, err
}

func (c *Client) ImportEnvironment(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("POST", "/environments/import", params, &result)
	return result, err
}

func (c *Client) RequestEnvironment(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/environments/request", params, &result)
	return result, err
}

// Cluster CRUD via environment

func (c *Client) AddClusterToEnvironment(envID string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/environment/%s/clusters", envID), params, &result)
	return result, err
}

func (c *Client) ImportCluster(envID string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("POST", fmt.Sprintf("/environment/%s/clusters/import", envID), params, &result)
	return result, err
}

func (c *Client) RestoreClusterIntoEnvironment(envID string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/environment/%s/clusters/restore", envID), params, &result)
	return result, err
}
