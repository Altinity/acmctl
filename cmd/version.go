package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "0.1.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print acmctl version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("acmctl version %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
