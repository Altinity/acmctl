package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var notificationCmd = &cobra.Command{
	Use:     "notification",
	Aliases: []string{"notif"},
	Short:   "Manage notifications",
}

var notificationUserListCmd = &cobra.Command{
	Use:   "user-list",
	Short: "List user notifications",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"all": "all", "page": "page", "limit": "limit", "filter": "filter", "order": "order",
		})
		result, err := apiClient.ListUserNotifications(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var notificationListCmd = &cobra.Command{
	Use:   "list",
	Short: "List admin notifications",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"sticky": "sticky", "page": "page", "limit": "limit",
			"filter": "filter", "order": "order",
		})
		result, err := apiClient.ListAdminNotifications(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var notificationCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an admin notification",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := notificationFlags(cmd)
		result, err := apiClient.CreateAdminNotification(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var notificationGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a notification",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetNotification(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var notificationUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Modify an admin notification",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := notificationFlags(cmd)
		result, err := apiClient.UpdateNotification(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var notificationDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Remove an admin notification",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteNotification(args[0]); err != nil {
			return err
		}
		fmt.Printf("Notification %s deleted.\n", args[0])
		return nil
	},
}

var notificationAckCmd = &cobra.Command{
	Use:   "ack <id>",
	Short: "Acknowledge a notification",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.AckNotification(args[0]); err != nil {
			return err
		}
		fmt.Printf("Notification %s acknowledged.\n", args[0])
		return nil
	},
}

func notificationFlags(cmd *cobra.Command) map[string]string {
	return flagsToParams(cmd, map[string]string{
		"message": "message", "level": "level", "recipients": "recipients",
		"expiry": "expiry", "sticky": "sticky", "send": "send",
		"channel-email": "channelEmail", "channel-popup": "channelPopup",
	})
}

func init() {
	for _, c := range []*cobra.Command{notificationUserListCmd, notificationListCmd} {
		for _, f := range []string{"all", "sticky", "page", "limit", "filter", "order"} {
			c.Flags().String(f, "", "")
		}
	}

	for _, c := range []*cobra.Command{notificationCreateCmd, notificationUpdateCmd} {
		for _, f := range []string{"message", "level", "recipients", "expiry",
			"sticky", "send", "channel-email", "channel-popup"} {
			c.Flags().String(f, "", "")
		}
	}

	for _, c := range []*cobra.Command{
		notificationUserListCmd, notificationListCmd, notificationCreateCmd,
		notificationGetCmd, notificationUpdateCmd, notificationDeleteCmd, notificationAckCmd,
	} {
		notificationCmd.AddCommand(c)
	}
	rootCmd.AddCommand(notificationCmd)
}
