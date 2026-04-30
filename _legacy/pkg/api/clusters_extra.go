package api

import "fmt"

// Inspection / introspection

func (c *Client) GetClusterMetrics(clusterID string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/metrics", clusterID), nil, &result)
	return result, err
}

func (c *Client) GetClusterAudit(clusterID string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/audit", clusterID), params, &result)
	return result, err
}

func (c *Client) GetClusterBackupConfigModifications(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/backup-config-modifications", id), nil, &result)
	return result, err
}

func (c *Client) ListCloneDatabaseTasks(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/clone-database", id), nil, &result)
	return result, err
}

func (c *Client) GetClusterConsistency(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/consistency", id), nil, &result)
	return result, err
}

func (c *Client) GetClusterCrashes(id, node string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if node != "" {
		params["node"] = node
	}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/crashes", id), params, &result)
	return result, err
}

func (c *Client) GetClusterDataDistribution(id, node string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if node != "" {
		params["node"] = node
	}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/data-distribution", id), params, &result)
	return result, err
}

func (c *Client) GetDataRebalanceQueries(id, node string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if node != "" {
		params["node"] = node
	}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/data-rebalance-queries", id), params, &result)
	return result, err
}

func (c *Client) GetDataTransferLogs(id, limit string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if limit != "" {
		params["limit"] = limit
	}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/data-transfer-logs", id), params, &result)
	return result, err
}

func (c *Client) GetClusterDetachedParts(id, node string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if node != "" {
		params["node"] = node
	}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/detached-parts", id), params, &result)
	return result, err
}

func (c *Client) GetClusterErrors(id, node string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if node != "" {
		params["node"] = node
	}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/errors", id), params, &result)
	return result, err
}

func (c *Client) ExportCluster(id, format string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if format != "" {
		params["format"] = format
	}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/export", id), params, &result)
	return result, err
}

func (c *Client) GetKafkaConfigurations(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/kafka-configurations", id), nil, &result)
	return result, err
}

func (c *Client) GetKafkaTables(id, node string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if node != "" {
		params["node"] = node
	}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/kafka-tables", id), params, &result)
	return result, err
}

func (c *Client) GetClusterLayouts(id, cluster string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if cluster != "" {
		params["cluster"] = cluster
	}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/layouts", id), params, &result)
	return result, err
}

func (c *Client) GetClusterActionLog(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/log", id), params, &result)
	return result, err
}

func (c *Client) GetClusterQueries(id, node string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if node != "" {
		params["node"] = node
	}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/queries", id), params, &result)
	return result, err
}

func (c *Client) GetClusterSystemTables(id, node string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if node != "" {
		params["node"] = node
	}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/system-tables", id), params, &result)
	return result, err
}

func (c *Client) DescribeClusterTable(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/table", id), params, &result)
	return result, err
}

func (c *Client) GetTablePartitions(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/table-partitions", id), params, &result)
	return result, err
}

func (c *Client) GetUnusedTables(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/unused-tables", id), params, &result)
	return result, err
}

func (c *Client) GetWorkloadKafka(id, node string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if node != "" {
		params["node"] = node
	}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/workload-kafka", id), params, &result)
	return result, err
}

func (c *Client) GetWorkloadMutations(id, node string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if node != "" {
		params["node"] = node
	}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/workload-mutations", id), params, &result)
	return result, err
}

func (c *Client) GetWorkloadQueries(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/workload-queries", id), params, &result)
	return result, err
}

func (c *Client) GetWorkloadQuery(id, queryID string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/workload-query/%s", id, queryID), params, &result)
	return result, err
}

func (c *Client) GetWorkloadReplication(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/cluster/%s/workload-replication", id), params, &result)
	return result, err
}

func (c *Client) CheckTablesSchema(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/tables", params, &result)
	return result, err
}

// Lifecycle / data ops

func (c *Client) ValidateMigration(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/cluster/validate-migration", params, &result)
	return result, err
}

func (c *Client) MarkBackupManual(id, tags string) error {
	params := map[string]string{}
	if tags != "" {
		params["tags"] = tags
	}
	return c.Do("POST", fmt.Sprintf("/cluster/%s/backup/manual", id), params, nil)
}

func (c *Client) CloneCluster(id, name string) (interface{}, error) {
	var result interface{}
	params := map[string]string{}
	if name != "" {
		params["name"] = name
	}
	err := c.Do("POST", fmt.Sprintf("/cluster/%s/clone", id), params, &result)
	return result, err
}

func (c *Client) CloneDatabase(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/cluster/%s/clone-database", id), params, &result)
	return result, err
}

func (c *Client) ConvertToReplicated(id, tables string) error {
	params := map[string]string{}
	if tables != "" {
		params["tables"] = tables
	}
	return c.Do("POST", fmt.Sprintf("/cluster/%s/convert-to-replicated", id), params, nil)
}

func (c *Client) CreateDatalake(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("POST", fmt.Sprintf("/cluster/%s/datalake", id), params, &result)
	return result, err
}

func (c *Client) ResetDefaults(id, profile string) error {
	if profile == "" {
		return c.Do("POST", fmt.Sprintf("/cluster/%s/defaults", id), nil, nil)
	}
	return c.Do("POST", fmt.Sprintf("/cluster/%s/defaults/%s", id, profile), nil, nil)
}

func (c *Client) DeleteDetachedParts(id, parts string) error {
	params := map[string]string{}
	if parts != "" {
		params["parts"] = parts
	}
	return c.Do("POST", fmt.Sprintf("/cluster/%s/delete-detached-parts", id), params, nil)
}

