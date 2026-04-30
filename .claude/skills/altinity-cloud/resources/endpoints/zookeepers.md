# Zookeepers

DELETE /zookeeper/{id} — Removes given Zookeeper Cluster from a the environment
GET /environment/{environment}/zookeepers — Lists available zookeeper clusters [showDedicated]
GET /zookeeper/{id}/status — Checks out Zookeeper Cluster status
POST /environment/{environment}/zookeepers — Adds a Zookeeper Cluster into the specified environment [tag, hosts, suffix]
POST /environment/{environment}/zookeepers/launch — Launches a Zookeeper cluster inside given cloud environment [tag, size, nodeType]
POST /zookeeper/{id} — Modifies a Zookeeper Cluster [tag, hosts]
PUT /zookeeper/{id}/push — Publishes Zookeeper configuration to the cloud provider
PUT /zookeeper/{id}/rescale — Rescales a given Zookeeper Cluster [size, nodeType]
