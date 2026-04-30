package cmd

import (
	"github.com/spf13/cobra"
)

var cloudCmd = &cobra.Command{
	Use:   "cloud",
	Short: "Cloud option discovery and GraphQL proxy",
}

var cloudOptionsCmd = &cobra.Command{
	Use:   "options",
	Short: "List cloud options for a platform",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"platform": "platform", "type": "type", "region": "region", "provider": "provider",
		})
		result, err := apiClient.GetCloudOptions(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var cloudEnvOptionsCmd = &cobra.Command{
	Use:   "env-options <env-id>",
	Short: "List cloud options for a specific environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"platform": "platform", "type": "type", "region": "region", "provider": "provider",
		})
		result, err := apiClient.GetEnvironmentCloudOptions(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var cloudQueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Proxy a GraphQL query to the cloud controller",
	RunE: func(cmd *cobra.Command, args []string) error {
		query, _ := cmd.Flags().GetString("query")
		variables, _ := cmd.Flags().GetString("variables")
		resolvedQuery, err := resolveValue(query)
		if err != nil {
			return err
		}
		resolvedVars, err := resolveValue(variables)
		if err != nil {
			return err
		}
		params := map[string]string{}
		if resolvedQuery != "" {
			params["query"] = resolvedQuery
		}
		if resolvedVars != "" {
			params["variables"] = resolvedVars
		}
		result, err := apiClient.CloudQuery(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

func init() {
	for _, c := range []*cobra.Command{cloudOptionsCmd, cloudEnvOptionsCmd} {
		c.Flags().String("platform", "", "platform")
		c.Flags().String("type", "", "option type")
		c.Flags().String("region", "", "region")
		c.Flags().String("provider", "", "provider")
	}

	cloudQueryCmd.Flags().String("query", "", "GraphQL query (use @file or @- for stdin)")
	cloudQueryCmd.Flags().String("variables", "", "variables JSON (use @file or @- for stdin)")

	cloudCmd.AddCommand(cloudOptionsCmd)
	cloudCmd.AddCommand(cloudEnvOptionsCmd)
	cloudCmd.AddCommand(cloudQueryCmd)
	rootCmd.AddCommand(cloudCmd)
}
