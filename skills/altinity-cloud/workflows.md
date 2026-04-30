# ACM Multi-Step Workflows

Assumes `ACM_TOKEN`, `ACM_URL`, and the `acm()` helper from SKILL.md are set up.

## Launch a cluster and wait until ready

```bash
ENV=$(acm GET /environments | jq '.[0].id')

CLUSTER=$(acm POST /environment/$ENV/clusters/launch \
  --data-urlencode "name=test-cluster" \
  --data-urlencode "nodeType=s1" \
  --data-urlencode "version=24.3" \
  | jq -r '.id')

# Poll status until uptime > 0 (means it's accepting connections)
until [ "$(acm GET /cluster/$CLUSTER/status | jq '.uptime // 0')" -gt 0 ]; do
  sleep 15
done

# Get connection info
acm GET /cluster/$CLUSTER

# Optional: get temp creds for Altinity-side debugging
acm GET /cluster/$CLUSTER/support/credentials
```

## Diagnose a slow / failing query

```bash
# 1. Cluster health
acm GET /cluster/$ID/status

# 2. Find recent errors per node
acm GET /cluster/$ID/errors

# 3. List slow running queries
acm GET /cluster/$ID/workload-queries \
  | jq '.[] | select(.elapsed > 10) | {query_id, query, elapsed, node}'

# 4. Drill into a specific query
acm GET /cluster/$ID/workload-query/$QUERY_ID

# 5. Kill if needed
acm POST /cluster/$ID/query-kill \
  --data-urlencode "queryIds=$QUERY_ID" \
  --data-urlencode "node=$NODE"
```

## Refresh Altinity support access

```bash
acm POST /cluster/$ID/support/refresh
resp=$(acm GET /cluster/$ID/support/credentials)

# Response shape varies — see SKILL.md "Conventions". Handle both:
user=$(echo "$resp" | jq -r 'if type == "object" then .login // empty else empty end')
pass=$(echo "$resp" | jq -r 'if type == "object" then .password else . end')
[ -z "$user" ] && user="${EXPERT_CH_USER:-}"   # fall back to session user
```

## Restore a cluster from S3 backup

```bash
acm POST /cluster/$ID/restore \
  --data-urlencode "type=s3" \
  --data-urlencode "bucket=my-backup-bucket" \
  --data-urlencode "path=/backups/2026-04-29" \
  --data-urlencode "accessKey=$AWS_ACCESS_KEY" \
  --data-urlencode "secretKey=$AWS_SECRET_KEY" \
  --data-urlencode "region=us-east-1"

# If it fails partway, retry without redownloading already-fetched parts
acm POST /cluster/$ID/restore-retry --data-urlencode "skipDownload=1"
```

## Apply a large XML setting (Kafka config, custom XML)

```bash
# @file syntax loads from disk
acm POST /cluster/$ID/kafka-configuration \
  --data-urlencode "filename=kafka.xml" \
  --data-urlencode "xml@./kafka-config.xml"
```

## Bulk update cluster settings via field map

```bash
acm POST /cluster/$ID \
  --data-urlencode "alertsEmail=ops@example.com" \
  --data-urlencode "ipWhitelist=10.0.0.0/8" \
  --data-urlencode "uptime=24x7"
```
