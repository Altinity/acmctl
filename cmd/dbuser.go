package cmd

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/output"
	"github.com/spf13/cobra"
)

var dbuserCmd = &cobra.Command{
	Use:   "dbuser",
	Short: "Manage database users",
}

var dbuserListCmd = &cobra.Command{
	Use:   "list <cluster-id>",
	Short: "List database users for a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		users, err := apiClient.ListDbUsers(args[0])
		if err != nil {
			return err
		}
		return output.PrintTabulableList(cfg.Output, users)
	},
}

var dbuserCreateCmd = &cobra.Command{
	Use:   "create <cluster-id>",
	Short: "Create a database user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]string{}
		flagMap := map[string]string{
			"login": "login", "password": "password",
			"networks": "networks", "databases": "databases",
		}
		for flag, key := range flagMap {
			if v, _ := cmd.Flags().GetString(flag); v != "" {
				params[key] = v
			}
		}
		if v, _ := cmd.Flags().GetBool("access-management"); v {
			params["accessManagement"] = "true"
		}
		user, err := apiClient.CreateDbUser(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, user)
	},
}

var dbuserDeleteCmd = &cobra.Command{
	Use:   "delete <cluster-id> <user-id>",
	Short: "Delete a database user",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteDbUser(args[0], args[1]); err != nil {
			return err
		}
		fmt.Printf("Database user %s deleted.\n", args[1])
		return nil
	},
}

func init() {
	dbuserCreateCmd.Flags().String("login", "", "username")
	dbuserCreateCmd.Flags().String("password", "", "password")
	dbuserCreateCmd.Flags().String("networks", "", "allowed networks")
	dbuserCreateCmd.Flags().String("databases", "", "allowed databases")
	dbuserCreateCmd.Flags().Bool("access-management", false, "grant access management")

	dbuserCmd.AddCommand(dbuserListCmd)
	dbuserCmd.AddCommand(dbuserCreateCmd)
	dbuserCmd.AddCommand(dbuserDeleteCmd)
	rootCmd.AddCommand(dbuserCmd)
}
