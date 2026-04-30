package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Inspection / introspection commands

var clusterMetricsCmd = &cobra.Command{
	Use:   "metrics <id>",
	Short: "Get cluster metrics",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetClusterMetrics(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterAuditCmd = &cobra.Command{
	Use:   "audit <id>",
	Short: "List cluster audit log",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"page": "page", "limit": "limit", "filter": "filter", "order": "order",
		})
		result, err := apiClient.GetClusterAudit(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterBackupConfigModsCmd = &cobra.Command{
	Use:   "backup-config-modifications <id>",
	Short: "List advanced backup settings",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetClusterBackupConfigModifications(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterCloneDbTasksCmd = &cobra.Command{
	Use:   "clone-database-tasks <id>",
	Short: "Show running clone-database tasks",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ListCloneDatabaseTasks(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterConsistencyCmd = &cobra.Command{
	Use:   "consistency <id>",
	Short: "Get cluster consistency report",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetClusterConsistency(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterCrashesCmd = &cobra.Command{
	Use:   "crashes <id>",
	Short: "List cluster crashes",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		node, _ := cmd.Flags().GetString("node")
		result, err := apiClient.GetClusterCrashes(args[0], node)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterDataDistributionCmd = &cobra.Command{
	Use:   "data-distribution <id>",
	Short: "Show current data distribution",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		node, _ := cmd.Flags().GetString("node")
		result, err := apiClient.GetClusterDataDistribution(args[0], node)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterRebalanceQueriesCmd = &cobra.Command{
	Use:   "rebalance-queries <id>",
	Short: "List available data rebalance queries",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		node, _ := cmd.Flags().GetString("node")
		result, err := apiClient.GetDataRebalanceQueries(args[0], node)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterDataTransferLogsCmd = &cobra.Command{
	Use:   "data-transfer-logs <id>",
	Short: "Get data transfer logs",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		limit, _ := cmd.Flags().GetString("limit")
		result, err := apiClient.GetDataTransferLogs(args[0], limit)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterDetachedPartsCmd = &cobra.Command{
	Use:   "detached-parts <id>",
	Short: "List detached parts",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		node, _ := cmd.Flags().GetString("node")
		result, err := apiClient.GetClusterDetachedParts(args[0], node)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterErrorsCmd = &cobra.Command{
	Use:   "errors <id>",
	Short: "List cluster errors",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		node, _ := cmd.Flags().GetString("node")
		result, err := apiClient.GetClusterErrors(args[0], node)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterExportCmd = &cobra.Command{
	Use:   "export <id>",
	Short: "Export cluster configuration to JSON/YAML",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("format")
		result, err := apiClient.ExportCluster(args[0], format)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterKafkaConfigsCmd = &cobra.Command{
	Use:   "kafka-configurations <id>",
	Short: "Get existing Kafka configurations for the cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetKafkaConfigurations(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterKafkaTablesCmd = &cobra.Command{
	Use:   "kafka-tables <id>",
	Short: "List kafka tables",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		node, _ := cmd.Flags().GetString("node")
		result, err := apiClient.GetKafkaTables(args[0], node)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterLayoutsCmd = &cobra.Command{
	Use:   "layouts <id>",
	Short: "List cluster layouts",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cluster, _ := cmd.Flags().GetString("cluster")
		result, err := apiClient.GetClusterLayouts(args[0], cluster)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterActionLogCmd = &cobra.Command{
	Use:   "action-log <id>",
	Short: "List recent cluster actions",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{"all": "all", "undo": "undo"})
		result, err := apiClient.GetClusterActionLog(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterQueriesCmd = &cobra.Command{
	Use:   "queries <id>",
	Short: "List running queries",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		node, _ := cmd.Flags().GetString("node")
		result, err := apiClient.GetClusterQueries(args[0], node)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterSystemTablesCmd = &cobra.Command{
	Use:   "system-tables <id>",
	Short: "List system tables that can be cleaned up",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		node, _ := cmd.Flags().GetString("node")
		result, err := apiClient.GetClusterSystemTables(args[0], node)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterTableCmd = &cobra.Command{
	Use:   "table <id>",
	Short: "Describe a table",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"database": "database", "table": "table", "node": "node",
		})
		result, err := apiClient.DescribeClusterTable(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterTablePartitionsCmd = &cobra.Command{
	Use:   "table-partitions <id>",
	Short: "List partitions for a table",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"database": "database", "table": "table",
		})
		result, err := apiClient.GetTablePartitions(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterUnusedTablesCmd = &cobra.Command{
	Use:   "unused-tables <id>",
	Short: "List unused tables",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"node": "node", "not-queried-days": "notQueriedDays",
		})
		result, err := apiClient.GetUnusedTables(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterWorkloadKafkaCmd = &cobra.Command{
	Use:   "workload-kafka <id>",
	Short: "Show kafka workload",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		node, _ := cmd.Flags().GetString("node")
		result, err := apiClient.GetWorkloadKafka(args[0], node)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterWorkloadMutationsCmd = &cobra.Command{
	Use:   "workload-mutations <id>",
	Short: "Show mutation/replication queue",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		node, _ := cmd.Flags().GetString("node")
		result, err := apiClient.GetWorkloadMutations(args[0], node)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterWorkloadQueriesCmd = &cobra.Command{
	Use:   "workload-queries <id>",
	Short: "Show running queries (workload view)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"node": "node", "order-by": "orderBy", "order-direction": "orderDirection",
			"limit": "limit", "timeframe": "timeframe",
		})
		result, err := apiClient.GetWorkloadQueries(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterWorkloadQueryCmd = &cobra.Command{
	Use:   "workload-query <id> <query-id>",
	Short: "Get details for a specific query",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"node": "node", "timeframe": "timeframe",
		})
		result, err := apiClient.GetWorkloadQuery(args[0], args[1], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterWorkloadReplicationCmd = &cobra.Command{
	Use:   "workload-replication <id>",
	Short: "Show replication queue",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"node": "node", "order-by": "orderBy", "order-direction": "orderDirection",
		})
		result, err := apiClient.GetWorkloadReplication(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterTablesSchemaCmd = &cobra.Command{
	Use:   "tables-schema",
	Short: "Check the database schema of a cluster (cluster-less endpoint)",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"cluster": "cluster", "user": "user", "pass": "pass",
			"name": "name", "system": "system", "keywords": "keywords",
		})
		result, err := apiClient.CheckTablesSchema(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

// Lifecycle / data ops

var clusterValidateMigrationCmd = &cobra.Command{
	Use:   "validate-migration",
	Short: "Validate cluster authentication and free space for migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"user": "user", "password": "password", "cluster": "cluster",
			"name": "name", "type": "type",
		})
		result, err := apiClient.ValidateMigration(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterBackupManualCmd = &cobra.Command{
	Use:   "backup-manual <id>",
	Short: "Mark backups for a cluster as manual",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tags, _ := cmd.Flags().GetString("tags")
		if err := apiClient.MarkBackupManual(args[0], tags); err != nil {
			return err
		}
		fmt.Printf("Backups marked manual for %s.\n", args[0])
		return nil
	},
}

var clusterCloneCmd = &cobra.Command{
	Use:   "clone <id>",
	Short: "Clone a cluster (full)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		result, err := apiClient.CloneCluster(args[0], name)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterCloneDatabaseCmd = &cobra.Command{
	Use:   "clone-database <id>",
	Short: "Clone a database within a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"node": "node", "database": "database", "clone-name": "cloneName",
		})
		result, err := apiClient.CloneDatabase(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterConvertReplicatedCmd = &cobra.Command{
	Use:   "convert-to-replicated <id>",
	Short: "Convert MergeTree tables to replicated",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tables, _ := cmd.Flags().GetString("tables")
		if err := apiClient.ConvertToReplicated(args[0], tables); err != nil {
			return err
		}
		fmt.Printf("Cluster %s convert-to-replicated initiated.\n", args[0])
		return nil
	},
}

var clusterDatalakeCmd = &cobra.Command{
	Use:   "datalake <id>",
	Short: "Create a Data Lake catalog database",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"database": "database", "type": "type",
			"catalog": "catalog", "settings": "settings",
		})
		if v, ok := params["settings"]; ok {
			resolved, err := resolveValue(v)
			if err != nil {
				return err
			}
			params["settings"] = resolved
		}
		result, err := apiClient.CreateDatalake(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterResetDefaultsCmd = &cobra.Command{
	Use:   "reset-defaults <id> [profile]",
	Short: "Reset cluster or profile settings to defaults",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		profile := ""
		if len(args) == 2 {
			profile = args[1]
		}
		if err := apiClient.ResetDefaults(args[0], profile); err != nil {
			return err
		}
		fmt.Printf("Defaults reset for cluster %s.\n", args[0])
		return nil
	},
}

var clusterDeleteDetachedCmd = &cobra.Command{
	Use:   "delete-detached-parts <id>",
	Short: "Delete detached parts",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		parts, _ := cmd.Flags().GetString("parts")
		if err := apiClient.DeleteDetachedParts(args[0], parts); err != nil {
			return err
		}
		fmt.Printf("Detached parts deleted on %s.\n", args[0])
		return nil
	},
}

var clusterDisableSystemTablesCmd = &cobra.Command{
	Use:   "disable-system-tables <id>",
	Short: "Disable system tables",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tables, _ := cmd.Flags().GetString("tables")
		if err := apiClient.DisableSystemTables(args[0], tables); err != nil {
			return err
		}
		fmt.Printf("System tables disabled on %s.\n", args[0])
		return nil
	},
}

var clusterDropTablesCmd = &cobra.Command{
	Use:   "drop-tables <id>",
	Short: "Drop tables",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tables, _ := cmd.Flags().GetString("tables")
		if err := apiClient.DropTables(args[0], tables); err != nil {
			return err
		}
		fmt.Printf("Tables dropped on %s.\n", args[0])
		return nil
	},
}

var clusterEnableSystemTablesCmd = &cobra.Command{
	Use:   "enable-system-tables <id>",
	Short: "Enable system tables",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tables, _ := cmd.Flags().GetString("tables")
		if err := apiClient.EnableSystemTables(args[0], tables); err != nil {
			return err
		}
		fmt.Printf("System tables enabled on %s.\n", args[0])
		return nil
	},
}

var clusterImportDatasetCmd = &cobra.Command{
	Use:   "import-dataset <id>",
	Short: "Import a dataset",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"dataset": "dataset", "schemas": "schemas",
		})
		result, err := apiClient.ImportDataset(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterInterruptRebalanceCmd = &cobra.Command{
	Use:   "interrupt-rebalance <id>",
	Short: "Interrupt running data rebalance",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.InterruptRebalance(args[0]); err != nil {
			return err
		}
		fmt.Printf("Rebalance interrupted on %s.\n", args[0])
		return nil
	},
}

var clusterKafkaCheckCmd = &cobra.Command{
	Use:   "kafka-check <id>",
	Short: "Check Kafka broker connection",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"broker": "broker", "topic": "topic", "options": "options",
			"file": "file", "config": "config",
			"new-file": "newFile", "new-config": "newConfig",
		})
		result, err := apiClient.KafkaCheck(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterKafkaConfigSaveCmd = &cobra.Command{
	Use:   "kafka-config-save <id>",
	Short: "Save Kafka configuration file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filename, _ := cmd.Flags().GetString("filename")
		xml, _ := cmd.Flags().GetString("xml")
		resolvedXML, err := resolveValue(xml)
		if err != nil {
			return err
		}
		params := map[string]string{}
		if filename != "" {
			params["filename"] = filename
		}
		if resolvedXML != "" {
			params["xml"] = resolvedXML
		}
		result, err := apiClient.SaveKafkaConfiguration(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterKafkaConfigDeleteCmd = &cobra.Command{
	Use:   "kafka-config-delete <id> <file> <config>",
	Short: "Remove a Kafka configuration",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteKafkaConfiguration(args[0], args[1], args[2]); err != nil {
			return err
		}
		fmt.Printf("Kafka config %s/%s deleted on cluster %s.\n", args[1], args[2], args[0])
		return nil
	},
}

var clusterKafkaTablesRestartCmd = &cobra.Command{
	Use:   "kafka-tables-restart <id>",
	Short: "Restart kafka tables",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tables, _ := cmd.Flags().GetString("tables")
		if err := apiClient.RestartKafkaTables(args[0], tables); err != nil {
			return err
		}
		fmt.Printf("Kafka tables restarted on %s.\n", args[0])
		return nil
	},
}

var clusterMigrateVolumesCmd = &cobra.Command{
	Use:   "migrate-volumes <id>",
	Short: "Migrate cluster volumes",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.MigrateClusterVolumes(args[0]); err != nil {
			return err
		}
		fmt.Printf("Volume migration initiated on %s.\n", args[0])
		return nil
	},
}

var clusterMutationKillCmd = &cobra.Command{
	Use:   "mutation-kill <id>",
	Short: "Kill mutations",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		mutations, _ := cmd.Flags().GetString("mutations")
		if err := apiClient.MutationKill(args[0], mutations); err != nil {
			return err
		}
		fmt.Printf("Mutations killed on %s.\n", args[0])
		return nil
	},
}

var clusterQueryKillCmd = &cobra.Command{
	Use:   "query-kill <id>",
	Short: "Kill query processes",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"query-ids": "queryIds", "node": "node",
		})
		if err := apiClient.QueryKill(args[0], params); err != nil {
			return err
		}
		fmt.Printf("Query-kill initiated on %s.\n", args[0])
		return nil
	},
}

var clusterRebalanceDataCmd = &cobra.Command{
	Use:   "rebalance-data <id>",
	Short: "Run available data rebalance queries",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		node, _ := cmd.Flags().GetString("node")
		if err := apiClient.RebalanceData(args[0], node); err != nil {
			return err
		}
		fmt.Printf("Rebalance started on %s.\n", args[0])
		return nil
	},
}

var clusterRestartTableReplicaCmd = &cobra.Command{
	Use:   "restart-table-replica <id>",
	Short: "Restart a table replica",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"host": "host", "database": "database", "table": "table",
		})
		if err := apiClient.RestartTableReplica(args[0], params); err != nil {
			return err
		}
		fmt.Printf("Table replica restart initiated on %s.\n", args[0])
		return nil
	},
}

var clusterRestoreRetryCmd = &cobra.Command{
	Use:   "restore-retry <id>",
	Short: "Retry a cluster restore",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		skip, _ := cmd.Flags().GetString("skip-download")
		if err := apiClient.RetryRestore(args[0], skip); err != nil {
			return err
		}
		fmt.Printf("Restore retry initiated on %s.\n", args[0])
		return nil
	},
}

var clusterRollbackCmd = &cobra.Command{
	Use:   "rollback <id>",
	Short: "Roll back the cluster to a previous action",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		action, _ := cmd.Flags().GetString("action")
		if err := apiClient.RollbackCluster(args[0], action); err != nil {
			return err
		}
		fmt.Printf("Rollback initiated on %s.\n", args[0])
		return nil
	},
}

var clusterSetSysTableTTLCmd = &cobra.Command{
	Use:   "set-system-table-ttl <id>",
	Short: "Set TTL for a system table",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		table, _ := cmd.Flags().GetString("table")
		ttl, _ := cmd.Flags().GetString("ttl")
		if err := apiClient.SetSystemTableTTL(args[0], table, ttl); err != nil {
			return err
		}
		fmt.Printf("TTL set for %s.\n", args[0])
		return nil
	},
}

var clusterSwarmEnableCmd = &cobra.Command{
	Use:   "swarm-enable <id>",
	Short: "Enable swarm discovery",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.EnableSwarm(args[0]); err != nil {
			return err
		}
		fmt.Printf("Swarm enabled on %s.\n", args[0])
		return nil
	},
}

var clusterSyncSchemaCmd = &cobra.Command{
	Use:   "sync-schema <id>",
	Short: "Create an existing table on all nodes",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"database": "database", "table": "table",
		})
		if err := apiClient.SyncSchema(args[0], params); err != nil {
			return err
		}
		fmt.Printf("Schema sync initiated on %s.\n", args[0])
		return nil
	},
}

