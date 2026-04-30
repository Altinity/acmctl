package skillbundle

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

// Status enumerates per-file outcomes when installing a bundle.
type Status string

const (
	StatusCreated         Status = "created"
	StatusUpdated         Status = "updated"
	StatusUnchanged       Status = "unchanged"
	StatusWouldOverwrite  Status = "would_overwrite"
	StatusReplacedSymlink Status = "replaced_symlink"
)

// FileChange records a single (target-path, outcome) for the
// install summary.
type FileChange struct {
	Path   string
	Status Status
}

// Summary aggregates the result of an Install call. Conflicts is
// the set of files where contents differ from the bundle and
// neither --force nor update mode permitted overwriting; when
// non-empty, callers (e.g., `install` without --force) should
// abort and surface the list to the user.
type Summary struct {
	Created   int
	Updated   int
	Unchanged int
	Conflicts []string
	Files     []FileChange
}

// InstallOpts controls the behavior of an Install call.
type InstallOpts struct {
	// ScopeRoot is the parent dir under which the agent's subdir
	// lives. Typically $HOME (--scope global) or $CWD (--scope project).
	ScopeRoot string
	// Agent specifies the install path (e.g., ".claude/skills").
	Agent Agent
	// Force overwrites existing files even if they differ from the
	// bundle. Set by --force or by `update` mode.
	Force bool
	// DryRun stops just before any write — useful for preview.
	DryRun bool
}

// Install copies every file in the bundle into
//
//	{ScopeRoot}/{Agent.InstallSubdir}/{bundle.Skill}/<rel>
//
// idempotently. Files identical to the bundle are reported as
// Unchanged. Files that exist with different contents are reported
// as WouldOverwrite (recorded in Summary.Conflicts) unless Force is
// set, in which case they are Updated. Missing files are Created.
//
// When the existing target path is a symlink, it is removed before
// writing — never followed/written-through, so a sandbox setup
// pointing at a workspace source tree won't be clobbered. Symlink
// removals are recorded as ReplacedSymlink in the per-file list and
// counted toward Created (they didn't exist as files before).
func Install(b *Bundle, opts InstallOpts) (Summary, error) {
	if b == nil {
		return Summary{}, errors.New("nil bundle")
	}
	if opts.ScopeRoot == "" {
		return Summary{}, errors.New("scope root is empty")
	}
	dstRoot := filepath.Join(opts.ScopeRoot, opts.Agent.InstallSubdir, b.Skill)

	// Skill-directory-level symlink guard. If the skill directory
	// itself is a symlink (a common sandbox setup pointing at a
	// workspace source tree), per-file writes would resolve through
	// it and clobber the link's target. Remove the symlink before
	// touching anything below; the file loop will then create a
	// real directory and write fresh contents.
	if li, err := os.Lstat(dstRoot); err == nil && li.Mode()&fs.ModeSymlink != 0 {
		if !opts.DryRun {
			if rerr := os.Remove(dstRoot); rerr != nil {
				return Summary{}, fmt.Errorf("remove symlink at skill dir %s: %w", dstRoot, rerr)
			}
		}
	}

	// Sort for deterministic output.
	rels := make([]string, 0, len(b.Files))
	for k := range b.Files {
		rels = append(rels, k)
	}
	sort.Strings(rels)

	var s Summary
	for _, rel := range rels {
		dst := filepath.Join(dstRoot, rel)
		src := b.Files[rel]

		// Symlink guard before any write; never follow.
		replacedSymlink := false
		if li, err := os.Lstat(dst); err == nil && li.Mode()&fs.ModeSymlink != 0 {
			if !opts.DryRun {
				if rerr := os.Remove(dst); rerr != nil {
					return s, fmt.Errorf("remove symlink %s: %w", dst, rerr)
				}
			}
			replacedSymlink = true
		}

		existing, err := os.ReadFile(dst)
		switch {
		case errors.Is(err, fs.ErrNotExist) || replacedSymlink:
			if !opts.DryRun {
				if mkdirErr := os.MkdirAll(filepath.Dir(dst), 0o755); mkdirErr != nil {
					return s, fmt.Errorf("mkdir %s: %w", filepath.Dir(dst), mkdirErr)
				}
				if writeErr := os.WriteFile(dst, src, 0o644); writeErr != nil {
					return s, fmt.Errorf("write %s: %w", dst, writeErr)
				}
			}
			status := StatusCreated
			if replacedSymlink {
				status = StatusReplacedSymlink
			}
			s.Files = append(s.Files, FileChange{Path: dst, Status: status})
			s.Created++

		case err != nil:
			return s, fmt.Errorf("read existing %s: %w", dst, err)

		case bytes.Equal(existing, src):
			s.Files = append(s.Files, FileChange{Path: dst, Status: StatusUnchanged})
			s.Unchanged++

		case opts.Force:
			if !opts.DryRun {
				if writeErr := os.WriteFile(dst, src, 0o644); writeErr != nil {
					return s, fmt.Errorf("write %s: %w", dst, writeErr)
				}
			}
			s.Files = append(s.Files, FileChange{Path: dst, Status: StatusUpdated})
			s.Updated++

		default:
			s.Files = append(s.Files, FileChange{Path: dst, Status: StatusWouldOverwrite})
			s.Conflicts = append(s.Conflicts, dst)
		}
	}

	return s, nil
}

// IsInstalled reports whether the skill appears installed for an
// agent under a given scope. We check for the SKILL.md anchor file;
// individual missing resources don't count.
func IsInstalled(scopeRoot string, agent Agent, skill string) bool {
	anchor := filepath.Join(scopeRoot, agent.InstallSubdir, skill, "SKILL.md")
	_, err := os.Stat(anchor)
	return err == nil
}

// ScopeRoot resolves a scope name ("global" or "project") to the
// concrete directory it maps to.
func ScopeRoot(scope string) (string, error) {
	switch scope {
	case "global":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("resolve home dir: %w", err)
		}
		return home, nil
	case "project":
		cwd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("resolve current directory: %w", err)
		}
		return cwd, nil
	default:
		return "", fmt.Errorf("scope must be 'global' or 'project', got %q", scope)
	}
}
