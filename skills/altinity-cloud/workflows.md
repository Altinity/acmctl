# ACM Multi-Step Workflows

Assumes `acmctl` is authenticated and available on `PATH`.

## Launch a cluster and wait until ready

```bash
ENV=$(acmctl env list | jq '.[0].id')

CLUSTER=$(cat <<'JSON' | acmctl cluster launch "$ENV" | jq -r '.id'
{ "name": "test-cluster", "nodeType": "s1", "version": "24.3" }
JSON
)

# Poll status until uptime > 0 (means it's accepting connections)
until [ "$(acmctl raw GET /cluster/$CLUSTER/status | jq '.uptime // 0')" -gt 0 ]; do
  sleep 15
done

# Get connection info
acmctl cluster get "$CLUSTER"

# Optional: get temp creds for Altinity-side debugging
acmctl cluster temp-creds "$CLUSTER"
```

## Launch a cluster on ClickHouse Keeper (instead of ZooKeeper)

Whether a launch uses Keeper is decided by the **payload**, not by the environment
flag alone: the UI sends `keeperOptions` when the environment has `useKeeper` (or
`useClickHouseAPIv2`) on, and `zookeeper: "launch"` otherwise. Over the API you must
send the right one yourself. A launch with neither is rejected
(`Either CH Keeper or Zookeeper options must be specified.`).

```bash
ENV=426

# 1. Is Keeper enabled for the environment? (informational; the API doesn't enforce it)
acmctl env get "$ENV" | jq '{useKeeper, useClickHouseAPIv2}'

# 2. Pick a Keeper node type (scope "zookeeper" is shared by ZooKeeper and Keeper)
acmctl raw GET "/environment/$ENV/nodetypes?scope=zookeeper" | jq -r '.[].code'

# 3. Launch. keeperOptions creates a dedicated Keeper together with the cluster.
#    name: keep it short; ha:true => 3 Keeper nodes (use for replicas > 1).
cat <<'JSON' | acmctl cluster launch "$ENV" | jq '{id, name, status, keeperName, id_zookeeper}'
{
  "name": "mycluster", "type": "kubernetes", "role": "dev",
  "shards": 1, "replicas": 2, "nodes": 2,
  "nodeType": "m8g.2xlarge", "version": "26.3.33.10001.altinitystable", "memory": 30720,
  "size": 750, "disks": 1, "storageClass": "gp3-encrypted", "throughput": 125, "iops": 3000,
  "lbType": "ingress", "secure": true, "zoneAwareness": true,
  "azlist": ["us-west-2a", "us-west-2b"],
  "adminUser": "admin", "adminPass": "CHANGE-ME", "mysqlProtocol": false, "uptime": "always",
  "replicateSchema": false, "sourceCluster": null, "backupSource": null,
  "keeperOptions": {"name": "mycl-abcd", "instanceType": "t4g.large", "ha": true}
}
JSON
```

Verify it took: the response must have `keeperName` set and `id_zookeeper` null.
On Kubernetes you should also see a `ClickHouseKeeperInstallation` and the CHI's
`spec.configuration.zookeeper.nodes` pointing at `keeperclient-<name>-{0,1,2}:2181`,
with no new `zookeeper-c*` pods.

To reuse an existing Keeper instead, send `"keeperName": "<existing>"` and omit
`keeperOptions`. (The UI does this when replicating a cluster: it copies the source's
`keeperName` or `id_zookeeper` into the launch payload; the API itself does not infer it
from `sourceCluster`.) List Keepers with `GET /environment/{env}/keepers`.

Notes:
- There is no in-place ZooKeeper -> Keeper switch for a running cluster; relaunch or
  restore into a new cluster.
- Don't reuse the same `keeperOptions.name` across clusters; the UI appends a random
  suffix (`<first 10 chars of cluster name>-<4 random letters>`).
- Delete a Keeper only after the clusters using it are gone; the API refuses otherwise.

## Diagnose a slow / failing query

```bash
# 1. Cluster health
acmctl raw GET /cluster/$ID/status

# 2. Find recent errors per node
acmctl raw GET /cluster/$ID/errors

# 3. List slow running queries
acmctl raw GET /cluster/$ID/workload-queries \
  | jq '.[] | select(.elapsed > 10) | {query_id, query, elapsed, node}'

# 4. Drill into a specific query
acmctl raw GET /cluster/$ID/workload-query/$QUERY_ID

# 5. Kill if needed
acmctl raw POST /cluster/$ID/query-kill \
  -F queryIds="$QUERY_ID" \
  -F node="$NODE"
```

## Refresh Altinity support access

```bash
acmctl raw POST /cluster/$ID/support/refresh
resp=$(acmctl cluster temp-creds "$ID")

# Response shape varies — see SKILL.md "Conventions". Handle both:
user=$(echo "$resp" | jq -r 'if type == "object" then .login // empty else empty end')
pass=$(echo "$resp" | jq -r 'if type == "object" then .password else . end')
[ -z "$user" ] && user="${EXPERT_CH_USER:-}"   # fall back to session user
```

## Restore a cluster from S3 backup

```bash
acmctl raw POST /cluster/$ID/restore \
  -F type=s3 \
  -F bucket=my-backup-bucket \
  -F path=/backups/2026-04-29 \
  -F accessKey="$AWS_ACCESS_KEY" \
  -F secretKey="$AWS_SECRET_KEY" \
  -F region=us-east-1

# If it fails partway, retry without redownloading already-fetched parts
acmctl raw POST /cluster/$ID/restore-retry -F skipDownload=1
```

## Apply a large XML setting (Kafka config, custom XML)

```bash
# @file syntax loads from disk
acmctl raw POST /cluster/$ID/kafka-configuration \
  -F filename=kafka.xml \
  -F xml=@./kafka-config.xml
```

## Bulk update cluster settings via field map

```bash
cat <<'JSON' | acmctl cluster update "$ID"
{ "alertsEmail": "ops@example.com", "ipWhitelist": "10.0.0.0/8", "uptime": "24x7" }
JSON
```
