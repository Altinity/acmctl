package config

import (
	"os"
	"path/filepath"
	"testing"
)

func writeFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	p := filepath.Join(dir, name)
	if err := os.WriteFile(p, []byte(content), 0o600); err != nil {
		t.Fatalf("write %s: %v", p, err)
	}
	return p
}

func TestLoad_NoFile(t *testing.T) {
	cfg, err := Load(filepath.Join(t.TempDir(), "missing.yaml"))
	if err != nil {
		t.Fatalf("Load missing: %v", err)
	}
	if cfg == nil {
		t.Fatal("Load missing returned nil cfg")
	}
	if len(cfg.Profiles) != 0 || cfg.DefaultProfile != "" {
		t.Errorf("missing file should yield empty profiles, got %+v", cfg)
	}
	if cfg.Output != "table" {
		t.Errorf("default output = %q, want table", cfg.Output)
	}
}

func TestLoad_LegacyFlatPromotion(t *testing.T) {
	dir := t.TempDir()
	p := writeFile(t, dir, "acmctl.yaml", `
url: https://acm.altinity.cloud/api
token: oldtok123
`)
	cfg, err := Load(p)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.DefaultProfile != "default" {
		t.Errorf("DefaultProfile = %q, want default", cfg.DefaultProfile)
	}
	got, ok := cfg.Profiles["default"]
	if !ok {
		t.Fatalf("synthetic 'default' profile missing")
	}
	if got.URL != "https://acm.altinity.cloud/api" || got.Token != "oldtok123" {
		t.Errorf("legacy promotion got %+v", got)
	}
}

func TestLoad_ProfileBased(t *testing.T) {
	dir := t.TempDir()
	p := writeFile(t, dir, "acmctl.yaml", `
default_profile: dev
profiles:
  prod:
    url: https://acm.altinity.cloud/api
    token: prod-tok
  dev:
    url: https://acm.dev.altinity.cloud/api
    token: dev-tok
`)
	cfg, err := Load(p)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.DefaultProfile != "dev" {
		t.Errorf("DefaultProfile = %q", cfg.DefaultProfile)
	}
	if len(cfg.Profiles) != 2 {
		t.Errorf("len(Profiles) = %d, want 2", len(cfg.Profiles))
	}
}

func TestLoad_KnownFieldsRejectsTypos(t *testing.T) {
	dir := t.TempDir()
	p := writeFile(t, dir, "acmctl.yaml", `
default_profile: prod
profiless:                     # typo
  prod: {url: https://x}
`)
	_, err := Load(p)
	if err == nil {
		t.Fatal("expected error on unknown YAML field, got nil")
	}
}

func TestActiveProfile_FlagOverride(t *testing.T) {
	cfg := &Config{
		DefaultProfile: "prod",
		Profiles: map[string]Profile{
			"prod": {URL: "https://prod"},
			"dev":  {URL: "https://dev"},
		},
	}
	t.Setenv("ACMCTL_PROFILE", "")
	p, name, err := cfg.ActiveProfile("dev")
	if err != nil {
		t.Fatal(err)
	}
	if name != "dev" || p.URL != "https://dev" {
		t.Errorf("got name=%q url=%q, want dev/https://dev", name, p.URL)
	}
}

func TestActiveProfile_EnvOverride(t *testing.T) {
	cfg := &Config{
		DefaultProfile: "prod",
		Profiles: map[string]Profile{
			"prod":  {URL: "https://prod"},
			"stage": {URL: "https://stage"},
		},
	}
	t.Setenv("ACMCTL_PROFILE", "stage")
	p, name, err := cfg.ActiveProfile("")
	if err != nil {
		t.Fatal(err)
	}
	if name != "stage" || p.URL != "https://stage" {
		t.Errorf("got name=%q url=%q, want stage/https://stage", name, p.URL)
	}
}

func TestActiveProfile_DefaultFallback(t *testing.T) {
	cfg := &Config{
		DefaultProfile: "prod",
		Profiles:       map[string]Profile{"prod": {URL: "https://prod"}},
	}
	t.Setenv("ACMCTL_PROFILE", "")
	p, name, err := cfg.ActiveProfile("")
	if err != nil {
		t.Fatal(err)
	}
	if name != "prod" || p.URL != "https://prod" {
		t.Errorf("got name=%q url=%q", name, p.URL)
	}
}

func TestActiveProfile_SingleProfileFallback(t *testing.T) {
	cfg := &Config{
		Profiles: map[string]Profile{"only": {URL: "https://x"}},
	}
	t.Setenv("ACMCTL_PROFILE", "")
	p, name, err := cfg.ActiveProfile("")
	if err != nil {
		t.Fatal(err)
	}
	if name != "only" || p.URL != "https://x" {
		t.Errorf("got name=%q url=%q", name, p.URL)
	}
}

func TestActiveProfile_NoneSelected(t *testing.T) {
	cfg := &Config{
		Profiles: map[string]Profile{
			"prod": {URL: "https://prod"},
			"dev":  {URL: "https://dev"},
		},
	}
	t.Setenv("ACMCTL_PROFILE", "")
	if _, _, err := cfg.ActiveProfile(""); err == nil {
		t.Error("expected error when no profile selected and 2+ exist, got nil")
	}
}

func TestActiveProfile_UnknownName(t *testing.T) {
	cfg := &Config{
		Profiles: map[string]Profile{"prod": {URL: "https://prod"}},
	}
	t.Setenv("ACMCTL_PROFILE", "")
	if _, _, err := cfg.ActiveProfile("missing"); err == nil {
		t.Error("expected error for unknown profile name, got nil")
	}
}

func TestSaveLoad_RoundTrip_DropsLegacyFields(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "acmctl.yaml")
	cfg := &Config{
		DefaultProfile: "prod",
		Profiles: map[string]Profile{
			"prod": {URL: "https://prod", Token: "t1"},
			"dev":  {URL: "https://dev"},
		},
		Output: "json",
		// Set legacy fields too — Save should NOT emit them.
		URL:   "https://legacy",
		Token: "legacy-tok",
	}
	if err := Save(p, cfg); err != nil {
		t.Fatal(err)
	}
	data, _ := os.ReadFile(p)
	if got := string(data); contains(got, "https://legacy") || contains(got, "legacy-tok") {
		t.Errorf("Save emitted legacy fields:\n%s", got)
	}

	round, err := Load(p)
	if err != nil {
		t.Fatal(err)
	}
	if round.DefaultProfile != "prod" || len(round.Profiles) != 2 || round.Output != "json" {
		t.Errorf("round-trip lost data: %+v", round)
	}
}

func TestSetRemoveProfile(t *testing.T) {
	cfg := &Config{Profiles: map[string]Profile{}}
	cfg.SetProfile("a", Profile{URL: "https://a"})
	cfg.SetProfile("b", Profile{URL: "https://b"})
	cfg.DefaultProfile = "a"
	if len(cfg.Profiles) != 2 {
		t.Errorf("got %d profiles", len(cfg.Profiles))
	}
	cfg.RemoveProfile("a")
	if _, ok := cfg.Profiles["a"]; ok {
		t.Error("removed profile still present")
	}
	if cfg.DefaultProfile != "" {
		t.Error("removing default should clear DefaultProfile")
	}
}

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
