package cmd

import (
	"fmt"

	"github.com/altinity/acmctl/pkg/output"
	"github.com/spf13/cobra"
)

var clusterCmd = &cobra.Command{
	Use:     "cluster",
	Aliases: []string{"cl"},
	Short:   "Manage clusters",
}

var clusterListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all clusters",
	RunE: func(cmd *cobra.Command, args []string) error {
		clusters, err := apiClient.ListClusters()
		if err != nil {
			return err
		}
		return output.PrintTabulableList(cfg.Output, clusters)
	},
}

var clusterGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get cluster details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cluster, err := apiClient.GetCluster(args[0])
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, cluster)
	},
}

var clusterDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		terminate, _ := cmd.Flags().GetBool("terminate")
		if err := apiClient.DeleteCluster(args[0], terminate); err != nil {
			return err
		}
		fmt.Printf("Cluster %s deleted.\n", args[0])
		return nil
	},
}

var clusterLaunchCmd = &cobra.Command{
	Use:   "launch <env-id>",
	Short: "Launch a new cluster in an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]string{}
		flagMap := map[string]string{
			"name": "name", "type": "type", "nodes": "nodes", "shards": "shards",
			"replicas": "replicas", "node-type": "nodeType", "version": "version",
			"version-image": "versionImage", "size": "size", "disks": "disks",
			"admin-pass": "adminPass", "memory": "memory", "storage-class": "storageClass",
			"throughput": "throughput", "iops": "iops", "lb-type": "lbType",
			"secure": "secure", "uptime": "uptime", "timezone": "timezone", "role": "role",
			"zookeeper": "zookeeper", "zookeeper-options": "zookeeperOptions",
			"host": "host", "port": "port", "ssh-port": "sshPort", "http-port": "httpPort",
			"data-path": "dataPath", "source-cluster": "sourceCluster",
			"replicate-schema": "replicateSchema", "backup-source": "backupSource",
			"ip-whitelist": "ipWhitelist", "restore-options": "restoreOptions",
			"backup-options": "backupOptions", "datadog-settings": "datadogSettings",
			"uptime-settings": "uptimeSettings", "alternate-endpoints": "alternateEndpoints",
			"annotations": "annotations", "azlist": "azlist",
			"custom-lb-annotations": "customLBAnnotations",
			"chguard-settings": "chGuardSettings",
			"mysql-protocol": "mysqlProtocol", "mysql-port": "mysqlPort",
			"zone-awareness": "zoneAwareness",
		}
		for flag, key := range flagMap {
			if v, _ := cmd.Flags().GetString(flag); v != "" {
				params[key] = v
			}
		}
		cluster, err := apiClient.LaunchCluster(args[0], params)
		if err != nil {
			return err
		}
		return output.Print(cfg.Output, cluster)
	},
}

var clusterRestartCmd = &cobra.Command{
	Use:   "restart <id>",
	Short: "Restart a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		method, _ := cmd.Flags().GetString("method")
		if err := apiClient.RestartCluster(args[0], method); err != nil {
			return err
		}
		fmt.Printf("Cluster %s restart initiated.\n", args[0])
		return nil
	},
}

var clusterStopCmd = &cobra.Command{
	Use:   "stop <id>",
	Short: "Stop a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.StopCluster(args[0]); err != nil {
			return err
		}
		fmt.Printf("Cluster %s stop initiated.\n", args[0])
		return nil
	},
}

var clusterResumeCmd = &cobra.Command{
	Use:   "resume <id>",
	Short: "Resume a stopped cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]string{}
		if v, _ := cmd.Flags().GetString("node-type"); v != "" {
			params["nodeType"] = v
		}
		if v, _ := cmd.Flags().GetString("shards"); v != "" {
			params["shards"] = v
		}
		if v, _ := cmd.Flags().GetString("replicas"); v != "" {
			params["replicas"] = v
		}
		if err := apiClient.ResumeCluster(args[0], params); err != nil {
			return err
		}
		fmt.Printf("Cluster %s resume initiated.\n", args[0])
		return nil
	},
}

var clusterUpgradeCmd = &cobra.Command{
	Use:   "upgrade <id>",
	Short: "Upgrade cluster ClickHouse version",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]string{}
		if v, _ := cmd.Flags().GetString("version"); v != "" {
			params["version"] = v
		}
		if v, _ := cmd.Flags().GetString("type"); v != "" {
			params["type"] = v
		}
		if v, _ := cmd.Flags().GetString("image"); v != "" {
			params["image"] = v
		}
		if err := apiClient.UpgradeCluster(args[0], params); err != nil {
			return err
		}
		fmt.Printf("Cluster %s upgrade initiated.\n", args[0])
		return nil
	},
}

var clusterRescaleCmd = &cobra.Command{
	Use:   "rescale <id>",
	Short: "Rescale a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]string{}
		flagMap := map[string]string{
			"node-type": "nodeType", "shards": "shards", "replicas": "replicas",
			"size": "size", "disks": "disks", "throughput": "throughput",
		}
		for flag, key := range flagMap {
			if v, _ := cmd.Flags().GetString(flag); v != "" {
				params[key] = v
			}
		}
		if err := apiClient.RescaleCluster(args[0], params); err != nil {
			return err
		}
		fmt.Printf("Cluster %s rescale initiated.\n", args[0])
		return nil
	},
}

var clusterBackupCmd = &cobra.Command{
	Use:   "backup <id>",
	Short: "Trigger a cluster backup",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.BackupCluster(args[0]); err != nil {
			return err
		}
		fmt.Printf("Cluster %s backup initiated.\n", args[0])
		return nil
	},
}

