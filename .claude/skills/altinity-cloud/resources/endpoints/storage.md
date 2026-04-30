# Storage

DELETE /cluster/{cluster}/object-storages/{name} — Removes object storage from cluster
DELETE /storage/cluster-volume/{clusterVolume} — Removes cluster volume
GET /storage/cluster/{cluster}/volumes — Lists all cluster volumes
PATCH /storage/cluster-volume/{clusterVolume} — Modifies cluster volume [size, throughput, type, iops]
POST /cluster/{cluster}/object-storages — Adds object storage to cluster [type, bucket, region, accessKey, secretKey, createBucket]
POST /storage/cluster-volume/{clusterVolume}/cordon — Sets cordon value for cluster volume [cordon]
POST /storage/cluster-volume/{clusterVolume}/free — Frees up cluster volume
POST /storage/cluster-volume/{clusterVolume}/validate-modification-pvc — Validates cluster volume modification based on PVC status
POST /storage/cluster/{cluster}/interrupt-free — Interrupts running free volume queries
POST /storage/cluster/{cluster}/volumes — Modifies cluster volume [type, size, throughput, iops]
