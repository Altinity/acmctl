# Console

DELETE /console/task/{id} — Removes a new scheduled task from the database
GET /console/info — Returns some information about background operations
GET /console/logs — Displays internal audit log for all environments [page, limit, filter, order]
GET /console/logs/filtered — Gets logs for a given set of labels [labels]
GET /console/logs/{environment} — Gets tail logs for a given cluster inside environment [count, cluster, filter, download, section]
GET /console/settings — Gets the list of system operations settings
GET /console/tasks — Gets the list of system background tasks [page, limit, filter, order]
PATCH /console/settings — Updates system operations settings [maxProcesses, maxBackupProcesses, maxClusterProcesses, autoChargePercentCondition, currentCronHost, cronProcessTimeLimit, backupProcessTimeLimit, clusterProcessTimeLimit]
POST /console/task/{id} — Update a running system task [interrupt]
POST /console/tasks — Schedules a new background task [action, data]
