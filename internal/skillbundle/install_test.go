package skillbundle

import (
	"os"
	"path/filepath"
	"testing"
)

// makeBundle returns a minimal Bundle for tests with two files at
// known paths. Tests can mutate the returned map to simulate
// upstream changes.
func makeBundle() *Bundle {
	return &Bundle{
		Skill: "altinity-cloud",
		Files: map[string][]byte{
			"SKILL.md":             []byte("# skill v1\n"),
			"endpoints/clusters.md": []byte("# clusters\n"),
		},
	}
}

func newAgent() Agent {
	return Agent{Name: "claude", InstallSubdir: ".claude/skills"}
}

func TestInstall_FreshTarget(t *testing.T) {
	root := t.TempDir()
	b := makeBundle()
	s, err := Install(b, InstallOpts{ScopeRoot: root, Agent: newAgent()})
	if err != nil {
		t.Fatalf("Install: %v", err)
	}
	if s.Created != 2 || s.Updated != 0 || s.Unchanged != 0 || len(s.Conflicts) != 0 {
		t.Fatalf("unexpected summary: %+v", s)
	}
	got, _ := os.ReadFile(filepath.Join(root, ".claude/skills/altinity-cloud/SKILL.md"))
	if string(got) != "# skill v1\n" {
		t.Errorf("SKILL.md content = %q, want %q", got, "# skill v1\n")
	}
}

func TestInstall_Idempotent(t *testing.T) {
	root := t.TempDir()
	b := makeBundle()
	if _, err := Install(b, InstallOpts{ScopeRoot: root, Agent: newAgent()}); err != nil {
		t.Fatal(err)
	}
	s, err := Install(b, InstallOpts{ScopeRoot: root, Agent: newAgent()})
	if err != nil {
		t.Fatal(err)
	}
	if s.Created != 0 || s.Updated != 0 || s.Unchanged != 2 {
		t.Errorf("re-install should be all unchanged, got %+v", s)
	}
}

