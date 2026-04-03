package models

type ClusterSetting struct {
	ID          string `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	Value       string `json:"value" yaml:"value"`
	ValueFrom   string `json:"valueFrom" yaml:"valueFrom"`
	Description string `json:"description" yaml:"description"`
}

func (s ClusterSetting) TableHeaders() []string {
	return []string{"ID", "NAME", "VALUE", "DESCRIPTION"}
}

func (s ClusterSetting) TableRow() []string {
	return []string{
		s.ID,
		s.Name,
		s.Value,
		s.Description,
	}
}

type ClusterEnvSetting struct {
	ID        string `json:"id" yaml:"id"`
	Name      string `json:"name" yaml:"name"`
	Value     string `json:"value" yaml:"value"`
	ValueFrom string `json:"valueFrom" yaml:"valueFrom"`
}

type Profile struct {
	ID          string `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
}

func (p Profile) TableHeaders() []string {
	return []string{"ID", "NAME", "DESCRIPTION"}
}

func (p Profile) TableRow() []string {
	return []string{
		p.ID,
		p.Name,
		p.Description,
	}
}

type ProfileSetting struct {
	ID          string `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	Value       string `json:"value" yaml:"value"`
	Description string `json:"description" yaml:"description"`
	Parsed      bool   `json:"parsed" yaml:"parsed"`
}
