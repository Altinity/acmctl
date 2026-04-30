package cmd

import (
	"github.com/altinity/acmctl/pkg/output"
	"github.com/spf13/cobra"
)

var auditReportCmd = &cobra.Command{
	Use:     "audit-report",
	Aliases: []string{"audit"},
	Short:   "Manage cluster audit reports",
}

var auditReportGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get an audit report by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetAuditReport(args[0])
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

var auditReportListCmd = &cobra.Command{
	Use:   "list <cluster-id>",
	Short: "List audit reports for a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ListClusterAuditReports(args[0])
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

var auditReportCreateCmd = &cobra.Command{
	Use:   "create <cluster-id>",
	Short: "Create an audit report for a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.CreateClusterAuditReport(args[0])
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

func init() {
	auditReportCmd.AddCommand(auditReportGetCmd)
	auditReportCmd.AddCommand(auditReportListCmd)
	auditReportCmd.AddCommand(auditReportCreateCmd)
	rootCmd.AddCommand(auditReportCmd)
}
