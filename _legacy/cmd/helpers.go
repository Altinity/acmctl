package cmd

import (
	"fmt"
	"strings"

	"github.com/altinity/acmctl/pkg/output"
	"github.com/spf13/cobra"
)

// printJSON prints any value as JSON. Used by commands whose responses don't
// have a typed schema in reference_auth.json, where rendering as table/yaml
// would be misleading.
func printJSON(v interface{}) error {
	return output.Print("json", v)
}

// collectFieldFlags reads repeatable --field key=value flags into a params map.
// Supports the @file / @- expansion via resolveValue() so large XML/JSON values can
// be loaded from a file or stdin.
func collectFieldFlags(cmd *cobra.Command) (map[string]string, error) {
	raw, _ := cmd.Flags().GetStringSlice("field")
	out := map[string]string{}
	for _, kv := range raw {
		idx := strings.Index(kv, "=")
		if idx <= 0 {
			return nil, fmt.Errorf("invalid --field %q (expected key=value)", kv)
		}
		k := kv[:idx]
		v := kv[idx+1:]
		resolved, err := resolveValue(v)
		if err != nil {
			return nil, err
		}
		out[k] = resolved
	}
	return out, nil
}

// flagsToParams reads a list of named string flags into a params map keyed by API
// field name. Empty values are omitted. Each entry is "flagName:apiKey".
func flagsToParams(cmd *cobra.Command, mapping map[string]string) map[string]string {
	out := map[string]string{}
	for flag, key := range mapping {
		if v, _ := cmd.Flags().GetString(flag); v != "" {
			out[key] = v
		}
	}
	return out
}