func TestInstall_ConflictWithoutForce(t *testing.T) {
	root := t.TempDir()
	b := makeBundle()
	if _, err := Install(b, InstallOpts{ScopeRoot: root, Agent: newAgent()}); err != nil {
		t.Fatal(err)
	}
	// Local edit: change one file, leave the other.
	dst := filepath.Join(root, ".claude/skills/altinity-cloud/SKILL.md")
	if err := os.WriteFile(dst, []byte("# tampered\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	s, err := Install(b, InstallOpts{ScopeRoot: root, Agent: newAgent()})
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Conflicts) != 1 {
		t.Errorf("expected 1 conflict, got %d (files: %+v)", len(s.Conflicts), s.Files)
	}
	// Tampered file should NOT have been overwritten without --force.
	got, _ := os.ReadFile(dst)
	if string(got) != "# tampered\n" {
		t.Errorf("file overwritten unexpectedly: %q", got)
	}
}

func TestInstall_ForceOverwrites(t *testing.T) {
	root := t.TempDir()
	b := makeBundle()
	if _, err := Install(b, InstallOpts{ScopeRoot: root, Agent: newAgent()}); err != nil {
		t.Fatal(err)
	}
	dst := filepath.Join(root, ".claude/skills/altinity-cloud/SKILL.md")
	if err := os.WriteFile(dst, []byte("# tampered\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	s, err := Install(b, InstallOpts{ScopeRoot: root, Agent: newAgent(), Force: true})
	if err != nil {
		t.Fatal(err)
	}
	if s.Updated != 1 || s.Unchanged != 1 || len(s.Conflicts) != 0 {
		t.Errorf("--force summary = %+v, want Updated=1 Unchanged=1", s)
	}
	got, _ := os.ReadFile(dst)
	if string(got) != "# skill v1\n" {
		t.Errorf("file not restored: %q", got)
	}
}

func TestInstall_DirSymlinkReplaced(t *testing.T) {
	root := t.TempDir()
	// Pre-create a symlink at the skill-dir level pointing at an
	// arbitrary external dir (simulates the iso-acm sandbox layout).
	src := t.TempDir()
	if err := os.WriteFile(filepath.Join(src, "SKILL.md"), []byte("# external\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	skillsDir := filepath.Join(root, ".claude/skills")
	if err := os.MkdirAll(skillsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(skillsDir, "altinity-cloud")
	if err := os.Symlink(src, link); err != nil {
		t.Fatal(err)
	}

	b := makeBundle()
	s, err := Install(b, InstallOpts{ScopeRoot: root, Agent: newAgent()})
	if err != nil {
		t.Fatal(err)
	}
	if s.Created != 2 {
		t.Errorf("expected 2 created (symlink replaced), got %+v", s)
	}

	// Symlink should have become a real directory.
	li, err := os.Lstat(link)
	if err != nil {
		t.Fatalf("lstat: %v", err)
	}
	if li.Mode()&os.ModeSymlink != 0 {
		t.Errorf("path is still a symlink after install")
	}
	if !li.IsDir() {
		t.Errorf("path is not a directory after install: mode=%v", li.Mode())
	}

	// External target must NOT have been clobbered.
	got, _ := os.ReadFile(filepath.Join(src, "SKILL.md"))
	if string(got) != "# external\n" {
		t.Errorf("symlink target was clobbered: %q", got)
	}
}

func TestInstall_DirSymlinkDryRun(t *testing.T) {
	root := t.TempDir()
	src := t.TempDir()
	if err := os.WriteFile(filepath.Join(src, "SKILL.md"), []byte("# skill v1\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(src, "endpoints", "clusters.md"), []byte("# clusters\n"), 0o644); err != nil {
		// Need to mkdir first.
		if err := os.MkdirAll(filepath.Join(src, "endpoints"), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(src, "endpoints", "clusters.md"), []byte("# clusters\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	skillsDir := filepath.Join(root, ".claude/skills")
	if err := os.MkdirAll(skillsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(skillsDir, "altinity-cloud")
	if err := os.Symlink(src, link); err != nil {
		t.Fatal(err)
	}

	b := makeBundle()
	s, err := Install(b, InstallOpts{ScopeRoot: root, Agent: newAgent(), DryRun: true})
	if err != nil {
		t.Fatal(err)
	}
	// Even though the symlink target's contents happen to match the
	// bundle byte-for-byte, dry-run must report "would replace" /
	// "created" — not "unchanged" — because a real run will replace
	// the symlink. The previous bug reported all-unchanged here.
	if s.Created != 2 || s.Unchanged != 0 {
		t.Errorf("dry-run with dir symlink: got %+v, want Created=2 Unchanged=0", s)
	}
	// Symlink must still be intact after dry-run.
	li, _ := os.Lstat(link)
	if li.Mode()&os.ModeSymlink == 0 {
		t.Errorf("dry-run modified the symlink")
	}
}

func TestInstall_FileSymlinkReplaced(t *testing.T) {
	root := t.TempDir()
	dst := filepath.Join(root, ".claude/skills/altinity-cloud/SKILL.md")
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		t.Fatal(err)
	}
	external := filepath.Join(t.TempDir(), "external.md")
	if err := os.WriteFile(external, []byte("# external\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(external, dst); err != nil {
		t.Fatal(err)
	}

	b := makeBundle()
	if _, err := Install(b, InstallOpts{ScopeRoot: root, Agent: newAgent()}); err != nil {
		t.Fatal(err)
	}

	// File-level symlink replaced.
	li, err := os.Lstat(dst)
	if err != nil {
		t.Fatal(err)
	}
	if li.Mode()&os.ModeSymlink != 0 {
		t.Errorf("file-level symlink not replaced")
	}
	got, _ := os.ReadFile(dst)
	if string(got) != "# skill v1\n" {
		t.Errorf("content = %q, want %q", got, "# skill v1\n")
	}
	// External target untouched.
	gotExt, _ := os.ReadFile(external)
	if string(gotExt) != "# external\n" {
		t.Errorf("symlink target clobbered: %q", gotExt)
	}
}

func TestInstall_BundleVersionDrift(t *testing.T) {
	root := t.TempDir()
	b := makeBundle()
	if _, err := Install(b, InstallOpts{ScopeRoot: root, Agent: newAgent()}); err != nil {
		t.Fatal(err)
	}
	// Simulate upstream change: bundle now has different content.
	b.Files["SKILL.md"] = []byte("# skill v2\n")

	// Without force: conflict.
	s, err := Install(b, InstallOpts{ScopeRoot: root, Agent: newAgent()})
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Conflicts) != 1 {
		t.Errorf("expected 1 conflict for upstream drift, got %+v", s)
	}

	// With force (mimics `update`): apply.
	s, err = Install(b, InstallOpts{ScopeRoot: root, Agent: newAgent(), Force: true})
	if err != nil {
		t.Fatal(err)
	}
	if s.Updated != 1 || s.Unchanged != 1 {
		t.Errorf("force-update summary = %+v, want Updated=1 Unchanged=1", s)
	}
	got, _ := os.ReadFile(filepath.Join(root, ".claude/skills/altinity-cloud/SKILL.md"))
	if string(got) != "# skill v2\n" {
		t.Errorf("content = %q, want v2", got)
	}
}

func TestIsInstalled(t *testing.T) {
	root := t.TempDir()
	a := newAgent()
	if IsInstalled(root, a, "altinity-cloud") {
		t.Error("empty root: should not be installed")
	}

	b := makeBundle()
	if _, err := Install(b, InstallOpts{ScopeRoot: root, Agent: a}); err != nil {
		t.Fatal(err)
	}
	if !IsInstalled(root, a, "altinity-cloud") {
		t.Error("after install: should report installed")
	}

	if IsInstalled(root, Agent{Name: "codex", InstallSubdir: ".codex/skills"}, "altinity-cloud") {
		t.Error("different agent: should not be installed")
	}
}

func TestSelectAgents(t *testing.T) {
	cases := []struct {
		name    string
		names   []string
		all     bool
		want    []string // expected agent names in order
		wantErr bool
	}{
		{"default", nil, false, []string{"claude"}, false},
		{"all", nil, true, []string{"claude", "codex"}, false},
		{"explicit one", []string{"codex"}, false, []string{"codex"}, false},
		{"explicit two", []string{"codex", "claude"}, false, []string{"codex", "claude"}, false},
		{"dedupe", []string{"claude", "claude"}, false, []string{"claude"}, false},
		{"unknown", []string{"gemini"}, false, nil, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SelectAgents(tc.names, tc.all)
			if (err != nil) != tc.wantErr {
				t.Fatalf("err=%v wantErr=%v", err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if len(got) != len(tc.want) {
				t.Fatalf("len(got)=%d, want %d", len(got), len(tc.want))
			}
			for i, w := range tc.want {
				if got[i].Name != w {
					t.Errorf("got[%d]=%s, want %s", i, got[i].Name, w)
				}
			}
		})
	}
}

func TestScopeRoot(t *testing.T) {
	cases := []struct {
		scope   string
		wantErr bool
	}{
		{"global", false},
		{"project", false},
		{"wrong", true},
		{"", true},
	}
	for _, tc := range cases {
		t.Run(tc.scope, func(t *testing.T) {
			got, err := ScopeRoot(tc.scope)
			if (err != nil) != tc.wantErr {
				t.Fatalf("err=%v wantErr=%v", err, tc.wantErr)
			}
			if !tc.wantErr && got == "" {
				t.Errorf("got empty path with no error")
			}
		})
	}
}
