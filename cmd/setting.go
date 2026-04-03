package cmd

import (
	"fmt"

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

var settingCreateCmd = &cobra.Command{
	Use:   "create <cluster-id>",
	Short: "Create a cluster setting",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]string{}
		flagMap := map[string]string{
			"name": "name", "value": "value", "description": "description",
		}
		for flag, key := range flagMap {
			if v, _ := cmd.Flags().GetString(flag); v != "" {
				params[key] = v
			}
		}
		setting, err := apiClient.CreateClusterSetting(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, setting)
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
	settingCreateCmd.Flags().String("value", "", "setting value")
	settingCreateCmd.Flags().String("description", "", "setting description")

	settingCmd.AddCommand(settingListCmd)
	settingCmd.AddCommand(settingCreateCmd)
	settingCmd.AddCommand(settingDeleteCmd)
	rootCmd.AddCommand(settingCmd)
}
