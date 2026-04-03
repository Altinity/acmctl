package models

import "fmt"

type ActionLog struct {
	ID          string                 `json:"id" yaml:"id"`
	Created     string                 `json:"created" yaml:"created"`
	Updated     string                 `json:"updated" yaml:"updated"`
	Controller  string                 `json:"controller" yaml:"controller"`
	Action      string                 `json:"action" yaml:"action"`
	URL         string                 `json:"url" yaml:"url"`
	ClientIP    string                 `json:"clientIP" yaml:"clientIP"`
	Type        string                 `json:"type" yaml:"type"`
	State       map[string]interface{} `json:"state" yaml:"state"`
	Request     map[string]interface{} `json:"request" yaml:"request"`
	Success     bool                   `json:"success" yaml:"success"`
	Confirmed   bool                   `json:"confirmed" yaml:"confirmed"`
	Error       string                 `json:"error" yaml:"error"`
	UserInfo    map[string]interface{} `json:"userInfo" yaml:"userInfo"`
	ClusterInfo map[string]interface{} `json:"clusterInfo" yaml:"clusterInfo"`
}

func (a ActionLog) TableHeaders() []string {
	return []string{"ID", "CREATED", "ACTION", "SUCCESS", "TYPE"}
}

func (a ActionLog) TableRow() []string {
	return []string{
		a.ID,
		a.Created,
		a.Action,
		fmt.Sprintf("%v", a.Success),
		a.Type,
	}
}
