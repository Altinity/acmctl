package models

// User is the response shape of POST /login. Only Token is
// load-bearing for the CLI today (saved to ~/.acmctl.yaml); the
// other fields are kept for completeness and to surface a
// human-readable login confirmation message.
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