func (c *Client) DisableSystemTables(id, tables string) error {
	params := map[string]string{}
	if tables != "" {
		params["tables"] = tables
	}
	return c.Do("POST", fmt.Sprintf("/cluster/%s/disable-system-tables", id), params, nil)
}

func (c *Client) DropTables(id, tables string) error {
	params := map[string]string{}
	if tables != "" {
		params["tables"] = tables
	}
	return c.Do("POST", fmt.Sprintf("/cluster/%s/drop-tables", id), params, nil)
}

func (c *Client) EnableSystemTables(id, tables string) error {
	params := map[string]string{}
	if tables != "" {
		params["tables"] = tables
	}
	return c.Do("POST", fmt.Sprintf("/cluster/%s/enable-system-tables", id), params, nil)
}

func (c *Client) ImportDataset(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/cluster/%s/import-dataset", id), params, &result)
	return result, err
}

func (c *Client) InterruptRebalance(id string) error {
	return c.Do("POST", fmt.Sprintf("/cluster/%s/interrupt-rebalance-data", id), nil, nil)
}

func (c *Client) KafkaCheck(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("POST", fmt.Sprintf("/cluster/%s/kafka-check", id), params, &result)
	return result, err
}

func (c *Client) SaveKafkaConfiguration(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.DoForm("POST", fmt.Sprintf("/cluster/%s/kafka-configuration", id), params, &result)
	return result, err
}

func (c *Client) DeleteKafkaConfiguration(id, file, config string) error {
	return c.Do("DELETE", fmt.Sprintf("/cluster/%s/kafka-configuration/%s/%s", id, file, config), nil, nil)
}

func (c *Client) RestartKafkaTables(id, tables string) error {
	params := map[string]string{}
	if tables != "" {
		params["tables"] = tables
	}
	return c.Do("POST", fmt.Sprintf("/cluster/%s/kafka-tables-restart", id), params, nil)
}

func (c *Client) MigrateClusterVolumes(id string) error {
	return c.Do("POST", fmt.Sprintf("/cluster/%s/migrate-volumes", id), nil, nil)
}

func (c *Client) MutationKill(id, mutations string) error {
	params := map[string]string{}
	if mutations != "" {
		params["mutations"] = mutations
	}
	return c.Do("POST", fmt.Sprintf("/cluster/%s/mutation-kill", id), params, nil)
}

func (c *Client) QueryKill(id string, params map[string]string) error {
	return c.Do("POST", fmt.Sprintf("/cluster/%s/query-kill", id), params, nil)
}

func (c *Client) RebalanceData(id, node string) error {
	params := map[string]string{}
	if node != "" {
		params["node"] = node
	}
	return c.Do("POST", fmt.Sprintf("/cluster/%s/rebalance-data", id), params, nil)
}

func (c *Client) RestartTableReplica(id string, params map[string]string) error {
	return c.Do("POST", fmt.Sprintf("/cluster/%s/restart-table-replica", id), params, nil)
}

func (c *Client) RetryRestore(id, skipDownload string) error {
	params := map[string]string{}
	if skipDownload != "" {
		params["skipDownload"] = skipDownload
	}
	return c.Do("POST", fmt.Sprintf("/cluster/%s/restore-retry", id), params, nil)
}

func (c *Client) RollbackCluster(id, action string) error {
	params := map[string]string{}
	if action != "" {
		params["action"] = action
	}
	return c.Do("POST", fmt.Sprintf("/cluster/%s/rollback", id), params, nil)
}

func (c *Client) SetSystemTableTTL(id, table, ttl string) error {
	params := map[string]string{}
	if table != "" {
		params["table"] = table
	}
	if ttl != "" {
		params["ttl"] = ttl
	}
	return c.Do("POST", fmt.Sprintf("/cluster/%s/set-system-table-ttl", id), params, nil)
}

func (c *Client) EnableSwarm(id string) error {
	return c.Do("POST", fmt.Sprintf("/cluster/%s/swarm/enable", id), nil, nil)
}

func (c *Client) SyncSchema(id string, params map[string]string) error {
	return c.Do("POST", fmt.Sprintf("/cluster/%s/sync-schema", id), params, nil)
}

func (c *Client) TruncateTables(id, tables string) error {
	params := map[string]string{}
	if tables != "" {
		params["tables"] = tables
	}
	return c.Do("POST", fmt.Sprintf("/cluster/%s/truncate-tables", id), params, nil)
}

func (c *Client) StartDataTransfer(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/data-transfer", params, &result)
	return result, err
}

func (c *Client) PatchClusterLocked(id, locked string) error {
	params := map[string]string{}
	if locked != "" {
		params["locked"] = locked
	}
	return c.Do("PATCH", fmt.Sprintf("/cluster/%s/locked", id), params, nil)
}

func (c *Client) PatchClusterStoragePolicy(id, storagePolicy, moveFactor string) error {
	params := map[string]string{}
	if storagePolicy != "" {
		params["storagePolicy"] = storagePolicy
	}
	if moveFactor != "" {
		params["moveFactor"] = moveFactor
	}
	return c.Do("PATCH", fmt.Sprintf("/cluster/%s/storage-policy", id), params, nil)
}

func (c *Client) SetNodeLBExclusion(nodeID, exclude string) error {
	params := map[string]string{}
	if exclude != "" {
		params["exclude"] = exclude
	}
	return c.Do("POST", fmt.Sprintf("/node/%s/lb-exclusion", nodeID), params, nil)
}
