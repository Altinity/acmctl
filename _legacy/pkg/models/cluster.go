package models

import "fmt"

type Cluster struct {
	ID                        string      `json:"id" yaml:"id"`
	Name                      string      `json:"name" yaml:"name"`
	Type                      interface{} `json:"type" yaml:"type"`
	AdminPass                 string      `json:"adminPass,omitempty" yaml:"adminPass,omitempty"`
	StartupTime               interface{} `json:"startupTime" yaml:"startupTime"`
	TroubleshootingMode       bool        `json:"troubleshootingMode" yaml:"troubleshootingMode"`
	Role                      string      `json:"role" yaml:"role"`
	BackupSource              interface{} `json:"backupSource" yaml:"backupSource"`
	ReplicaSource             interface{} `json:"replicaSource" yaml:"replicaSource"`
	AlertsEmail               string      `json:"alertsEmail" yaml:"alertsEmail"`
	AlertsSettings            interface{} `json:"alertsSettings" yaml:"alertsSettings"`
	DatadogSettings           interface{} `json:"datadogSettings" yaml:"datadogSettings"`
	SystemVersion             string      `json:"systemVersion" yaml:"systemVersion"`
	StoragePolicy             string      `json:"storagePolicy" yaml:"storagePolicy"`
	MoveFactor                interface{} `json:"moveFactor" yaml:"moveFactor"`
	ObjectStorages            interface{} `json:"objectStorages" yaml:"objectStorages"`
	SecretHandler             string      `json:"secretHandler" yaml:"secretHandler"`
	Nodes                     interface{} `json:"nodes" yaml:"nodes"`
	Replicas                  string      `json:"replicas" yaml:"replicas"`
	Shards                    string      `json:"shards" yaml:"shards"`
	Options                   interface{} `json:"options" yaml:"options"`
	BackupOptions             interface{} `json:"backupOptions" yaml:"backupOptions"`
	BackupConfigModifications interface{} `json:"backupConfigModifications" yaml:"backupConfigModifications"`
	Locked                    bool        `json:"locked" yaml:"locked"`
	Endpoint                  string      `json:"endpoint" yaml:"endpoint"`
	EndpointHTTP              string      `json:"endpointHTTP" yaml:"endpointHTTP"`
	AlternateEndpoints        interface{} `json:"alternateEndpoints" yaml:"alternateEndpoints"`
	EndpointsEnabled          interface{} `json:"endpointsEnabled" yaml:"endpointsEnabled"`
	Annotations               interface{} `json:"annotations" yaml:"annotations"`
	Secure                    bool        `json:"secure" yaml:"secure"`
	Status                    string      `json:"status" yaml:"status"`
	CheckTime                 string      `json:"checkTime" yaml:"checkTime"`
	LatestBackupTime          string      `json:"latestBackupTime" yaml:"latestBackupTime"`
	ChGuardSettings           interface{} `json:"chGuardSettings" yaml:"chGuardSettings"`
	CustomLBAnnotations       interface{} `json:"customLBAnnotations" yaml:"customLBAnnotations"`
	ZoneAwareness             bool        `json:"zoneAwareness" yaml:"zoneAwareness"`
	Uptime                    string      `json:"uptime" yaml:"uptime"`
	UptimeSettings            interface{} `json:"uptimeSettings" yaml:"uptimeSettings"`
	ActionSchedule            interface{} `json:"actionSchedule" yaml:"actionSchedule"`
	Timezone                  interface{} `json:"timezone" yaml:"timezone"`
	AltinitySupport           string      `json:"altinitySupport" yaml:"altinitySupport"`
	MysqlProtocol             bool        `json:"mysqlProtocol" yaml:"mysqlProtocol"`
	MysqlPort                 interface{} `json:"mysqlPort" yaml:"mysqlPort"`
	State                     string      `json:"state" yaml:"state"`
	// Extra fields from detailed response
	IDEnvironment   string      `json:"id_environment,omitempty" yaml:"id_environment,omitempty"`
	IDZookeeper     string      `json:"id_zookeeper,omitempty" yaml:"id_zookeeper,omitempty"`
	Version         string      `json:"version,omitempty" yaml:"version,omitempty"`
	NormalizedName  string      `json:"normalizedName,omitempty" yaml:"normalizedName,omitempty"`
	Environment     interface{} `json:"environment,omitempty" yaml:"environment,omitempty"`
	Owner           interface{} `json:"owner,omitempty" yaml:"owner,omitempty"`
	Layouts         interface{} `json:"layouts,omitempty" yaml:"layouts,omitempty"`
	Endpoints       interface{} `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
	Storage         interface{} `json:"storage,omitempty" yaml:"storage,omitempty"`
}

func (c Cluster) TableHeaders() []string {
	return []string{"ID", "NAME", "STATE", "STATUS", "VERSION", "ENDPOINT"}
}

func (c Cluster) TableRow() []string {
	v := c.Version
	if v == "" {
		v = c.SystemVersion
	}
	return []string{
		c.ID,
		c.Name,
		c.State,
		c.Status,
		v,
		c.Endpoint,
	}
}

func (c Cluster) String() string {
	return fmt.Sprintf("%s (%s)", c.Name, c.ID)
}
