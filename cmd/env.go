package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:     "env",
	Aliases: []string{"environment"},
	Short:   "Inspect environments",
}

var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all environments",
	RunE: func(cmd *cobra.Command, args []string) error {
		var raw json.RawMessage
		if err := apiClient.Do("GET", "/environments", nil, &raw); err != nil {
			return err
		}
		return printJSON(raw)
	},
}

var envGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get environment details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var raw json.RawMessage
		if err := apiClient.Do("GET", fmt.Sprintf("/environment/%s", args[0]), nil, &raw); err != nil {
			return err
		}
		return printJSON(raw)
	},
}

func init() {
	envCmd.AddCommand(envListCmd)
	envCmd.AddCommand(envGetCmd)
	rootCmd.AddCommand(envCmd)
}
