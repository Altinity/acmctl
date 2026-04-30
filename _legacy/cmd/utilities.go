package cmd

import "github.com/spf13/cobra"

var datasetsCmd = &cobra.Command{
	Use:   "datasets",
	Short: "List available datasets",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ListDatasets()
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Get Prometheus monitoring metrics",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetMetrics()
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var referenceCmd = &cobra.Command{
	Use:   "reference",
	Short: "Get the API specification (reference.json)",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetReferenceSpec()
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get system status",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetStatus()
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

func init() {
	rootCmd.AddCommand(datasetsCmd)
	rootCmd.AddCommand(metricsCmd)
	rootCmd.AddCommand(referenceCmd)
	rootCmd.AddCommand(statusCmd)
}
