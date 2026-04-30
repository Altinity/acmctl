# ClusterSettings

DELETE /cluster-env-setting/{id} — Removes an environment setting from a given cluster
DELETE /cluster-setting/{id} — Removes a cluster setting from a given cluster
GET /cluster/{cluster}/env-settings — Lists all cluster environment settings
GET /cluster/{cluster}/settings — Lists all cluster settings
GET /cluster/{cluster}/system-settings — Lists all cluster system settings
POST /cluster-env-setting/{id} — Modifies an environment setting for a given cluster [name, value, valueFrom]
POST /cluster-setting/{id} — Modifies a cluster setting for a given cluster [name, value, description, valueFrom]
POST /cluster/{cluster}/env-settings — Adds an environemnt setting to a given cluster [name, value, valueFrom]
POST /cluster/{cluster}/settings — Adds a cluster setting to a given cluster [name, value, description, valueFrom]
