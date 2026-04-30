package api

import "fmt"

func (c *Client) GetConsoleInfo() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/console/info", nil, &result)
	return result, err
}

func (c *Client) GetConsoleLogs(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/console/logs", params, &result)
	return result, err
}

func (c *Client) GetEnvironmentConsoleLogs(envID string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/console/logs/%s", envID), params, &result)
	return result, err
}

func (c *Client) GetFilteredConsoleLogs(labels string) (interface{}, error) {
	var result interface{}
	q := map[string]string{}
	if labels != "" {
		q["labels"] = labels
	}
	err := c.Do("GET", "/console/logs/filtered", q, &result)
	return result, err
}

func (c *Client) GetConsoleSettings() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/console/settings", nil, &result)
	return result, err
}

func (c *Client) UpdateConsoleSettings(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("PATCH", "/console/settings", params, &result)
	return result, err
}

func (c *Client) ListConsoleTasks(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/console/tasks", params, &result)
	return result, err
}

func (c *Client) ScheduleConsoleTask(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("POST", "/console/tasks", params, &result)
	return result, err
}

func (c *Client) UpdateConsoleTask(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/console/task/%s", id), params, &result)
	return result, err
}

func (c *Client) DeleteConsoleTask(id string) error {
	return c.Do("DELETE", fmt.Sprintf("/console/task/%s", id), nil, nil)
}
