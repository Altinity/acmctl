package models

import "fmt"

type ZookeeperCluster struct {
	ID        string `json:"id" yaml:"id"`
	Tag       string `json:"tag" yaml:"tag"`
	Launched  bool   `json:"launched" yaml:"launched"`
	Dedicated bool   `json:"dedicated" yaml:"dedicated"`
	State     string `json:"state" yaml:"state"`
	Size      string `json:"size" yaml:"size"`
	Suffix    string `json:"suffix" yaml:"suffix"`
}

func (z ZookeeperCluster) TableHeaders() []string {
	return []string{"ID", "TAG", "STATE", "SIZE", "LAUNCHED", "DEDICATED"}
}

func (z ZookeeperCluster) TableRow() []string {
	return []string{
		z.ID,
		z.Tag,
		z.State,
		z.Size,
		fmt.Sprintf("%v", z.Launched),
		fmt.Sprintf("%v", z.Dedicated),
	}
}

type Zookeeper struct {
	ID          string      `json:"id" yaml:"id"`
	Index       interface{} `json:"index" yaml:"index"`
	Host        string      `json:"host" yaml:"host"`
	Port        interface{} `json:"port" yaml:"port"`
	Type        string      `json:"type" yaml:"type"`
	ContainerID interface{} `json:"containerId" yaml:"containerId"`
	State       string      `json:"state" yaml:"state"`
	Terminate   interface{} `json:"terminate" yaml:"terminate"`
	Options     interface{} `json:"options" yaml:"options"`
}