var clusterTruncateTablesCmd = &cobra.Command{
	Use:   "truncate-tables <id>",
	Short: "Truncate tables",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tables, _ := cmd.Flags().GetString("tables")
		if err := apiClient.TruncateTables(args[0], tables); err != nil {
			return err
		}
		fmt.Printf("Tables truncated on %s.\n", args[0])
		return nil
	},
}

var clusterDataTransferCmd = &cobra.Command{
	Use:   "data-transfer",
	Short: "Start a data migration job from one cluster to another",
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectFieldFlags(cmd)
		if err != nil {
			return err
		}
		for k, v := range flagsToParams(cmd, map[string]string{
			"source": "source", "name": "name", "target": "target",
		}) {
			params[k] = v
		}
		result, err := apiClient.StartDataTransfer(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var clusterLockedCmd = &cobra.Command{
	Use:   "locked <id>",
	Short: "Modify cluster locked status",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		locked, _ := cmd.Flags().GetString("locked")
		if err := apiClient.PatchClusterLocked(args[0], locked); err != nil {
			return err
		}
		fmt.Printf("Cluster %s locked status updated.\n", args[0])
		return nil
	},
}

var clusterStoragePolicyCmd = &cobra.Command{
	Use:   "storage-policy <id>",
	Short: "Modify cluster storage policy",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		policy, _ := cmd.Flags().GetString("storage-policy")
		moveFactor, _ := cmd.Flags().GetString("move-factor")
		if err := apiClient.PatchClusterStoragePolicy(args[0], policy, moveFactor); err != nil {
			return err
		}
		fmt.Printf("Cluster %s storage policy updated.\n", args[0])
		return nil
	},
}

