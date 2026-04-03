package cmd

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/output"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:     "node",
	Aliases: []string{"nd"},
	Short:   "Manage nodes",
}

var nodeGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get node details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		node, err := apiClient.GetNode(args[0])
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, node)
	},
}

var nodeRestartCmd = &cobra.Command{
	Use:   "restart <id>",
	Short: "Restart a node",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		hard, _ := cmd.Flags().GetBool("hard")
		if err := apiClient.RestartNode(args[0], hard); err != nil {
			return err
		}
		fmt.Printf("Node %s restart initiated.\n", args[0])
		return nil
	},
}

var nodeStatusCmd = &cobra.Command{
	Use:   "status <id>",
	Short: "Get node status",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetNodeStatus(args[0])
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

var nodeMetricsCmd = &cobra.Command{
	Use:   "metrics <id>",
	Short: "Get node metrics",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		detailed, _ := cmd.Flags().GetBool("detailed")
		result, err := apiClient.GetNodeMetrics(args[0], detailed)
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

func init() {
	nodeRestartCmd.Flags().Bool("hard", false, "hard restart")
	nodeMetricsCmd.Flags().Bool("detailed", false, "show detailed metrics")

	nodeCmd.AddCommand(nodeGetCmd)
	nodeCmd.AddCommand(nodeRestartCmd)
	nodeCmd.AddCommand(nodeStatusCmd)
	nodeCmd.AddCommand(nodeMetricsCmd)
	rootCmd.AddCommand(nodeCmd)
}
