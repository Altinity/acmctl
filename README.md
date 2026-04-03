# acmctl

CLI tool for [Altinity Cloud Manager API](https://acm.altinity.cloud/docs/).

## Installation

```bash
go build -o acmctl .
cp acmctl ~/bin/
```

## Authentication

Set your API token via environment variable:

```bash
export ACMCTL_TOKEN="your-api-token"
```

Or save it to config:

```bash
acmctl login --token "your-api-token"
```

Or login interactively with email/password:

```bash
acmctl login
```

Token is stored in `~/.acmctl.yaml`.

## Usage

### Output formats

All list/get commands support `--output` (`-o`) flag: `table` (default), `json`, `yaml`.

```bash
acmctl env list
acmctl env list -o json
acmctl env list -o yaml
```

### Environments

```bash
# List all environments
acmctl env list

# Get environment details
acmctl env get 37
acmctl env get 37 -o json

# List clusters in an environment
acmctl env clusters 37

# List node types available in an environment
acmctl env nodetypes 37

# List zookeeper clusters in an environment
acmctl env zookeepers 37

# Delete an environment
acmctl env delete 123
```

### Clusters

```bash
# List all clusters
acmctl cluster list

# Filter clusters (using shell tools)
acmctl cluster list | grep demo

# Get cluster details
acmctl cluster get 337
acmctl cluster get 337 -o json
acmctl cluster get 337 -o yaml

# Get cluster status
acmctl cluster status 337

# List cluster nodes
acmctl cluster nodes 337

# List cluster backups
acmctl cluster backups 337

# Launch a new cluster in an environment
acmctl cluster launch 37 --name my-cluster --node-type s1 --version 24.3

# Restart a cluster
acmctl cluster restart 337

# Stop a cluster
acmctl cluster stop 337

# Resume a stopped cluster
acmctl cluster resume 337

# Upgrade ClickHouse version
acmctl cluster upgrade 337 --version 24.8

# Rescale a cluster
acmctl cluster rescale 337 --node-type m1 --replicas 2

# Trigger a backup
acmctl cluster backup 337

# Restore from backup
acmctl cluster restore 337 --type s3 --bucket my-bucket --path /backups/latest

# Execute a SQL query
acmctl cluster query 337 --query "SELECT count() FROM system.tables"

# Push cluster configuration
acmctl cluster push 337

# Delete a cluster
acmctl cluster delete 337
acmctl cluster delete 337 --terminate
```

### Nodes

```bash
# List nodes in a cluster
acmctl cluster nodes 337

# Get node status
acmctl node status 193

# Get node metrics
acmctl node metrics 193
acmctl node metrics 193 --detailed

# Restart a node
acmctl node restart 193
acmctl node restart 193 --hard
```

### Database Users

```bash
# List database users for a cluster
acmctl dbuser list 337

# Create a database user
acmctl dbuser create 337 --login myuser --password secret123

# Delete a database user
acmctl dbuser delete 337 990
```

### Cluster Settings

```bash
# List cluster settings
acmctl setting list 337

# Create a setting
acmctl setting create 337 --name max_memory_usage --value 10000000000

# Delete a setting
acmctl setting delete 456
```

### Accounts

```bash
# Get current account info
acmctl account get

# List all accounts
acmctl account list

# Create an account
acmctl account create --email user@example.com --name "John Doe" --password secret

# Delete an account
acmctl account delete 42
```

### Debugging

```bash
# Verbose mode — shows HTTP requests and responses
acmctl env list -v
acmctl cluster get 337 -v
```

### Override settings per command

```bash
# Use a different token
acmctl --token other-token env list

# Use a different ACM instance
acmctl --url https://other-acm.example.com/api/ env list

# Use a different config file
acmctl --config /path/to/config.yaml env list
```

### Shell completion

```bash
# Bash
source <(acmctl completion bash)

# Zsh
source <(acmctl completion zsh)

# Fish
acmctl completion fish | source
```

## Configuration

Config file: `~/.acmctl.yaml`

```yaml
url: https://acm.altinity.cloud/api/
token: your-api-token
output: table
```

Environment variables (override config file):

- `ACMCTL_TOKEN` — API token
- `ACMCTL_URL` — API base URL

## API Reference

Full API spec: [reference.json](reference.json) / [reference_auth.json](reference_auth.json)

Swagger UI: https://acm.altinity.cloud/docs/
