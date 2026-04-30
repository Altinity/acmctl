package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rawCmd = &cobra.Command{
	Use:   "raw <METHOD> <path>",
	Short: "Call any ACM endpoint directly. Body auto-detected from stdin (JSON) or -F flags (form-urlencoded).",
	Long: `Generic passthrough to the ACM API.

Body shape is auto-detected:
  - JSON body:    pipe a JSON object/array on stdin
                  e.g.  echo '{"name":"foo"}' | acmctl raw POST /cluster/337
  - Form fields:  use -F key=value (repeatable)
                  e.g.  acmctl raw POST /cluster/337/query -F query='SELECT 1'
  - No body:      omit both (most GET/DELETE endpoints)

The same auth header and base URL as other commands. Output is JSON.

Path is everything after the API root (e.g. /cluster/337/status, /clusters).
Don't include the host or /api/ prefix.`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		method := strings.ToUpper(args[0])
		path := args[1]

		fields, _ := cmd.Flags().GetStringSlice("field")

		stdinPiped := !isatty(os.Stdin)
		hasFields := len(fields) > 0

		if stdinPiped && hasFields {
			return fmt.Errorf("cannot combine stdin JSON body with -F form fields")
		}

		var raw json.RawMessage
		switch {
		case stdinPiped:
			body, err := io.ReadAll(os.Stdin)
			if err != nil {
				return fmt.Errorf("read stdin: %w", err)
			}
			if len(body) > 0 && !json.Valid(body) {
				return fmt.Errorf("stdin is not valid JSON")
			}
			err = apiClient.DoJSON(method, path, body, &raw)
			if err != nil {
				return err
			}
		case hasFields:
			form, err := parseFormFields(fields)
			if err != nil {
				return err
			}
			err = apiClient.DoForm(method, path, form, &raw)
			if err != nil {
				return err
			}
		default:
			err := apiClient.Do(method, path, nil, &raw)
			if err != nil {
				return err
			}
		}
		return printJSON(raw)
	},
}

// parseFormFields turns ["k=v", "k2=v2", "k3=@file"] into a map. The "@file"
// suffix loads value from disk; "@-" reads from stdin (only one such allowed).
func parseFormFields(fields []string) (map[string]string, error) {
	out := map[string]string{}
	for _, kv := range fields {
		eq := strings.Index(kv, "=")
		if eq <= 0 {
			return nil, fmt.Errorf("invalid -F %q (expected key=value)", kv)
		}
		k, v := kv[:eq], kv[eq+1:]
		if strings.HasPrefix(v, "@") && v != "@" {
			if v == "@-" {
				b, err := io.ReadAll(os.Stdin)
				if err != nil {
					return nil, fmt.Errorf("read stdin for %s: %w", k, err)
				}
				v = string(b)
			} else {
				b, err := os.ReadFile(v[1:])
				if err != nil {
					return nil, fmt.Errorf("read %s: %w", v[1:], err)
				}
				v = string(b)
			}
		}
		out[k] = v
	}
	return out, nil
}

func init() {
	rawCmd.Flags().StringSliceP("field", "F", nil, "form field key=value (repeatable; @file loads from disk)")
	rootCmd.AddCommand(rawCmd)
}
