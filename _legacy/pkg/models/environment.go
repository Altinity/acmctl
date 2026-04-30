package models

type Environment struct {
	ID                string                 `json:"id" yaml:"id"`
	Name              string                 `json:"name" yaml:"name"`
	NormalizedName    string                 `json:"normalizedName" yaml:"normalizedName"`
	DisplayName       string                 `json:"displayName" yaml:"displayName"`
	Created           string                 `json:"created" yaml:"created"`
	Type              string                 `json:"type" yaml:"type"`
	Domain            string                 `json:"domain" yaml:"domain"`
	State             string                 `json:"state" yaml:"state"`
	Status            string                 `json:"status" yaml:"status"`
	Managed           bool                   `json:"managed" yaml:"managed"`
	Remote            bool                   `json:"remote" yaml:"remote"`
	HostedByAltinity  bool                   `json:"hostedByAltinity" yaml:"hostedByAltinity"`
	KubeProvider      string                 `json:"kubeProvider" yaml:"kubeProvider"`
	KubeNamespace     string                 `json:"kubeNamespace" yaml:"kubeNamespace"`
	KubeAPIUrl        string                 `json:"kubeAPIUrl" yaml:"kubeAPIUrl"`
	KubeCHOVersion    string                 `json:"kubeCHOVersion" yaml:"kubeCHOVersion"`
	Monitoring        bool                   `json:"monitoring" yaml:"monitoring"`
	MonitoringUrl     string                 `json:"monitoringUrl" yaml:"monitoringUrl"`
	DashboardUrl      string                 `json:"dashboardUrl" yaml:"dashboardUrl"`
	ImageRegistry     string                 `json:"imageRegistry" yaml:"imageRegistry"`
	Notes             string                 `json:"notes" yaml:"notes"`
	SniProxyForCH     bool                   `json:"sniProxyForCH" yaml:"sniProxyForCH"`
	PublicLB          bool                   `json:"publicLB" yaml:"publicLB"`
	PrivateLB         bool                   `json:"privateLB" yaml:"privateLB"`
	VpcEndpoints      bool                   `json:"vpcEndpoints" yaml:"vpcEndpoints"`
	BackupOptions     map[string]interface{} `json:"backupOptions" yaml:"backupOptions"`
	Options           map[string]interface{} `json:"options" yaml:"options"`
	ResourceLimits    map[string]interface{} `json:"resourceLimits" yaml:"resourceLimits"`
}

func (e Environment) TableHeaders() []string {
	return []string{"ID", "NAME", "TYPE", "STATE", "STATUS", "PROVIDER", "DOMAIN"}
}

func (e Environment) TableRow() []string {
	return []string{
		e.ID,
		e.Name,
		e.Type,
		e.State,
		e.Status,
		e.KubeProvider,
		e.Domain,
	}
}
