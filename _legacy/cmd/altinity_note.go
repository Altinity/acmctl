package cmd

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/output"
	"github.com/spf13/cobra"
)

var altinityNoteCmd = &cobra.Command{
	Use:     "altinity-note",
	Aliases: []string{"note"},
	Short:   "Manage Altinity notes attached to environments",
}

var altinityNoteListCmd = &cobra.Command{
	Use:   "list <env-id>",
	Short: "List Altinity notes for an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ListAltinityNotes(args[0])
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

var altinityNoteCreateCmd = &cobra.Command{
	Use:   "create <env-id>",
	Short: "Create an Altinity note for an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectFieldFlags(cmd)
		if err != nil {
			return err
		}
		for k, v := range flagsToParams(cmd, map[string]string{
			"severity": "severity", "title": "title", "message": "message",
		}) {
			params[k] = v
		}
		result, err := apiClient.CreateAltinityNote(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, result)
	},
}

var altinityNoteUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update an Altinity note",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectFieldFlags(cmd)
		if err != nil {
			return err
		}
		for k, v := range flagsToParams(cmd, map[string]string{
			"severity": "severity", "title": "title", "message": "message",
		}) {
			params[k] = v
		}
		result, err := apiClient.UpdateAltinityNote(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, result)
	},
}

var altinityNoteDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete an Altinity note",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteAltinityNote(args[0]); err != nil {
			return err
		}
		fmt.Printf("Note %s deleted.\n", args[0])
		return nil
	},
}

func init() {
	for _, c := range []*cobra.Command{altinityNoteCreateCmd, altinityNoteUpdateCmd} {
		c.Flags().String("severity", "", "severity (e.g. info, warning, critical)")
		c.Flags().String("title", "", "note title")
		c.Flags().String("message", "", "note message (use @file or @- for stdin)")
		c.Flags().StringSliceP("field", "F", nil, "key=value (repeatable)")
	}

	altinityNoteCmd.AddCommand(altinityNoteListCmd)
	altinityNoteCmd.AddCommand(altinityNoteCreateCmd)
	altinityNoteCmd.AddCommand(altinityNoteUpdateCmd)
	altinityNoteCmd.AddCommand(altinityNoteDeleteCmd)
	rootCmd.AddCommand(altinityNoteCmd)
}
