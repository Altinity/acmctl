package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const DefaultURL = "https://acm.altinity.cloud/api/"

type Config struct {
	URL    string `yaml:"url"`
	Token  string `yaml:"token"`
	Output string `yaml:"output"`
}

// DefaultPath returns ~/.acmctl.yaml. Returns an error rather than
// silently swallowing a missing $HOME — that condition normally only
// happens in pathological environments (locked-down containers,
// misconfigured CI runners), but the config command is the wrong
// place to silently degrade.
func DefaultPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolve home dir: %w", err)
	}
	return filepath.Join(home, ".acmctl.yaml"), nil
}

func Load(path string) (*Config, error) {
	cfg := &Config{
		URL:    DefaultURL,
		Output: "table",
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}
	// KnownFields(true) catches typos like `urls:` instead of `url:`
	// in the user's config — without it, an unknown key silently
	// leaves the struct field at its default and the user is left
	// wondering why their override didn't take effect.
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	if err := dec.Decode(cfg); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	if cfg.URL == "" {
		cfg.URL = DefaultURL
	}
	if cfg.Output == "" {
		cfg.Output = "table"
	}
	return cfg, nil
}

func Save(path string, cfg *Config) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}
