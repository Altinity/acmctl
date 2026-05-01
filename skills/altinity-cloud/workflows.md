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
