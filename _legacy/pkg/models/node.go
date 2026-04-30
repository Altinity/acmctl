package models

import "fmt"

type Node struct {
	ID          string      `json:"id" yaml:"id"`
	Type        string      `json:"type" yaml:"type"`
	ContainerID interface{} `json:"containerId" yaml:"containerId"`
	VolumeID    interface{} `json:"volumeId" yaml:"volumeId"`
	Terminate   interface{} `json:"terminate" yaml:"terminate"`
	Status      string      `json:"status" yaml:"status"`
	State       string      `json:"state" yaml:"state"`
	Host        string      `json:"host" yaml:"host"`
	PublicHost  string      `json:"publicHost" yaml:"publicHost"`
	Port        interface{} `json:"port" yaml:"port"`
	PortTLS     interface{} `json:"portTLS" yaml:"portTLS"`
	HTTPPort    interface{} `json:"httpPort" yaml:"httpPort"`
	HTTPPortTLS interface{} `json:"httpPortTLS" yaml:"httpPortTLS"`
	User        string      `json:"user" yaml:"user"`
	Pass        string      `json:"pass,omitempty" yaml:"pass,omitempty"`
	Version     string      `json:"version" yaml:"version"`
	CheckTime   string      `json:"checkTime" yaml:"checkTime"`
	AwsOptions  interface{} `json:"awsOptions" yaml:"awsOptions"`
	DataPath    string      `json:"dataPath" yaml:"dataPath"`
	Metrics     interface{} `json:"metrics" yaml:"metrics"`
}

func (n Node) TableHeaders() []string {
	return []string{"ID", "TYPE", "STATE", "STATUS", "HOST", "PORT", "VERSION"}
}

func (n Node) TableRow() []string {
	return []string{
		n.ID,
		n.Type,
		n.State,
		n.Status,
		n.Host,
		fmt.Sprintf("%v", n.Port),
		n.Version,
	}
}

type NodeType struct {
	ID           string      `json:"id" yaml:"id"`
	Scope        string      `json:"scope" yaml:"scope"`
	Code         string      `json:"code" yaml:"code"`
	Name         string      `json:"name" yaml:"name"`
	IsSpot       bool        `json:"isSpot" yaml:"isSpot"`
	StorageClass string      `json:"storageClass" yaml:"storageClass"`
	ExtraSpec    interface{} `json:"extraSpec" yaml:"extraSpec"`
	Tolerations  interface{} `json:"tolerations" yaml:"tolerations"`
	NodeSelector interface{} `json:"nodeSelector" yaml:"nodeSelector"`
	Capacity     interface{} `json:"capacity" yaml:"capacity"`
	CPU          interface{} `json:"cpu" yaml:"cpu"`
	Memory       interface{} `json:"memory" yaml:"memory"`
}

func (n NodeType) TableHeaders() []string {
	return []string{"ID", "CODE", "NAME", "CPU", "MEMORY", "CAPACITY", "SPOT"}
}

func (n NodeType) TableRow() []string {
	return []string{
		n.ID,
		n.Code,
		n.Name,
		fmt.Sprintf("%v", n.CPU),
		fmt.Sprintf("%v", n.Memory),
		fmt.Sprintf("%v", n.Capacity),
		fmt.Sprintf("%v", n.IsSpot),
	}
}
