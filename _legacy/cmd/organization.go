package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var organizationCmd = &cobra.Command{
	Use:     "organization",
	Aliases: []string{"org"},
	Short:   "Manage organizations",
}

var organizationListCmd = &cobra.Command{
	Use:   "list",
	Short: "List organizations",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"blocked": "blocked", "billing": "billing",
		})
		result, err := apiClient.ListOrganizations(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var organizationGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get an organization",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetOrganization(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var organizationCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an organization",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := organizationFlags(cmd)
		result, err := apiClient.CreateOrganization(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var organizationUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Modify an organization",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := organizationFlags(cmd)
		result, err := apiClient.UpdateOrganization(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var organizationDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Remove an organization",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		terminate, _ := cmd.Flags().GetBool("terminate")
		if err := apiClient.DeleteOrganization(args[0], terminate); err != nil {
			return err
		}
		fmt.Printf("Organization %s deleted.\n", args[0])
		return nil
	},
}

var organizationLoginsCmd = &cobra.Command{
	Use:   "logins <id>",
	Short: "Configure organization login settings",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"name": "name", "blocked-password": "blockedPassword", "blocked-api": "blockedAPI",
			"allow-admin-password": "allowAdminPassword", "enable-2fa": "enable2FA",
			"opened": "opened", "default-user-role": "id_defaultUserRole",
			"user-sync-settings": "userSyncSettings",
		})
		result, err := apiClient.UpdateOrganizationLogins(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

func organizationFlags(cmd *cobra.Command) map[string]string {
	return flagsToParams(cmd, map[string]string{
		"name": "name", "email-domain": "emailDomain",
		"inherit-environment": "inheritEnvironment", "opened": "opened",
		"default-user-role": "id_defaultUserRole", "limited": "limited",
		"blocked": "blocked", "blocked-password": "blockedPassword",
		"blocked-api": "blockedAPI", "allow-admin-password": "allowAdminPassword",
		"enable-2fa": "enable2FA", "environments": "environments",
		"trial-expiry": "trialExpiry", "auto-charge": "autoCharge",
		"user-sync-settings": "userSyncSettings", "auto-charge-no-limit": "autoChargeNoLimit",
	})
}

func init() {
	organizationListCmd.Flags().String("blocked", "", "filter blocked")
	organizationListCmd.Flags().String("billing", "", "include billing")

	for _, c := range []*cobra.Command{organizationCreateCmd, organizationUpdateCmd} {
		for _, f := range []string{
			"name", "email-domain", "inherit-environment", "opened",
			"default-user-role", "limited", "blocked", "blocked-password",
			"blocked-api", "allow-admin-password", "enable-2fa", "environments",
			"trial-expiry", "auto-charge", "user-sync-settings", "auto-charge-no-limit",
		} {
			c.Flags().String(f, "", "")
		}
	}

	for _, f := range []string{
		"name", "blocked-password", "blocked-api", "allow-admin-password",
		"enable-2fa", "opened", "default-user-role", "user-sync-settings",
	} {
		organizationLoginsCmd.Flags().String(f, "", "")
	}

	organizationDeleteCmd.Flags().Bool("terminate", false, "terminate associated resources")

	for _, c := range []*cobra.Command{
		organizationListCmd, organizationGetCmd, organizationCreateCmd,
		organizationUpdateCmd, organizationDeleteCmd, organizationLoginsCmd,
	} {
		organizationCmd.AddCommand(c)
	}
	rootCmd.AddCommand(organizationCmd)
}
