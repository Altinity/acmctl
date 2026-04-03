package models

type DbUser struct {
	ID               string      `json:"id" yaml:"id"`
	Login            string      `json:"login" yaml:"login"`
	Password         string      `json:"password,omitempty" yaml:"password,omitempty"`
	Networks         interface{} `json:"networks" yaml:"networks"`
	Databases        interface{} `json:"databases" yaml:"databases"`
	AccessManagement bool        `json:"accessManagement" yaml:"accessManagement"`
}

func (d DbUser) TableHeaders() []string {
	return []string{"ID", "LOGIN", "ACCESS MGMT"}
}

func (d DbUser) TableRow() []string {
	return []string{
		d.ID,
		d.Login,
		boolStr(d.AccessManagement),
	}
}

func boolStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