var nodeLBExclusionCmd = &cobra.Command{
	Use:   "lb-exclusion <node-id>",
	Short: "Set node load-balancer exclusion status",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		exclude, _ := cmd.Flags().GetString("exclude")
		if err := apiClient.SetNodeLBExclusion(args[0], exclude); err != nil {
			return err
		}
		fmt.Printf("Node %s LB exclusion updated.\n", args[0])
		return nil
	},
}

func init() {
	for _, c := range []*cobra.Command{
		clusterCrashesCmd, clusterDataDistributionCmd, clusterRebalanceQueriesCmd,
		clusterDetachedPartsCmd, clusterErrorsCmd, clusterKafkaTablesCmd,
		clusterQueriesCmd, clusterSystemTablesCmd, clusterWorkloadKafkaCmd,
		clusterWorkloadMutationsCmd, clusterRebalanceDataCmd,
	} {
		c.Flags().String("node", "", "target node")
	}

	clusterAuditCmd.Flags().String("page", "", "page number")
	clusterAuditCmd.Flags().String("limit", "", "page size")
	clusterAuditCmd.Flags().String("filter", "", "filter (JSON)")
	clusterAuditCmd.Flags().String("order", "", "order (JSON)")

	clusterDataTransferLogsCmd.Flags().String("limit", "", "max records")
	clusterExportCmd.Flags().String("format", "", "json or yaml")
	clusterLayoutsCmd.Flags().String("cluster", "", "filter by cluster")

	clusterActionLogCmd.Flags().String("all", "", "include all")
	clusterActionLogCmd.Flags().String("undo", "", "undoable only")

	for _, c := range []*cobra.Command{clusterTableCmd, clusterTablePartitionsCmd} {
		c.Flags().String("database", "", "database")
		c.Flags().String("table", "", "table")
	}
	clusterTableCmd.Flags().String("node", "", "target node")

	clusterUnusedTablesCmd.Flags().String("node", "", "target node")
	clusterUnusedTablesCmd.Flags().String("not-queried-days", "", "days threshold")

	for _, name := range []string{"node", "order-by", "order-direction", "limit", "timeframe"} {
		clusterWorkloadQueriesCmd.Flags().String(name, "", "")
	}
	for _, name := range []string{"node", "timeframe"} {
		clusterWorkloadQueryCmd.Flags().String(name, "", "")
	}
	for _, name := range []string{"node", "order-by", "order-direction"} {
		clusterWorkloadReplicationCmd.Flags().String(name, "", "")
	}

	for _, name := range []string{"cluster", "user", "pass", "name", "system", "keywords"} {
		clusterTablesSchemaCmd.Flags().String(name, "", "")
	}

	for _, name := range []string{"user", "password", "cluster", "name", "type"} {
		clusterValidateMigrationCmd.Flags().String(name, "", "")
	}

	clusterBackupManualCmd.Flags().String("tags", "", "tags to mark")
	clusterCloneCmd.Flags().String("name", "", "cloned cluster name")

	clusterCloneDatabaseCmd.Flags().String("node", "", "node")
	clusterCloneDatabaseCmd.Flags().String("database", "", "database to clone")
	clusterCloneDatabaseCmd.Flags().String("clone-name", "", "new database name")

	clusterConvertReplicatedCmd.Flags().String("tables", "", "tables to convert")

	clusterDatalakeCmd.Flags().String("database", "", "database")
	clusterDatalakeCmd.Flags().String("type", "", "catalog type")
	clusterDatalakeCmd.Flags().String("catalog", "", "catalog")
	clusterDatalakeCmd.Flags().String("settings", "", "settings (use @file or @- for stdin)")

	clusterDeleteDetachedCmd.Flags().String("parts", "", "parts list")
	for _, c := range []*cobra.Command{clusterDisableSystemTablesCmd, clusterDropTablesCmd, clusterEnableSystemTablesCmd, clusterTruncateTablesCmd, clusterKafkaTablesRestartCmd} {
		c.Flags().String("tables", "", "tables list")
	}

	clusterImportDatasetCmd.Flags().String("dataset", "", "dataset")
	clusterImportDatasetCmd.Flags().String("schemas", "", "schemas")

	for _, name := range []string{"broker", "topic", "options", "file", "config", "new-file", "new-config"} {
		clusterKafkaCheckCmd.Flags().String(name, "", "")
	}

	clusterKafkaConfigSaveCmd.Flags().String("filename", "", "filename")
	clusterKafkaConfigSaveCmd.Flags().String("xml", "", "XML content (use @file or @- for stdin)")

	clusterMutationKillCmd.Flags().String("mutations", "", "mutation IDs")
	clusterQueryKillCmd.Flags().String("query-ids", "", "query IDs")
	clusterQueryKillCmd.Flags().String("node", "", "target node")

	for _, name := range []string{"host", "database", "table"} {
		clusterRestartTableReplicaCmd.Flags().String(name, "", "")
	}

	clusterRestoreRetryCmd.Flags().String("skip-download", "", "skip download")
	clusterRollbackCmd.Flags().String("action", "", "action ID to roll back to")

	clusterSetSysTableTTLCmd.Flags().String("table", "", "table")
	clusterSetSysTableTTLCmd.Flags().String("ttl", "", "TTL value")

	clusterSyncSchemaCmd.Flags().String("database", "", "database")
	clusterSyncSchemaCmd.Flags().String("table", "", "table")

	clusterDataTransferCmd.Flags().String("source", "", "source cluster")
	clusterDataTransferCmd.Flags().String("name", "", "transfer name")
	clusterDataTransferCmd.Flags().String("target", "", "target cluster")
	clusterDataTransferCmd.Flags().StringSliceP("field", "F", nil, "key=value (repeatable)")

	clusterLockedCmd.Flags().String("locked", "", "true/false")
	clusterStoragePolicyCmd.Flags().String("storage-policy", "", "policy")
	clusterStoragePolicyCmd.Flags().String("move-factor", "", "move factor")

	nodeLBExclusionCmd.Flags().String("exclude", "", "true/false")

	for _, c := range []*cobra.Command{
		clusterMetricsCmd, clusterAuditCmd, clusterBackupConfigModsCmd,
		clusterCloneDbTasksCmd, clusterConsistencyCmd,
		clusterCrashesCmd, clusterDataDistributionCmd, clusterRebalanceQueriesCmd,
		clusterDataTransferLogsCmd, clusterDetachedPartsCmd, clusterErrorsCmd,
		clusterExportCmd, clusterKafkaConfigsCmd, clusterKafkaTablesCmd,
		clusterLayoutsCmd, clusterActionLogCmd, clusterQueriesCmd,
		clusterSystemTablesCmd, clusterTableCmd, clusterTablePartitionsCmd,
		clusterUnusedTablesCmd, clusterWorkloadKafkaCmd, clusterWorkloadMutationsCmd,
		clusterWorkloadQueriesCmd, clusterWorkloadQueryCmd, clusterWorkloadReplicationCmd,
		clusterTablesSchemaCmd,
		clusterValidateMigrationCmd, clusterBackupManualCmd, clusterCloneCmd,
		clusterCloneDatabaseCmd, clusterConvertReplicatedCmd, clusterDatalakeCmd,
		clusterResetDefaultsCmd, clusterDeleteDetachedCmd, clusterDisableSystemTablesCmd,
		clusterDropTablesCmd, clusterEnableSystemTablesCmd, clusterImportDatasetCmd,
		clusterInterruptRebalanceCmd, clusterKafkaCheckCmd, clusterKafkaConfigSaveCmd,
		clusterKafkaConfigDeleteCmd, clusterKafkaTablesRestartCmd,
		clusterMigrateVolumesCmd, clusterMutationKillCmd, clusterQueryKillCmd,
		clusterRebalanceDataCmd, clusterRestartTableReplicaCmd, clusterRestoreRetryCmd,
		clusterRollbackCmd, clusterSetSysTableTTLCmd, clusterSwarmEnableCmd,
		clusterSyncSchemaCmd, clusterTruncateTablesCmd, clusterDataTransferCmd,
		clusterLockedCmd, clusterStoragePolicyCmd,
	} {
		clusterCmd.AddCommand(c)
	}

	nodeCmd.AddCommand(nodeLBExclusionCmd)
}
