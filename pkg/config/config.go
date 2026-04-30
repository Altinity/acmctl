package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

const DefaultURL = "https://acm.altinity.cloud/api/"

// Profile is one ACM environment's settings: API URL + token.
// Multiple profiles live under Config.Profiles, keyed by a
// user-chosen name (e.g., "prod", "dev", "stage").
type Profile struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token,omitempty"`
}

// Config is the on-disk shape of ~/.acmctl.yaml.
//
// Two layouts coexist for back-compat:
//
//  1. Profile-based (preferred):
//     default_profile: prod
//     profiles:
//       prod: { url: ..., token: ... }
//       dev:  { url: ..., token: ... }
//
//  2. Legacy flat (read-only on the way in; never written back):
//     url: ...
//     token: ...
//
// On Load, a legacy flat config is automatically promoted to a
// synthetic "default" profile in memory, so callers always see the
// profile-based view.
type Config struct {
	DefaultProfile string             `yaml:"default_profile,omitempty"`
	Profiles       map[string]Profile `yaml:"profiles,omitempty"`
	Output         string             `yaml:"output,omitempty"`

	// Legacy flat fields. Keep declared so KnownFields(true) doesn't
	// reject older configs. Promoted to a profile in Load().
	URL   string `yaml:"url,omitempty"`
	Token string `yaml:"token,omitempty"`
}

// DefaultPath returns ~/.acmctl.yaml. Returns an error rather than
// silently swallowing a missing $HOME — that condition normally only
// happens in pathological environments (locked-down containers,
// misconfigured CI runners).
func DefaultPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolve home dir: %w", err)
	}
	return filepath.Join(home, ".acmctl.yaml"), nil
}

func Load(path string) (*Config, error) {
	cfg := &Config{
		Profiles: map[string]Profile{},
		Output:   "table",
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	if err := dec.Decode(cfg); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	if cfg.Profiles == nil {
		cfg.Profiles = map[string]Profile{}
	}
	if cfg.Output == "" {
		cfg.Output = "table"
	}

	// Legacy flat config → synthetic "default" profile. We don't
	// rewrite the file unless the user mutates a profile (so existing
	// flat configs are left untouched until a Save() happens).
	if len(cfg.Profiles) == 0 && (cfg.URL != "" || cfg.Token != "") {
		cfg.Profiles["default"] = Profile{URL: cfg.URL, Token: cfg.Token}
		if cfg.DefaultProfile == "" {
			cfg.DefaultProfile = "default"
		}
	}

	return cfg, nil
}

// Save writes config to path. We always emit the profile-based
// layout — never the legacy flat fields.
func Save(path string, cfg *Config) error {
	out := struct {
		DefaultProfile string             `yaml:"default_profile,omitempty"`
		Profiles       map[string]Profile `yaml:"profiles,omitempty"`
		Output         string             `yaml:"output,omitempty"`
	}{
		DefaultProfile: cfg.DefaultProfile,
		Profiles:       cfg.Profiles,
		Output:         cfg.Output,
	}
	data, err := yaml.Marshal(out)
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o600)
}

// ActiveProfile resolves which profile to use, applying the
// resolution order:
//
//  1. flagOverride (--profile)
//  2. ACMCTL_PROFILE env var
//  3. cfg.DefaultProfile
//  4. if exactly one profile is defined, use it
//  5. error
//
// Returns the chosen profile, its name, and any error.
func (cfg *Config) ActiveProfile(flagOverride string) (Profile, string, error) {
	name := flagOverride
	if name == "" {
		name = os.Getenv("ACMCTL_PROFILE")
	}
	if name == "" {
		name = cfg.DefaultProfile
	}
	if name == "" && len(cfg.Profiles) == 1 {
		for k, p := range cfg.Profiles {
			return p, k, nil
		}
	}
	if name == "" {
		return Profile{}, "", fmt.Errorf("no profile selected — set --profile, ACMCTL_PROFILE env, or default_profile in config (`acmctl config use-profile <name>`)")
	}
	p, ok := cfg.Profiles[name]
	if !ok {
		return Profile{}, name, fmt.Errorf("profile %q not found (known: %s)", name, profileNames(cfg.Profiles))
	}
	return p, name, nil
}

// SetProfile replaces (or creates) the named profile.
func (cfg *Config) SetProfile(name string, p Profile) {
	if cfg.Profiles == nil {
		cfg.Profiles = map[string]Profile{}
	}
	cfg.Profiles[name] = p
}

// RemoveProfile deletes a profile by name. If the removed profile
// was the default, the default is cleared.
func (cfg *Config) RemoveProfile(name string) {
	delete(cfg.Profiles, name)
	if cfg.DefaultProfile == name {
		cfg.DefaultProfile = ""
	}
}

func profileNames(m map[string]Profile) string {
	if len(m) == 0 {
		return "<none>"
	}
	names := make([]string, 0, len(m))
	for k := range m {
		names = append(names, k)
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}
