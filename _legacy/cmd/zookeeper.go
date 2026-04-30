package cmd

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/output"
	"github.com/spf13/cobra"
)

var zookeeperCmd = &cobra.Command{
	Use:     "zookeeper",
	Aliases: []string{"zk"},
	Short:   "Manage Zookeeper clusters",
}

var zookeeperAddCmd = &cobra.Command{
	Use:   "add <env-id>",
	Short: "Add an existing Zookeeper cluster to an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"tag": "tag", "hosts": "hosts", "suffix": "suffix",
		})
		zk, err := apiClient.AddZookeeperCluster(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, zk)
	},
}

var zookeeperLaunchCmd = &cobra.Command{
	Use:   "launch <env-id>",
	Short: "Launch a new Zookeeper cluster in an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"tag": "tag", "size": "size", "node-type": "nodeType",
		})
		zk, err := apiClient.LaunchZookeeperCluster(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, zk)
	},
}

var zookeeperUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Modify a Zookeeper cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{"tag": "tag", "hosts": "hosts"})
		zk, err := apiClient.UpdateZookeeperCluster(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, zk)
	},
}

var zookeeperDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Remove a Zookeeper cluster from its environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteZookeeperCluster(args[0]); err != nil {
			return err
		}
		fmt.Printf("Zookeeper %s deleted.\n", args[0])
		return nil
	},
}

var zookeeperPushCmd = &cobra.Command{
	Use:   "push <id>",
	Short: "Publish Zookeeper configuration to the cloud provider",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.PushZookeeperCluster(args[0]); err != nil {
			return err
		}
		fmt.Printf("Zookeeper %s push initiated.\n", args[0])
		return nil
	},
}

var zookeeperRescaleCmd = &cobra.Command{
	Use:   "rescale <id>",
	Short: "Rescale a Zookeeper cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{"size": "size", "node-type": "nodeType"})
		if err := apiClient.RescaleZookeeperCluster(args[0], params); err != nil {
			return err
		}
		fmt.Printf("Zookeeper %s rescale initiated.\n", args[0])
		return nil
	},
}

var zookeeperStatusCmd = &cobra.Command{
	Use:   "status <id>",
	Short: "Check Zookeeper cluster status",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetZookeeperStatus(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

func init() {
	for _, name := range []string{"tag", "hosts", "suffix"} {
		zookeeperAddCmd.Flags().String(name, "", "")
	}
	for _, name := range []string{"tag", "size", "node-type"} {
		zookeeperLaunchCmd.Flags().String(name, "", "")
	}
	for _, name := range []string{"tag", "hosts"} {
		zookeeperUpdateCmd.Flags().String(name, "", "")
	}
	zookeeperRescaleCmd.Flags().String("size", "", "new size")
	zookeeperRescaleCmd.Flags().String("node-type", "", "node type")

	for _, c := range []*cobra.Command{
		zookeeperAddCmd, zookeeperLaunchCmd, zookeeperUpdateCmd, zookeeperDeleteCmd,
		zookeeperPushCmd, zookeeperRescaleCmd, zookeeperStatusCmd,
	} {
		zookeeperCmd.AddCommand(c)
	}
	rootCmd.AddCommand(zookeeperCmd)
}
