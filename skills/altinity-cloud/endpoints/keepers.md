# Keepers

DELETE /environment/{environment}/keeper/{name} — Deletes specified CH Keeper within given env (refused while a cluster uses it)
GET /environment/{environment}/keeper/{name}/status — Checks the status of given CH keeper
GET /environment/{environment}/keepers — Lists available CH keepers within given env (each row has `clusters`)
POST /environment/{environment}/keeper/{name} — Updates a given CH Keeper within given env [zones, instanceType, image, ha, settings]
POST /environment/{environment}/keepers — Launches a CH Keeper within given env [name, zones, instanceType, image, ha, settings]

Notes:
- `instanceType` must be a node type with `scope=zookeeper` in the environment
  (`GET /environment/{env}/nodetypes?scope=zookeeper`).
- `image` defaults to the ACM-configured Keeper image; `ha: true` runs 3 nodes, otherwise 1.
- A cluster normally gets its own Keeper at launch via `keeperOptions`; see
  `workflows.md` ("Launch a cluster on ClickHouse Keeper"). Use these endpoints to inspect or
  manage Keepers afterwards.
- ACM has no in-place ZooKeeper -> Keeper conversion for a running cluster.
