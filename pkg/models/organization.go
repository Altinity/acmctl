package models

import "fmt"

type Organization struct {
	ID                 string                 `json:"id" yaml:"id"`
	Name               string                 `json:"name" yaml:"name"`
	CompanyName        string                 `json:"companyName" yaml:"companyName"`
	EmailDomain        string                 `json:"emailDomain" yaml:"emailDomain"`
	InheritEnvironment bool                   `json:"inheritEnvironment" yaml:"inheritEnvironment"`
	Opened             bool                   `json:"opened" yaml:"opened"`
	Limited            bool                   `json:"limited" yaml:"limited"`
	Blocked            bool                   `json:"blocked" yaml:"blocked"`
	BlockedPassword    bool                   `json:"blockedPassword" yaml:"blockedPassword"`
	BlockedAPI         bool                   `json:"blockedAPI" yaml:"blockedAPI"`
	AllowAdminPassword bool                   `json:"allowAdminPassword" yaml:"allowAdminPassword"`
	Enable2FA          bool                   `json:"enable2FA" yaml:"enable2FA"`
	TrialExpiry        string                 `json:"trialExpiry" yaml:"trialExpiry"`
	SupportPlan        string                 `json:"supportPlan" yaml:"supportPlan"`
	AutoCharge         bool                   `json:"autoCharge" yaml:"autoCharge"`
	OktaSettings       map[string]interface{} `json:"oktaSettings" yaml:"oktaSettings"`
	Quotas             map[string]interface{} `json:"quotas" yaml:"quotas"`
}

func (o Organization) TableHeaders() []string {
	return []string{"ID", "NAME", "COMPANY", "BLOCKED", "SUPPORT PLAN"}
}

func (o Organization) TableRow() []string {
	return []string{
		o.ID,
		o.Name,
		o.CompanyName,
		fmt.Sprintf("%v", o.Blocked),
		o.SupportPlan,
	}
}
