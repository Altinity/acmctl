package models

import "fmt"

type User struct {
	ID          string `json:"id" yaml:"id"`
	Email       string `json:"email" yaml:"email"`
	Name        string `json:"name" yaml:"name"`
	Token       string `json:"token,omitempty" yaml:"token,omitempty"`
	TokenExpiry string `json:"tokenExpiry,omitempty" yaml:"tokenExpiry,omitempty"`
	Origins     string `json:"origins" yaml:"origins"`
	Blocked     bool   `json:"blocked" yaml:"blocked"`
	DarkTheme   bool   `json:"darkTheme" yaml:"darkTheme"`
	LastLogin   string `json:"lastLogin" yaml:"lastLogin"`
}

type UserRole struct {
	ID     string                 `json:"id" yaml:"id"`
	Name   string                 `json:"name" yaml:"name"`
	Rights map[string]interface{} `json:"rights" yaml:"rights"`
}

func (u User) TableHeaders() []string {
	return []string{"ID", "EMAIL", "NAME", "BLOCKED", "LAST LOGIN"}
}

func (u User) TableRow() []string {
	return []string{
		u.ID,
		u.Email,
		u.Name,
		fmt.Sprintf("%v", u.Blocked),
		u.LastLogin,
	}
}
