package api

func (c *Client) ListDatasets() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/datasets", nil, &result)
	return result, err
}

func (c *Client) GetMetrics() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/metrics", nil, &result)
	return result, err
}

func (c *Client) GetReferenceSpec() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/reference.json", nil, &result)
	return result, err
}

func (c *Client) GetStatus() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/status", nil, &result)
	return result, err
}
