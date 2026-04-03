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

func init() {
	accountCreateCmd.Flags().String("email", "", "email address")
	accountCreateCmd.Flags().String("name", "", "display name")
	accountCreateCmd.Flags().String("password", "", "password")
	accountCreateCmd.Flags().String("role", "", "role ID")
	accountCreateCmd.Flags().String("organization", "", "organization ID")

	accountCmd.AddCommand(accountListCmd)
	accountCmd.AddCommand(accountGetCmd)
	accountCmd.AddCommand(accountCreateCmd)
	accountCmd.AddCommand(accountDeleteCmd)
	rootCmd.AddCommand(accountCmd)
}