var clusterRestoreCmd = &cobra.Command{
	Use:   "restore <id>",
	Short: "Restore a cluster from backup",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]string{}
		flagMap := map[string]string{
			"type": "type", "provider": "provider", "access-key": "accessKey",
			"secret-key": "secretKey", "region": "region", "bucket": "bucket",
			"path": "path", "tag": "tag", "table": "table",
		}
		for flag, key := range flagMap {
			if v, _ := cmd.Flags().GetString(flag); v != "" {
				params[key] = v
			}
		}
		if err := apiClient.RestoreCluster(args[0], params); err != nil {
			return err
		}
		fmt.Printf("Cluster %s restore initiated.\n", args[0])
		return nil
	},
}

var clusterQueryCmd = &cobra.Command{
	Use:   "query <id>",
	Short: "Execute a SQL query on a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]string{}
		if v, _ := cmd.Flags().GetString("query"); v != "" {
			params["query"] = v
		}
		if v, _ := cmd.Flags().GetString("node"); v != "" {
			params["node"] = v
		}
		if v, _ := cmd.Flags().GetString("user"); v != "" {
			params["user"] = v
		}
		if v, _ := cmd.Flags().GetString("password"); v != "" {
			params["password"] = v
		}
		result, err := apiClient.QueryCluster(args[0], params)
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

var clusterNodesCmd = &cobra.Command{
	Use:   "nodes <id>",
	Short: "List nodes in a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		nodes, err := apiClient.ListClusterNodes(args[0])
		if err != nil {
			return err
		}
		return output.PrintTabulableList(cfg.Output, nodes)
	},
}

var clusterStatusCmd = &cobra.Command{
	Use:   "status <id>",
	Short: "Get cluster status",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetClusterStatus(args[0])
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

var clusterBackupsCmd = &cobra.Command{
	Use:   "backups <id>",
	Short: "List cluster backups",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ListClusterBackups(args[0])
		if err != nil {
			return err
		}
		return output.Print("json", result)
	},
}

var clusterPushCmd = &cobra.Command{
	Use:   "push <id>",
	Short: "Push cluster configuration",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.PushCluster(args[0]); err != nil {
			return err
		}
		fmt.Printf("Cluster %s push initiated.\n", args[0])
		return nil
	},
}

func init() {
	// delete
	clusterDeleteCmd.Flags().Bool("terminate", false, "terminate cluster resources")

	// launch
	for _, f := range []string{"name", "type", "nodes", "shards", "replicas", "node-type", "version",
		"version-image", "size", "disks", "admin-pass", "memory", "storage-class",
		"throughput", "iops", "lb-type", "secure", "uptime", "timezone", "role",
		"zookeeper", "zookeeper-options",
		"host", "port", "ssh-port", "http-port", "data-path",
		"source-cluster", "replicate-schema", "backup-source",
		"ip-whitelist", "restore-options", "backup-options", "datadog-settings",
		"uptime-settings", "alternate-endpoints", "annotations", "azlist",
		"custom-lb-annotations", "chguard-settings",
		"mysql-protocol", "mysql-port", "zone-awareness"} {
		clusterLaunchCmd.Flags().String(f, "", "")
	}
	_ = clusterLaunchCmd.MarkFlagRequired("name")

	// restart
	clusterRestartCmd.Flags().String("method", "", "restart method")

	// resume
	clusterResumeCmd.Flags().String("node-type", "", "node type")
	clusterResumeCmd.Flags().String("shards", "", "number of shards")
	clusterResumeCmd.Flags().String("replicas", "", "number of replicas")

	// upgrade
	clusterUpgradeCmd.Flags().String("version", "", "target ClickHouse version")
	clusterUpgradeCmd.Flags().String("type", "", "upgrade type")
	clusterUpgradeCmd.Flags().String("image", "", "custom image")

	// rescale
	for _, f := range []string{"node-type", "shards", "replicas", "size", "disks", "throughput"} {
		clusterRescaleCmd.Flags().String(f, "", "")
	}

	// restore
	for _, f := range []string{"type", "provider", "access-key", "secret-key", "region", "bucket", "path", "tag", "table"} {
		clusterRestoreCmd.Flags().String(f, "", "")
	}

	// query
	clusterQueryCmd.Flags().StringP("query", "q", "", "SQL query to execute")
	clusterQueryCmd.Flags().String("node", "", "target node")
	clusterQueryCmd.Flags().String("user", "", "ClickHouse user")
	clusterQueryCmd.Flags().String("password", "", "ClickHouse password")

	clusterCmd.AddCommand(clusterListCmd)
	clusterCmd.AddCommand(clusterGetCmd)
	clusterCmd.AddCommand(clusterDeleteCmd)
	clusterCmd.AddCommand(clusterLaunchCmd)
	clusterCmd.AddCommand(clusterRestartCmd)
	clusterCmd.AddCommand(clusterStopCmd)
	clusterCmd.AddCommand(clusterResumeCmd)
	clusterCmd.AddCommand(clusterUpgradeCmd)
	clusterCmd.AddCommand(clusterRescaleCmd)
	clusterCmd.AddCommand(clusterBackupCmd)
	clusterCmd.AddCommand(clusterRestoreCmd)
	clusterCmd.AddCommand(clusterQueryCmd)
	clusterCmd.AddCommand(clusterNodesCmd)
	clusterCmd.AddCommand(clusterStatusCmd)
	clusterCmd.AddCommand(clusterBackupsCmd)
	clusterCmd.AddCommand(clusterPushCmd)
	rootCmd.AddCommand(clusterCmd)
}
