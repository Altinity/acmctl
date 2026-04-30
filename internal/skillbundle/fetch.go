package skillbundle

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

// Defaults for locating the published skill zip. Override
// ACMCTL_SKILLS_RELEASE_BASE for forks or local testing.
const (
	defaultRepoOwner = "Altinity"
	defaultRepoName  = "acmctl"

	// DefaultRef is the rolling pre-release tag the GitHub Action
	// publishes to on every push to main that touches skills/**.
	DefaultRef = "skill-bundle"
)

// Bundle is the in-memory representation of a downloaded skill zip.
// All entries are loaded into a map keyed by the path inside the
// zip ("altinity-cloud/SKILL.md" etc.), trimmed of any leading
// directory components.
type Bundle struct {
	// Skill is the name of the skill (matches the zip's top-level
	// directory and the filename minus .zip).
	Skill string
	// Files maps relative path within the skill (e.g., "SKILL.md",
	// "endpoints/clusters.md") to its byte contents.
	Files map[string][]byte
}

// Fetch downloads <release>/<skill>.zip from the configured GitHub
// release and returns a parsed Bundle. ref defaults to DefaultRef
// when empty.
func Fetch(ctx context.Context, skill, ref string) (*Bundle, error) {
	if ref == "" {
		ref = DefaultRef
	}
	url := archiveURL(ref, skill)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "acmctl-skills-fetch")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("download %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download %s: HTTP %d", url, resp.StatusCode)
	}

	// Read into memory — skill zips are tiny (<<1 MB).
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response from %s: %w", url, err)
	}

	zr, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return nil, fmt.Errorf("open zip from %s: %w", url, err)
	}

	bundle := &Bundle{Skill: skill, Files: map[string][]byte{}}
	prefix := skill + "/"
	for _, f := range zr.File {
		if f.FileInfo().IsDir() {
			continue
		}
		// Skip __MACOSX/ noise zip on macOS hosts may include.
		if strings.HasPrefix(f.Name, "__MACOSX/") {
			continue
		}
		// Defensive: refuse path traversal.
		clean := path.Clean(f.Name)
		if strings.HasPrefix(clean, "..") || strings.HasPrefix(clean, "/") {
			return nil, fmt.Errorf("zip contains unsafe path %q", f.Name)
		}
		// Strip the top-level "<skill>/" prefix; entries outside it
		// (rare, but possible if the zip is malformed) are dropped
		// with a clear error.
		if !strings.HasPrefix(clean, prefix) {
			return nil, fmt.Errorf("zip entry %q is not under expected prefix %q", clean, prefix)
		}
		rel := strings.TrimPrefix(clean, prefix)
		if rel == "" {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("open %s in zip: %w", f.Name, err)
		}
		data, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return nil, fmt.Errorf("read %s from zip: %w", f.Name, err)
		}
		bundle.Files[rel] = data
	}

	if len(bundle.Files) == 0 {
		return nil, fmt.Errorf("zip from %s is empty (no files under %q)", url, prefix)
	}

	return bundle, nil
}

// archiveURL returns the URL for a given ref + skill. Format:
//
//	{base}/{ref}/{skill}.zip
//
// Where {base} defaults to https://github.com/Altinity/acmctl/releases/download
// and is overridable via ACMCTL_SKILLS_RELEASE_BASE for testing.
func archiveURL(ref, skill string) string {
	base := os.Getenv("ACMCTL_SKILLS_RELEASE_BASE")
	if base == "" {
		base = fmt.Sprintf("https://github.com/%s/%s/releases/download", defaultRepoOwner, defaultRepoName)
	}
	return fmt.Sprintf("%s/%s/%s.zip", strings.TrimRight(base, "/"), ref, skill)
}
