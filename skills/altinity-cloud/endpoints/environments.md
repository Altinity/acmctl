# Environments

DELETE /environment/{id} — Removes environment along with all related resources like clusters, ZK descriptors, etc. [namespace, clusters, skipKube, skipCloud]
DELETE /environment/{id}/chop-configuration/{name} — Deletes given CHOP configuration from Kubernetes environment
GET /environment/{id} — Gets complete environment specification specified by unique ID or name
GET /environment/{id}/acc-check — Checks the connection between Cloud Connector and Cloud Controller [noWait]
GET /environment/{id}/alerts — Lists all alerts for given environment [resolved]
GET /environment/{id}/buckets — Lists all buckets in the environment [type]
GET /environment/{id}/chop-configurations — Lists all CHOP configurations from Kubernetes environment
GET /environment/{id}/cluster-launch-validity — Checks whether cluster launch is valid
GET /environment/{id}/configuration-templates — Lists all configuration templates from Kubernetes environment
GET /environment/{id}/discover — Discovers Kubernetes environment for existing clusters
GET /environment/{id}/export — Exports environment specification into JSON file [format]
GET /environment/{id}/iceberg — Extract Iceberg Catalog related settings
GET /environment/{id}/invite-details — Display default organization to which user will be assigned
GET /environment/{id}/kube-configmaps — Lists all ConfigMaps from Kubernetes environment
GET /environment/{id}/kube-map — Displays kubernetes resources map
GET /environment/{id}/log — Displays environment action log (Internal & Control Plane) [source, page, limit, filter, order]
GET /environment/{id}/resource — Displays kubernetes resource spec [kind, name, apiVersion]
GET /environment/{id}/resources — Lists all available environment resources [limits]
GET /environment/{id}/usage — Lists environment resource usage [skipClusterId]
GET /environments — Lists all available environments [organization, checkClusters]
POST /backups — Returns the list of entries within given external bucket [type, provider, accessKey, secretKey, arn, endpoint, region, bucket, path, check]
POST /environment/{environment}/backups — Returns the list of entries within given external bucket [type, path, check, schedule]
POST /environment/{id} — Modifies environment details (Deployment type, connection credentials, etc) [name, displayName, type, domain, externalDNS, sslCertificateARN, user, pass, awsKey, awsSecretKey, awsKeyPairName, awsPrivateKey, awsSettings, awsSettingsAuto, kubeProvider, kubeNamespace, kubeNamespaceManage, kubeAPIUrl, kubeToken, kubeAuthOptions, kubeNodeLabel, kubeLBType, kubeStartupTime, options, resourceLimits, autoPush, monitoring, dashboardUrl, backupOptions, notes, datadogSettings, sniProxyForCH, kubeManagedPVs, logsStorage, imageRegistry, managed, remote, metricStorage, logsOptions, maintenanceWindowSchedules, publicLB, privateLB, tags, dynamicTags, kubeZoneKey, loadBalancePROXYProtocolV2, applyToClusters, iceberg, internalEncryption, vpcEndpointsReference]
POST /environment/{id}/acc-connect — Sets up given Environment into a state of awaiting connection [resources]
POST /environment/{id}/approve — Approval for the environment setup request [reason]
POST /environment/{id}/chop-configurations — Adds new CHOP configuration to Kubernetes environment [name, spec]
POST /environment/{id}/config-apply — Applies all Kubernetes ClickHouse operator ConfigMap, Configuration Template, CHOP Configuration changes by restarting pods
POST /environment/{id}/discover — Confirms which Kubernetes clusters should be checked out [clusters]
POST /environment/{id}/get-acc-token — Returns a token for establishing a connection between Cloud Connector and Cloud Controller
POST /environment/{id}/invite — Invite a user to the Environment [email, id_role]
POST /environment/{id}/kube-cho — Handles installation and removal of Clickhouse Operator inside Kubernetes type of Environment [action, version]
POST /environment/{id}/kube-update — Refreshes Kubernetes environment details from acm-env-details ConfigMap
POST /environments/connect — Connects an existing environment to ACM [name, type, host, port, user, pass, sshHost, sshPort, sshUser, sshPass, kubeAPIUrl, kubeToken, kubeAuthOptions, kubeNamespace, kubeNamespaceManage]
POST /environments/import — Imports an environment from a JSON source (previously exported with ACM Export API)
POST /environments/request — Request of provisioning a new environment [name, cloud_provider, aws_region, gcp_region, azure_region, hcloud_region, first]
PUT /environment/{id}/chop-configuration-apply — Apply all CHOP configuration changes
PUT /environment/{id}/chop-configuration/{name} — Patches given CHOP configuration inside Kubernetes environment [spec]
PUT /environment/{id}/configuration-template/{name} — Patches given configuration template inside Kubernetes environment [spec]
PUT /environment/{id}/kube-configmap/{name} — Patches given ConfigMap inside Kubernetes environment [data]
PUT /environment/{id}/reset — Resets given Environment into initial state
