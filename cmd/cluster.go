package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

// integerIDRegexp matches one or more digits with no other characters.
// Used to reject path-injection-shaped args like "337/foo" or "337?x=1"
// before they're spliced into URL paths via fmt.Sprintf.
var integerIDRegexp = regexp.MustCompile(`^[0-9]+$`)

// integerIDArg returns a cobra.PositionalArgs validator that requires
// exactly n integer-shaped arguments. Optional args (e.g., env-id for
// launch) should be validated by the command itself, not here.
func integerIDArg(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != n {
			return fmt.Errorf("accepts %d arg(s), received %d", n, len(args))
		}
		for i, a := range args {
			if !integerIDRegexp.MatchString(a) {
				return fmt.Errorf("arg %d (%q): expected integer ID", i, a)
			}
		}
		return nil
	}
}

var clusterCmd = &cobra.Command{
	Use:     "cluster",
	Aliases: []string{"cl"},
	Short:   "Manage clusters",
}

var clusterListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all clusters",
	RunE: func(cmd *cobra.Command, args []string) error {
		envFilter, _ := cmd.Flags().GetString("env")
		var raw json.RawMessage
		if err := apiClient.Do("GET", "/clusters", nil, &raw); err != nil {
			return err
		}
		if envFilter != "" {
			return printFiltered(raw, envFilter)
		}
		return printJSON(raw)
	},
}

var clusterGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get cluster details",
	Args:  integerIDArg(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var raw json.RawMessage
		if err := apiClient.Do("GET", fmt.Sprintf("/cluster/%s", args[0]), nil, &raw); err != nil {
			return err
		}
		return printJSON(raw)
	},
}

var clusterLaunchCmd = &cobra.Command{
	Use:   "launch [env-id]",
	Short: "Launch a cluster (JSON body on stdin). env-id falls back to ACM_API_ENV_ID.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("accepts at most 1 arg, received %d", len(args))
		}
		if len(args) == 1 && !integerIDRegexp.MatchString(args[0]) {
			return fmt.Errorf("env-id %q: expected integer", args[0])
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		envID := envIDFromArgsOrEnv(args)
		if envID == "" {
			return fmt.Errorf("env-id required (positional arg or ACM_API_ENV_ID)")
		}
		if !integerIDRegexp.MatchString(envID) {
			return fmt.Errorf("env-id %q: expected integer (check ACM_API_ENV_ID)", envID)
		}
		body, err := readStdinJSON()
		if err != nil {
			return err
		}
		var raw json.RawMessage
		if err := apiClient.DoJSON("POST", fmt.Sprintf("/environment/%s/clusters/launch", envID), body, &raw); err != nil {
			return err
		}
		return printJSON(raw)
	},
}

var clusterUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Modify a cluster (JSON body on stdin)",
	Args:  integerIDArg(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		body, err := readStdinJSON()
		if err != nil {
			return err
		}
		var raw json.RawMessage
		if err := apiClient.DoJSON("POST", fmt.Sprintf("/cluster/%s", args[0]), body, &raw); err != nil {
			return err
		}
		return printJSON(raw)
	},
}

var clusterDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a cluster (--terminate to also tear down resources)",
	Args:  integerIDArg(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		t := "0"
		if v, _ := cmd.Flags().GetBool("terminate"); v {
			t = "1"
		}
		return apiClient.Do("DELETE", fmt.Sprintf("/cluster/%s/%s", args[0], t), nil, nil)
	},
}

var clusterTempCredsCmd = &cobra.Command{
	Use:   "temp-creds <id>",
	Short: "Mint temporary Altinity-support credentials",
	Long: `GET /cluster/{id}/support/credentials.

Pass-through: outputs whatever the API returns inside the .data envelope,
without reshaping. Historically the response format has varied — sometimes a
plain password string, sometimes {login, password}. Caller should handle both
shapes (and falls back to its own session username if only a password is
returned).`,
	Args: integerIDArg(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var raw json.RawMessage
		if err := apiClient.Do("GET", fmt.Sprintf("/cluster/%s/support/credentials", args[0]), nil, &raw); err != nil {
			return err
		}
		return printJSON(raw)
	},
}

func envIDFromArgsOrEnv(args []string) string {
	if len(args) > 0 && args[0] != "" {
		return args[0]
	}
	return os.Getenv("ACM_API_ENV_ID")
}

func readStdinJSON() ([]byte, error) {
	if isatty(os.Stdin) {
		return nil, fmt.Errorf("expected JSON body on stdin (e.g. cat cluster.json | acmctl cluster launch)")
	}
	body, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("read stdin: %w", err)
	}
	if len(body) == 0 {
		return nil, fmt.Errorf("empty stdin")
	}
	if !json.Valid(body) {
		return nil, fmt.Errorf("stdin is not valid JSON")
	}
	return body, nil
}

func init() {
	clusterListCmd.Flags().String("env", "", "filter by environment ID (client-side)")
	clusterDeleteCmd.Flags().Bool("terminate", false, "terminate cluster resources")

	clusterCmd.AddCommand(clusterListCmd)
	clusterCmd.AddCommand(clusterGetCmd)
	clusterCmd.AddCommand(clusterLaunchCmd)
	clusterCmd.AddCommand(clusterUpdateCmd)
	clusterCmd.AddCommand(clusterDeleteCmd)
	clusterCmd.AddCommand(clusterTempCredsCmd)
	rootCmd.AddCommand(clusterCmd)
}
