# Clusters

DELETE /cluster/{id}/kafka-configuration/{file}/{config} — Removes Kafka configuration from Cluster
DELETE /cluster/{id}/{terminate} — Removes a cluster descriptior from the environment
GET /cluster/{cluster}/logs — Displays the system log tails of the given cluster [count, node, filter, download, section]
GET /cluster/{cluster}/metrics — Returns cluster metrics
GET /cluster/{cluster}/status — Returns cluster uptime status
GET /cluster/{id} — Gets cluster's complete information [schema]
GET /cluster/{id}/audit — List all audit logs for the given cluster [page, limit, filter, order]
GET /cluster/{id}/backup-config-modifications — Lists advanced backup settings
GET /cluster/{id}/backups — Returns the list of backups for a given cluster [schedule]
GET /cluster/{id}/clone-database — Display currently running clone database tasks
GET /cluster/{id}/consistency — Returns consistency report for cluster
GET /cluster/{id}/crashes — Lists crashes for the given cluster [node]
GET /cluster/{id}/data-distribution — Returns current data distribution on cluster [node]
GET /cluster/{id}/data-rebalance-queries — Returns list of available data rebalance queries or "in_progress" if rebalance is in progress [node]
GET /cluster/{id}/data-transfer-logs — Returns data transfer logs [limit]
GET /cluster/{id}/detached-parts — Returns a list of detached parts [node]
GET /cluster/{id}/errors — Lists errors for the given cluster [node]
GET /cluster/{id}/export — Exports cluster configuration options into JSON/YAML file [format]
GET /cluster/{id}/kafka-configurations — Tries to find existing configuration for Cluster
GET /cluster/{id}/kafka-tables — List kafka tables for the given cluster [node]
GET /cluster/{id}/layouts — Lists the cluster' layouts [cluster]
GET /cluster/{id}/log — List recent actions done with the cluster [all, undo]
GET /cluster/{id}/nodes — List the cluster's nodes [cluster]
GET /cluster/{id}/queries — List all running queries at the given cluster [node]
GET /cluster/{id}/support/credentials — Generates a temporary user with Altinity permissions
GET /cluster/{id}/system-tables — Returns list of system tables that can possibly be cleaned up [node]
GET /cluster/{id}/table — Describes the table of a given Cluster [database, table, node]
GET /cluster/{id}/table-partitions — Lists partitions for the given table [database, table]
GET /cluster/{id}/unused-tables — Returns list of unused tables that can possibly be cleaned up [node, notQueriedDays]
GET /cluster/{id}/workload-kafka — List kafka workload for the given cluster [node]
GET /cluster/{id}/workload-mutations — List replication queue for the given cluster [node]
GET /cluster/{id}/workload-queries — List all running queries at the given cluster [node, orderBy, orderDirection, limit, timeframe]
GET /cluster/{id}/workload-query/{queryId} — List details for given query [node, timeframe]
GET /cluster/{id}/workload-replication — List replication queue for the given cluster [node, orderBy, orderDirection]
GET /clusters — Lists all existing clusters [essentialsOnly]
GET /environment/{environment}/clusters — Lists available clusters of the given environment [withSettings]
GET /node/{id} — Gets information about given cluster node
GET /node/{id}/metrics — Gets metrics information about given cluster node [detailed]
GET /node/{id}/status — Returns node's uptime, health status and LB exclusion info
GET /tables — Checks cluster's database schema [cluster, user, pass, name, system, keywords]
PATCH /cluster/{id}/locked — Modifies cluster locked status [locked]
PATCH /cluster/{id}/storage-policy — Modifies cluster storage policy [storagePolicy, moveFactor]
POST /cluster/validate-migration — Returns cluster authentication data validity and free space [user, password, cluster, name, type]
POST /cluster/{id} — Modifies a cluster general information [name, nodes, startupTime, troubleshootingMode, ipWhitelist, alertsEmail, id_parent, alertsSettings, datadogSettings, uptime, uptimeSettings, alternateEndpoints, endpointsEnabled, annotations, backupOptions, backupSource, id_owner, role, disableZoneAwareness, lbType, customLBAnnotations, chGuardSettings, mysqlProtocol, mysqlPort, zkRoot, backupConfigModifications, actionSchedule, timezone, zoneAwareness]
POST /cluster/{id}/backup — Creates a data backup of the given cluster
POST /cluster/{id}/backup/manual — Sets backups for a given cluster as manual [tags]
POST /cluster/{id}/clone — Clones cluster in full, including all settings [name]
POST /cluster/{id}/clone-database — Clone database for the given cluster [node, database, cloneName]
POST /cluster/{id}/convert-to-replicated — Converts a MergeTree type cluster to replicated [tables]
POST /cluster/{id}/datalake — Executes CREATE database query for Data Lake catalog [database, type, catalog, settings]
POST /cluster/{id}/defaults{profile}) — Resets cluster or profile settings to defaults
POST /cluster/{id}/delete-detached-parts — Deletes detached parts [parts]
POST /cluster/{id}/disable-system-tables — Disables system tables [tables]
POST /cluster/{id}/drop-tables — Drops tables [tables]
POST /cluster/{id}/enable-system-tables — Enables system tables [tables]
POST /cluster/{id}/import-dataset — Imports a dataset [dataset, schemas]
POST /cluster/{id}/interrupt-rebalance-data — Interrupts running data rebalance queries
POST /cluster/{id}/kafka-check — Checks connection details to Kafka broker [broker, topic, options, file, config, newFile, newConfig]
POST /cluster/{id}/kafka-configuration — Saves Kafka configuration file to Cluster, overwriting existing one [filename, xml]
POST /cluster/{id}/kafka-tables-restart — Restarts kafka tables for the given cluster [tables]
POST /cluster/{id}/migrate-volumes — Migrates volumes for the cluster
POST /cluster/{id}/mutation-kill — Kills certain mutations in the given cluster [mutations]
POST /cluster/{id}/push — Publishes all recent configuration changes onto nodes of the given cluster [update-hosts]
POST /cluster/{id}/query — Executes a query on cluster (or a particular cluster node) [cluster, node, query, ddl, layout, user, password, timeout, swarm]
POST /cluster/{id}/query-kill — Kills query processes for a particular node of the given cluster [queryIds, node]
POST /cluster/{id}/rebalance-data — Runs available data rebalance queries [node]
POST /cluster/{id}/restart-table-replica — Restarts table replica for the given table [host, database, table]
POST /cluster/{id}/restore — Restores a cluster data from a backup [type, provider, accessKey, secretKey, region, bucket, path, tag, table, sourceEnvironment, tableFilterMode, tableFilterParam, engineFilterMode, engineFilterParam, skipEngines, skipDownload]
POST /cluster/{id}/restore-retry — Restores a cluster data from a backup [skipDownload]
POST /cluster/{id}/rollback — Rolls back the cluster to the state from the given action [action]
POST /cluster/{id}/set-system-table-ttl — Sets TTL for the given table [table, ttl]
POST /cluster/{id}/support/refresh — Refreshes Altinity support access
POST /cluster/{id}/swarm/enable — Enables swarm discovery to a given cluster
POST /cluster/{id}/sync-schema — Creates an existing table on all nodes [database, table]
POST /cluster/{id}/truncate-tables — Truncates tables [tables]
POST /data-transfer — Starts a data migration job from one cluster to another [source, name, target, tables, credentials]
POST /environment/{environment}/clusters — Adds a cluster descriptor [name, nodes, id_parent]
POST /environment/{environment}/clusters/import — Imports cluster configuration options from JSON file [skipEndpoints]
POST /environment/{environment}/clusters/launch — Launches a cluster inside given cloud environment [name, type, nodes, shards, replicas, host, port, sshPort, httpPort, zookeeper, zookeeperOptions, dataPath, nodeType, version, versionImage, size, disks, adminPass, memory, storageClass, throughput, iops, lbType, secure, sourceCluster, replicateSchema, backupSource, ipWhitelist, restoreOptions, backupOptions, datadogSettings, uptime, uptimeSettings, alternateEndpoints, annotations, azlist, role, customLBAnnotations, chGuardSettings, mysqlProtocol, mysqlPort, timezone, zoneAwareness]
POST /environment/{environment}/clusters/restore — Restores the cluster data and its configuration (optional) from a given external storage [type, provider, accessKey, secretKey, region, bucket, path, tag, cluster, restoreName, nodeType, version, versionImage, storageClass, size, disks, includeNodeVolumes, sourceEnvironment, tableFilterMode, tableFilterParam, engineFilterMode, engineFilterParam, skipEngines, skipDownload, dry]
POST /node/{id}/lb-exclusion — Sets node's LB exclusion status [exclude]
PUT /cluster/{id}/rescale — Adds/Removes a Shard/Replica of the given cluster [nodeType, shards, replicas, size, disks, onlyNew, throughput, forceNonReplicated, azlist]
PUT /cluster/{id}/restart — Restarts the cluster [method]
PUT /cluster/{id}/resume — Resumes previously suspended cluster [nodeType, shards, replicas]
PUT /cluster/{id}/stop — Suspends a given cluster (including all running nodes)
PUT /cluster/{id}/upgrade — Upgrades a Clickhouse version upon the cluster [version, type, image]
PUT /node/{id}/restart — Restarts the node [hard]
