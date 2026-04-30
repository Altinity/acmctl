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

var dbuserUpdateCmd = &cobra.Command{
	Use:   "update <cluster-id> <user-id>",
	Short: "Modify a database user",
	Long:  "Modify a database user with arbitrary fields. Use --field key=val (repeatable).",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectFieldFlags(cmd)
		if err != nil {
			return err
		}
		user, err := apiClient.UpdateDbUser(args[0], args[1], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, user)
	},
}

var dbuserModifyCmd = &cobra.Command{
	Use:   "modify <user-id>",
	Short: "Modify a database user via /user/{id} (no cluster ID)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"login": "login", "password": "password", "databases": "databases",
			"profile": "id_profile", "quota": "id_quota",
		})
		if v, _ := cmd.Flags().GetBool("access-management"); v {
			params["accessManagement"] = "true"
		}
		user, err := apiClient.UpdateDbUserByID(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, user)
	},
}

var dbuserRemoveCmd = &cobra.Command{
	Use:   "remove <user-id>",
	Short: "Remove a database user via /user/{id} (no cluster ID)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteDbUserByID(args[0]); err != nil {
			return err
		}
		fmt.Printf("User %s removed.\n", args[0])
		return nil
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

	dbuserUpdateCmd.Flags().StringSliceP("field", "F", nil, "key=value (repeatable)")

	dbuserModifyCmd.Flags().String("login", "", "username")
	dbuserModifyCmd.Flags().String("password", "", "password")
	dbuserModifyCmd.Flags().String("databases", "", "allowed databases")
	dbuserModifyCmd.Flags().String("profile", "", "settings profile ID")
	dbuserModifyCmd.Flags().String("quota", "", "quota ID")
	dbuserModifyCmd.Flags().Bool("access-management", false, "grant access management")

	dbuserCmd.AddCommand(dbuserListCmd)
	dbuserCmd.AddCommand(dbuserCreateCmd)
	dbuserCmd.AddCommand(dbuserUpdateCmd)
	dbuserCmd.AddCommand(dbuserModifyCmd)
	dbuserCmd.AddCommand(dbuserRemoveCmd)
	dbuserCmd.AddCommand(dbuserDeleteCmd)
	rootCmd.AddCommand(dbuserCmd)
}
