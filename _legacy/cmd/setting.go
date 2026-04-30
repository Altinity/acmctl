package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/altinity/acmctl/pkg/output"
	"github.com/spf13/cobra"
)

var settingCmd = &cobra.Command{
	Use:   "setting",
	Short: "Manage cluster settings",
}

var settingListCmd = &cobra.Command{
	Use:   "list <cluster-id>",
	Short: "List cluster settings",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		settings, err := apiClient.ListClusterSettings(args[0])
		if err != nil {
			return err
		}
		return output.PrintTabulableList(cfg.Output, settings)
	},
}

// resolveValue expands @file / @- syntax in a flag value.
//
//	"@/path/to/file" → contents of file
//	"@-"             → contents of stdin
//	"@@..."          → literal "@..." (escape hatch for values that should start with '@')
//	anything else    → returned unchanged
func resolveValue(v string) (string, error) {
	if !strings.HasPrefix(v, "@") {
		return v, nil
	}
	if strings.HasPrefix(v, "@@") {
		return v[1:], nil
	}
	if v == "@-" {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("read stdin: %w", err)
		}
		return string(b), nil
	}
	path := v[1:]
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", path, err)
	}
	return string(b), nil
}

// collectSettingFlags reads --name/--value/--description, expanding @file/@- where applicable.
func collectSettingFlags(cmd *cobra.Command) (map[string]string, error) {
	params := map[string]string{}
	for _, key := range []string{"name", "value", "description"} {
		v, _ := cmd.Flags().GetString(key)
		if v == "" {
			continue
		}
		// `name` is intentionally not @-expanded — it's always a short identifier.
		if key != "name" {
			resolved, err := resolveValue(v)
			if err != nil {
				return nil, err
			}
			v = resolved
		}
		params[key] = v
	}
	return params, nil
}

var settingCreateCmd = &cobra.Command{
	Use:   "create <cluster-id>",
	Short: "Create a cluster setting",
	Long: `Create a cluster setting.

Use @/path/to/file to load --value or --description from a file, or @- to read
from stdin. Prefix a literal value with @@ to escape (so "@@x" is sent as "@x").

Large values are sent in the request body, so XML/JSON content above Apache's
~8 KB URI limit works correctly.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectSettingFlags(cmd)
		if err != nil {
			return err
		}
		setting, err := apiClient.CreateClusterSetting(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, setting)
	},
}

var settingUpdateCmd = &cobra.Command{
	Use:   "update <setting-id>",
	Short: "Update a cluster setting",
	Long: `Update an existing cluster setting in place.

Use @/path/to/file or @- (stdin) for --value or --description, same as
'setting create'. Prefer this over delete+create — delete-then-failed-create
can leave the cluster in a broken state.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectSettingFlags(cmd)
		if err != nil {
			return err
		}
		setting, err := apiClient.UpdateClusterSetting(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, setting)
	},
}

var settingEnvListCmd = &cobra.Command{
	Use:   "env-list <cluster-id>",
	Short: "List cluster environment settings",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		settings, err := apiClient.ListClusterEnvSettings(args[0])
		if err != nil {
			return err
		}
		return output.PrintTabulableList(cfg.Output, settings)
	},
}

var settingEnvCreateCmd = &cobra.Command{
	Use:   "env-create <cluster-id>",
	Short: "Add an environment setting to a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectSettingFlags(cmd)
		if err != nil {
			return err
		}
		setting, err := apiClient.CreateClusterEnvSetting(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, setting)
	},
}

var settingEnvUpdateCmd = &cobra.Command{
	Use:   "env-update <setting-id>",
	Short: "Modify an environment setting for a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectSettingFlags(cmd)
		if err != nil {
			return err
		}
		setting, err := apiClient.UpdateClusterEnvSetting(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, setting)
	},
}

var settingEnvDeleteCmd = &cobra.Command{
	Use:   "env-delete <setting-id>",
	Short: "Remove an environment setting from a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteClusterEnvSetting(args[0]); err != nil {
			return err
		}
		fmt.Printf("Env setting %s deleted.\n", args[0])
		return nil
	},
}

var settingSystemListCmd = &cobra.Command{
	Use:   "system-list <cluster-id>",
	Short: "List cluster system settings",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ListClusterSystemSettings(args[0])
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

var settingDeleteCmd = &cobra.Command{
	Use:   "delete <setting-id>",
	Short: "Delete a cluster setting",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteClusterSetting(args[0]); err != nil {
			return err
		}
		fmt.Printf("Setting %s deleted.\n", args[0])
		return nil
	},
}

func init() {
	settingCreateCmd.Flags().String("name", "", "setting name")
	settingCreateCmd.Flags().String("value", "", "setting value (use @file or @- for stdin)")
	settingCreateCmd.Flags().String("description", "", "setting description (use @file or @- for stdin)")

	settingUpdateCmd.Flags().String("name", "", "setting name")
	settingUpdateCmd.Flags().String("value", "", "setting value (use @file or @- for stdin)")
	settingUpdateCmd.Flags().String("description", "", "setting description (use @file or @- for stdin)")

	for _, c := range []*cobra.Command{settingEnvCreateCmd, settingEnvUpdateCmd} {
		c.Flags().String("name", "", "setting name")
		c.Flags().String("value", "", "setting value (use @file or @- for stdin)")
		c.Flags().String("description", "", "setting description (use @file or @- for stdin)")
	}

	settingCmd.AddCommand(settingListCmd)
	settingCmd.AddCommand(settingCreateCmd)
	settingCmd.AddCommand(settingUpdateCmd)
	settingCmd.AddCommand(settingDeleteCmd)
	settingCmd.AddCommand(settingEnvListCmd)
	settingCmd.AddCommand(settingEnvCreateCmd)
	settingCmd.AddCommand(settingEnvUpdateCmd)
	settingCmd.AddCommand(settingEnvDeleteCmd)
	settingCmd.AddCommand(settingSystemListCmd)
	rootCmd.AddCommand(settingCmd)
}
