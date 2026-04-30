package cmd

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/output"
	"github.com/spf13/cobra"
)

var nodetypeCmd = &cobra.Command{
	Use:   "nodetype",
	Short: "Manage node types",
}

var nodetypeCreateCmd = &cobra.Command{
	Use:   "create <env-id>",
	Short: "Add a node type to an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := nodeTypeFlags(cmd)
		nt, err := apiClient.CreateNodeType(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, nt)
	},
}

var nodetypeUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Modify a node type",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := nodeTypeFlags(cmd)
		nt, err := apiClient.UpdateNodeType(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, nt)
	},
}

var nodetypeDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Remove a node type",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteNodeType(args[0]); err != nil {
			return err
		}
		fmt.Printf("Node type %s deleted.\n", args[0])
		return nil
	},
}

func nodeTypeFlags(cmd *cobra.Command) map[string]string {
	return flagsToParams(cmd, map[string]string{
		"name": "name", "scope": "scope", "code": "code",
		"memory": "memory", "cpu": "cpu", "storage-class": "storageClass",
		"capacity": "capacity", "extra-spec": "extraSpec",
		"node-selector": "nodeSelector", "tolerations": "tolerations",
		"is-spot": "isSpot",
	})
}

func init() {
	for _, c := range []*cobra.Command{nodetypeCreateCmd, nodetypeUpdateCmd} {
		for _, name := range []string{"name", "scope", "code", "memory", "cpu",
			"storage-class", "capacity", "extra-spec", "node-selector", "tolerations", "is-spot"} {
			c.Flags().String(name, "", "")
		}
	}

	nodetypeCmd.AddCommand(nodetypeCreateCmd)
	nodetypeCmd.AddCommand(nodetypeUpdateCmd)
	nodetypeCmd.AddCommand(nodetypeDeleteCmd)
	rootCmd.AddCommand(nodetypeCmd)
}
