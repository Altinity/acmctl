package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/term"
)

// printJSON writes value as indented JSON to stdout. value can be json.RawMessage,
// any encodable Go type, or nil (no-op).
func printJSON(v interface{}) error {
	if v == nil {
		return nil
	}
	if raw, ok := v.(json.RawMessage); ok && len(raw) == 0 {
		return nil
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

// printFiltered narrows a JSON array of cluster-like objects to ones whose
// id_environment / envId / environmentId / env_id matches the requested env.
// Used by `cluster list --env <id>`.
func printFiltered(raw json.RawMessage, envID string) error {
	var clusters []map[string]interface{}
	if err := json.Unmarshal(raw, &clusters); err != nil {
		return fmt.Errorf("expected array of clusters: %w", err)
	}
	out := []map[string]interface{}{}
	for _, c := range clusters {
		if matchesEnv(c, envID) {
			out = append(out, c)
		}
	}
	return printJSON(out)
}

func matchesEnv(c map[string]interface{}, envID string) bool {
	for _, key := range []string{"id_environment", "envId", "environmentId", "env_id"} {
		if v, ok := c[key]; ok {
			if fmt.Sprintf("%v", v) == envID {
				return true
			}
		}
	}
	return false
}

// isatty reports whether the file descriptor refers to a terminal.
func isatty(f *os.File) bool {
	return term.IsTerminal(int(f.Fd()))
}
