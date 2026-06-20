package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// clusterSetting is one ACM cluster config setting — either a config.d file
// (isFile=true) or a scalar server setting. Note: the API returns both id and
// id_cluster as strings.
type clusterSetting struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	IsFile    bool   `json:"isFile"`
	IDCluster string `json:"id_cluster"`
	System    bool   `json:"system"`
}

// fetchSettings returns all settings for a cluster (GET /cluster/<id>/settings).
func fetchSettings(cid string) ([]clusterSetting, error) {
	var out []clusterSetting
	if err := apiClient.Do("GET", fmt.Sprintf("/cluster/%s/settings", cid), nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// findSetting locates a setting by exact name or by id. Returns nil (no error)
// if absent. There is no GET-by-id endpoint, so this lists and filters.
func findSetting(cid, key string) (*clusterSetting, error) {
	all, err := fetchSettings(cid)
	if err != nil {
		return nil, err
	}
	for i := range all {
		if all[i].Name == key || all[i].ID == key {
			return &all[i], nil
		}
	}
	return nil, nil
}

// looksLikeFile guesses whether a setting name is a config.d file (vs a scalar
// server setting), used as the default for `isFile` when creating one.
func looksLikeFile(name string) bool {
	if strings.Contains(name, "/") {
		return true
	}
	for _, ext := range []string{".xml", ".sql", ".json", ".yaml", ".yml"} {
		if strings.HasSuffix(name, ext) {
			return true
		}
	}
	return false
}

func requireIntID(label, v string) error {
	if !integerIDRegexp.MatchString(v) {
		return fmt.Errorf("%s %q: expected integer ID", label, v)
	}
	return nil
}

var clusterSettingsCmd = &cobra.Command{
	Use:     "settings",
	Aliases: []string{"setting"},
	Short:   "Manage a cluster's config settings (config.d files + server settings)",
}

var clusterSettingsListCmd = &cobra.Command{
	Use:   "list <cluster-id>",
	Short: "List a cluster's settings (id, kind, name). --json for full objects.",
	Args:  integerIDArg(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		all, err := fetchSettings(args[0])
		if err != nil {
			return err
		}
		if v, _ := cmd.Flags().GetBool("json"); v {
			return printJSON(all) // includes values (may contain secrets)
		}
		for _, s := range all {
			kind := "scalar"
			if s.IsFile {
				kind = "file"
			}
			fmt.Printf("%-8s %-6s %s\n", s.ID, kind, s.Name)
		}
		return nil
	},
}

var clusterSettingsGetCmd = &cobra.Command{
	Use:   "get <cluster-id> <name|id>",
	Short: "Print one setting's value (resolved by name or id). --json for the full object.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireIntID("cluster-id", args[0]); err != nil {
			return err
		}
		s, err := findSetting(args[0], args[1])
		if err != nil {
			return err
		}
		if s == nil {
			return fmt.Errorf("no setting with name or id %q on cluster %s", args[1], args[0])
		}
		if v, _ := cmd.Flags().GetBool("json"); v {
			return printJSON(s)
		}
		fmt.Print(s.Value)
		if !strings.HasSuffix(s.Value, "\n") {
			fmt.Println()
		}
		return nil
	},
}

var clusterSettingsSetCmd = &cobra.Command{
	Use:   "set <cluster-id> <name>",
	Short: "Create or update a setting by name. Updates in place if it exists.",
	Long: `Create or update a cluster setting, idempotent by name.

Value comes from exactly one of --file (read from disk) or --value (inline).
If a setting with that name already exists it is updated in place by id
(POST /cluster-setting/<id>); otherwise it is created
(POST /cluster/<cid>/settings) with isFile inferred from the name
(config.d/*, *.xml, *.sql, …) unless --is-file is given.

Changes are staged; run 'acmctl cluster push <cluster-id>' to apply them
(which may restart ClickHouse).`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		cid, name := args[0], args[1]
		if err := requireIntID("cluster-id", cid); err != nil {
			return err
		}

		file, _ := cmd.Flags().GetString("file")
		valueSet := cmd.Flags().Changed("value")
		if (file != "") == valueSet { // both or neither
			return fmt.Errorf("provide exactly one of --file or --value")
		}
		value, _ := cmd.Flags().GetString("value")
		if file != "" {
			b, err := os.ReadFile(file)
			if err != nil {
				return fmt.Errorf("read %s: %w", file, err)
			}
			value = string(b)
		}

		existing, err := findSetting(cid, name)
		if err != nil {
			return err
		}

		var raw json.RawMessage
		if existing != nil {
			body, _ := json.Marshal(map[string]string{"name": existing.Name, "value": value})
			if err := apiClient.DoJSON("POST", fmt.Sprintf("/cluster-setting/%s", existing.ID), body, &raw); err != nil {
				return err
			}
			fmt.Fprintf(os.Stderr, "updated %s (id %s)\n", existing.Name, existing.ID)
		} else {
			isFile := looksLikeFile(name)
			if cmd.Flags().Changed("is-file") {
				isFile, _ = cmd.Flags().GetBool("is-file")
			}
			cidNum, _ := strconv.Atoi(cid)
			body, _ := json.Marshal(map[string]interface{}{
				"name": name, "value": value, "id_cluster": cidNum, "isFile": isFile,
			})
			if err := apiClient.DoJSON("POST", fmt.Sprintf("/cluster/%s/settings", cid), body, &raw); err != nil {
				return err
			}
			fmt.Fprintf(os.Stderr, "created %s (isFile=%t)\n", name, isFile)
		}
		return printJSON(raw)
	},
}

var clusterSettingsRmCmd = &cobra.Command{
	Use:     "rm <cluster-id> <name|id>",
	Aliases: []string{"delete"},
	Short:   "Delete a setting (resolved by name or id)",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireIntID("cluster-id", args[0]); err != nil {
			return err
		}
		id := args[1]
		if !integerIDRegexp.MatchString(id) {
			s, err := findSetting(args[0], args[1])
			if err != nil {
				return err
			}
			if s == nil {
				return fmt.Errorf("no setting with name %q on cluster %s", args[1], args[0])
			}
			id = s.ID
		}
		return apiClient.Do("DELETE", fmt.Sprintf("/cluster-setting/%s", id), nil, nil)
	},
}

var clusterPushCmd = &cobra.Command{
	Use:   "push <cluster-id>",
	Short: "Apply staged settings to the cluster (may restart ClickHouse)",
	Args:  integerIDArg(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var raw json.RawMessage
		if err := apiClient.DoJSON("POST", fmt.Sprintf("/cluster/%s/push", args[0]), nil, &raw); err != nil {
			return err
		}
		return printJSON(raw)
	},
}

func init() {
	clusterSettingsListCmd.Flags().Bool("json", false, "output full setting objects as JSON (includes values)")
	clusterSettingsGetCmd.Flags().Bool("json", false, "output the full setting object as JSON")
	clusterSettingsSetCmd.Flags().String("file", "", "read the value from a file")
	clusterSettingsSetCmd.Flags().String("value", "", "inline value")
	clusterSettingsSetCmd.Flags().Bool("is-file", false, "store as a config.d file (default: inferred from the name)")

	clusterSettingsCmd.AddCommand(clusterSettingsListCmd, clusterSettingsGetCmd, clusterSettingsSetCmd, clusterSettingsRmCmd)
	clusterCmd.AddCommand(clusterSettingsCmd)
	clusterCmd.AddCommand(clusterPushCmd)
}
