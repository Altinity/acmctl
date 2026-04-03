package models

import "fmt"

type ClusterVolume struct {
	ID         string      `json:"id" yaml:"id"`
	Name       string      `json:"name" yaml:"name"`
	Class      string      `json:"class" yaml:"class"`
	Type       string      `json:"type" yaml:"type"`
	Size       interface{} `json:"size" yaml:"size"`
	Throughput interface{} `json:"throughput" yaml:"throughput"`
	IOPS       interface{} `json:"iops" yaml:"iops"`
	IsCordoned bool        `json:"isCordoned" yaml:"isCordoned"`
	IsRemoved  bool        `json:"isRemoved" yaml:"isRemoved"`
}

func (v ClusterVolume) TableHeaders() []string {
	return []string{"ID", "NAME", "CLASS", "TYPE", "SIZE", "IOPS"}
}

func (v ClusterVolume) TableRow() []string {
	return []string{
		v.ID,
		v.Name,
		v.Class,
		v.Type,
		fmt.Sprintf("%v", v.Size),
		fmt.Sprintf("%v", v.IOPS),
	}
}
