# NodeTypes

DELETE /nodetype/{id} — Removes given Node Type from the environment
GET /environment/{environment}/nodetypes — Lists available node types for given Environment [scope, withUsed]
POST /environment/{environment}/nodetypes — Adds a Zookeeper Cluster into the specified environment [name, scope, code, memory, cpu, storageClass, capacity, extraSpec, nodeSelector, tolerations, isSpot]
POST /nodetype/{id} — Modifies given Node Type [name, scope, code, memory, cpu, storageClass, capacity, extraSpec, nodeSelector, tolerations, isSpot]
