package cmd

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/output"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage accounts",
}

var accountListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all accounts",
	RunE: func(cmd *cobra.Command, args []string) error {
		users, err := apiClient.ListAccounts()
		if err != nil {
			return err
		}
		return output.PrintTabulableList(cfg.Output, users)
	},
}

var accountGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get current account details",
	RunE: func(cmd *cobra.Command, args []string) error {
		user, err := apiClient.GetAccount()
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, user)
	},
}

var accountCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new account",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]string{}
		flagMap := map[string]string{
			"email": "email", "name": "name", "password": "password",
			"role": "id_role", "organization": "id_organization",
		}
		for flag, key := range flagMap {
			if v, _ := cmd.Flags().GetString(flag); v != "" {
				params[key] = v
			}
		}
		user, err := apiClient.CreateAccount(params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, user)
	},
}

var accountDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete an account",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteAccount(args[0]); err != nil {
			return err
		}
		fmt.Printf("Account %s deleted.\n", args[0])
		return nil
	},
}

var accountUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Modify an account",
	Long:  "Modify an account with arbitrary fields. Use --field key=val (repeatable).",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectFieldFlags(cmd)
		if err != nil {
			return err
		}
		user, err := apiClient.UpdateAccount(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, user)
	},
}

var accountUpdateSelfCmd = &cobra.Command{
	Use:   "update-self",
	Short: "Modify the current user's account",
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectFieldFlags(cmd)
		if err != nil {
			return err
		}
		result, err := apiClient.UpdateOwnAccount(params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, result)
	},
}

var accountAccessRightsCmd = &cobra.Command{
	Use:   "access-rights",
	Short: "List all available access rights settings",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetAccessRights()
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

var accountRolesListCmd = &cobra.Command{
	Use:   "roles",
	Short: "List all ACM account roles",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ListAccountRoles()
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

var accountRoleCreateCmd = &cobra.Command{
	Use:   "role-create",
	Short: "Create an ACM account role",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{"name": "name", "rights": "rights"})
		result, err := apiClient.CreateAccountRole(params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, result)
	},
}

var accountRoleUpdateCmd = &cobra.Command{
	Use:   "role-update <role-id>",
	Short: "Modify an ACM account role",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{"name": "name", "rights": "rights"})
		result, err := apiClient.UpdateAccountRole(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, result)
	},
}

var accountRoleDeleteCmd = &cobra.Command{
	Use:   "role-delete <role-id>",
	Short: "Remove an ACM account role",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteAccountRole(args[0]); err != nil {
			return err
		}
		fmt.Printf("Account role %s deleted.\n", args[0])
		return nil
	},
}

var accountTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Generate a random API key",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GenerateAccountToken()
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

var accountAnywhereTokenCmd = &cobra.Command{
	Use:   "anywhere-token",
	Short: "Generate a cloud.anywhere API key",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GenerateAnywhereToken()
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

var accountLogCmd = &cobra.Command{
	Use:   "log <user-id>",
	Short: "Show user account action log",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"page": "page", "limit": "limit", "filter": "filter", "order": "order",
		})
		result, err := apiClient.GetAccountLog(args[0], params)
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

var accountAccessCmd = &cobra.Command{
	Use:   "access <id>",
	Short: "Change environment access for an account",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		environments, _ := cmd.Flags().GetString("environments")
		params := map[string]string{}
		if environments != "" {
			params["environments"] = environments
		}
		if err := apiClient.SetAccountAccess(args[0], params); err != nil {
			return err
		}
		fmt.Printf("Access updated for account %s.\n", args[0])
		return nil
	},
}

func init() {
	accountCreateCmd.Flags().String("email", "", "email address")
	accountCreateCmd.Flags().String("name", "", "display name")
	accountCreateCmd.Flags().String("password", "", "password")
	accountCreateCmd.Flags().String("role", "", "role ID")
	accountCreateCmd.Flags().String("organization", "", "organization ID")

	accountUpdateCmd.Flags().StringSliceP("field", "F", nil, "key=value (repeatable)")
	accountUpdateSelfCmd.Flags().StringSliceP("field", "F", nil, "key=value (repeatable)")

	for _, c := range []*cobra.Command{accountRoleCreateCmd, accountRoleUpdateCmd} {
		c.Flags().String("name", "", "role name")
		c.Flags().String("rights", "", "rights expression")
	}

	accountLogCmd.Flags().String("page", "", "page number")
	accountLogCmd.Flags().String("limit", "", "page size")
	accountLogCmd.Flags().String("filter", "", "filter (JSON)")
	accountLogCmd.Flags().String("order", "", "order (JSON: by, direction)")

	accountAccessCmd.Flags().String("environments", "", "comma-separated environment IDs")

	accountCmd.AddCommand(accountListCmd)
	accountCmd.AddCommand(accountGetCmd)
	accountCmd.AddCommand(accountCreateCmd)
	accountCmd.AddCommand(accountUpdateCmd)
	accountCmd.AddCommand(accountUpdateSelfCmd)
	accountCmd.AddCommand(accountDeleteCmd)
	accountCmd.AddCommand(accountAccessRightsCmd)
	accountCmd.AddCommand(accountRolesListCmd)
	accountCmd.AddCommand(accountRoleCreateCmd)
	accountCmd.AddCommand(accountRoleUpdateCmd)
	accountCmd.AddCommand(accountRoleDeleteCmd)
	accountCmd.AddCommand(accountTokenCmd)
	accountCmd.AddCommand(accountAnywhereTokenCmd)
	accountCmd.AddCommand(accountLogCmd)
	accountCmd.AddCommand(accountAccessCmd)
	rootCmd.AddCommand(accountCmd)
}
