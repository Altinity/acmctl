package cmd

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/output"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:     "env",
	Aliases: []string{"environment"},
	Short:   "Manage environments",
}

var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all environments",
	RunE: func(cmd *cobra.Command, args []string) error {
		envs, err := apiClient.ListEnvironments()
		if err != nil {
			return err
		}
		return output.PrintTabulableList(cfg.Output, envs)
	},
}

var envGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get environment details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		env, err := apiClient.GetEnvironment(args[0])
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, env)
	},
}

var envDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteEnvironment(args[0]); err != nil {
			return err
		}
		fmt.Printf("Environment %s deleted.\n", args[0])
		return nil
	},
}

var envClustersCmd = &cobra.Command{
	Use:   "clusters <env-id>",
	Short: "List clusters in an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		clusters, err := apiClient.ListEnvironmentClusters(args[0])
		if err != nil {
			return err
		}
		return output.PrintTabulableList(cfg.Output, clusters)
	},
}

var envNodeTypesCmd = &cobra.Command{
	Use:   "nodetypes <env-id>",
	Short: "List node types in an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		nodeTypes, err := apiClient.ListEnvironmentNodeTypes(args[0])
		if err != nil {
			return err
		}
		return output.PrintTabulableList(cfg.Output, nodeTypes)
	},
}

var envZookeepersCmd = &cobra.Command{
	Use:   "zookeepers <env-id>",
	Short: "List zookeeper clusters in an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		zks, err := apiClient.ListEnvironmentZookeepers(args[0])
		if err != nil {
			return err
		}
		return output.PrintTabulableList(cfg.Output, zks)
	},
}

func init() {
	envCmd.AddCommand(envListCmd)
	envCmd.AddCommand(envGetCmd)
	envCmd.AddCommand(envDeleteCmd)
	envCmd.AddCommand(envClustersCmd)
	envCmd.AddCommand(envNodeTypesCmd)
	envCmd.AddCommand(envZookeepersCmd)
	rootCmd.AddCommand(envCmd)
}
