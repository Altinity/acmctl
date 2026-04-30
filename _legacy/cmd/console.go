package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Internal console: audit logs, system tasks, settings",
}

var consoleInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show background-operations info",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetConsoleInfo()
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var consoleLogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Show internal audit log across environments",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"page": "page", "limit": "limit", "filter": "filter", "order": "order",
		})
		result, err := apiClient.GetConsoleLogs(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var consoleEnvLogsCmd = &cobra.Command{
	Use:   "env-logs <env-id>",
	Short: "Tail logs for a cluster inside an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"count": "count", "cluster": "cluster", "filter": "filter",
			"download": "download", "section": "section",
		})
		result, err := apiClient.GetEnvironmentConsoleLogs(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var consoleFilteredLogsCmd = &cobra.Command{
	Use:   "logs-filtered",
	Short: "Get logs for a given set of labels",
	RunE: func(cmd *cobra.Command, args []string) error {
		labels, _ := cmd.Flags().GetString("labels")
		result, err := apiClient.GetFilteredConsoleLogs(labels)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var consoleSettingsGetCmd = &cobra.Command{
	Use:   "settings",
	Short: "Get system operations settings",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetConsoleSettings()
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var consoleSettingsUpdateCmd = &cobra.Command{
	Use:   "settings-update",
	Short: "Update system operations settings",
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectFieldFlags(cmd)
		if err != nil {
			return err
		}
		for k, v := range flagsToParams(cmd, map[string]string{
			"max-processes": "maxProcesses", "max-backup-processes": "maxBackupProcesses",
			"max-cluster-processes": "maxClusterProcesses",
			"auto-charge-percent-condition": "autoChargePercentCondition",
			"current-cron-host": "currentCronHost",
		}) {
			params[k] = v
		}
		result, err := apiClient.UpdateConsoleSettings(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var consoleTasksListCmd = &cobra.Command{
	Use:   "tasks",
	Short: "List system background tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"page": "page", "limit": "limit", "filter": "filter", "order": "order",
		})
		result, err := apiClient.ListConsoleTasks(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var consoleTaskCreateCmd = &cobra.Command{
	Use:   "task-create",
	Short: "Schedule a new background task",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{"action": "action", "data": "data"})
		if v, ok := params["data"]; ok {
			resolved, err := resolveValue(v)
			if err != nil {
				return err
			}
			params["data"] = resolved
		}
		result, err := apiClient.ScheduleConsoleTask(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var consoleTaskUpdateCmd = &cobra.Command{
	Use:   "task-update <task-id>",
	Short: "Update a running system task (e.g. interrupt)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{"interrupt": "interrupt"})
		result, err := apiClient.UpdateConsoleTask(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var consoleTaskDeleteCmd = &cobra.Command{
	Use:   "task-delete <task-id>",
	Short: "Remove a scheduled task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteConsoleTask(args[0]); err != nil {
			return err
		}
		fmt.Printf("Task %s deleted.\n", args[0])
		return nil
	},
}

func init() {
	for _, c := range []*cobra.Command{consoleLogsCmd, consoleTasksListCmd} {
		c.Flags().String("page", "", "page number")
		c.Flags().String("limit", "", "page size")
		c.Flags().String("filter", "", "filter (JSON)")
		c.Flags().String("order", "", "order (JSON)")
	}

	for _, name := range []string{"count", "cluster", "filter", "download", "section"} {
		consoleEnvLogsCmd.Flags().String(name, "", "")
	}

	consoleFilteredLogsCmd.Flags().String("labels", "", "labels expression")

	for _, name := range []string{"max-processes", "max-backup-processes", "max-cluster-processes",
		"auto-charge-percent-condition", "current-cron-host"} {
		consoleSettingsUpdateCmd.Flags().String(name, "", "")
	}
	consoleSettingsUpdateCmd.Flags().StringSliceP("field", "F", nil, "key=value (repeatable)")

	consoleTaskCreateCmd.Flags().String("action", "", "task action")
	consoleTaskCreateCmd.Flags().String("data", "", "task data (use @file or @- for stdin)")

	consoleTaskUpdateCmd.Flags().String("interrupt", "", "interrupt flag")

	for _, c := range []*cobra.Command{
		consoleInfoCmd, consoleLogsCmd, consoleEnvLogsCmd, consoleFilteredLogsCmd,
		consoleSettingsGetCmd, consoleSettingsUpdateCmd,
		consoleTasksListCmd, consoleTaskCreateCmd, consoleTaskUpdateCmd, consoleTaskDeleteCmd,
	} {
		consoleCmd.AddCommand(c)
	}
	rootCmd.AddCommand(consoleCmd)
}
