package cmd

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/output"
	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage cluster settings profiles and their settings",
}

var profileChOptionsCmd = &cobra.Command{
	Use:   "ch-options <profile-id>",
	Short: "List all available ClickHouse settings for a profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		settings, err := apiClient.ListProfileChOptions(args[0])
		if err != nil {
			return err
		}
		return output.PrintTabulableList(cfg.Output, settings)
	},
}

var profileSettingsListCmd = &cobra.Command{
	Use:   "settings <profile-id>",
	Short: "List modified settings for a profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		settings, err := apiClient.ListProfileSettings(args[0])
		if err != nil {
			return err
		}
		return output.PrintTabulableList(cfg.Output, settings)
	},
}

var profileSettingCreateCmd = &cobra.Command{
	Use:   "setting-create <profile-id>",
	Short: "Add a setting modification to a profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectSettingFlags(cmd)
		if err != nil {
			return err
		}
		setting, err := apiClient.CreateProfileSetting(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, setting)
	},
}

var profileSettingUpdateCmd = &cobra.Command{
	Use:   "setting-update <setting-id>",
	Short: "Modify a profile setting",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectSettingFlags(cmd)
		if err != nil {
			return err
		}
		setting, err := apiClient.UpdateProfileSetting(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, setting)
	},
}

var profileSettingDeleteCmd = &cobra.Command{
	Use:   "setting-delete <setting-id>",
	Short: "Remove a profile setting",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteProfileSetting(args[0]); err != nil {
			return err
		}
		fmt.Printf("Profile setting %s deleted.\n", args[0])
		return nil
	},
}

var profileCreateCmd = &cobra.Command{
	Use:   "create <cluster-id>",
	Short: "Create a settings profile for a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"name": "name", "description": "description",
		})
		profile, err := apiClient.CreateClusterProfile(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, profile)
	},
}

var profileUpdateCmd = &cobra.Command{
	Use:   "update <profile-id>",
	Short: "Modify a settings profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"name": "name", "description": "description",
		})
		result, err := apiClient.UpdateProfile(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var profileDeleteCmd = &cobra.Command{
	Use:   "delete <profile-id>",
	Short: "Remove a settings profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteProfile(args[0]); err != nil {
			return err
		}
		fmt.Printf("Profile %s deleted.\n", args[0])
		return nil
	},
}

func init() {
	for _, c := range []*cobra.Command{profileSettingCreateCmd, profileSettingUpdateCmd} {
		c.Flags().String("name", "", "setting name")
		c.Flags().String("value", "", "setting value (use @file or @- for stdin)")
		c.Flags().String("description", "", "setting description (use @file or @- for stdin)")
	}

	for _, c := range []*cobra.Command{profileCreateCmd, profileUpdateCmd} {
		c.Flags().String("name", "", "profile name")
		c.Flags().String("description", "", "profile description")
	}

	for _, c := range []*cobra.Command{
		profileChOptionsCmd, profileSettingsListCmd, profileSettingCreateCmd, profileSettingUpdateCmd,
		profileSettingDeleteCmd, profileCreateCmd, profileUpdateCmd, profileDeleteCmd,
	} {
		profileCmd.AddCommand(c)
	}
	rootCmd.AddCommand(profileCmd)
}
