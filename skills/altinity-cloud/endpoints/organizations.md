# Organizations

DELETE /organization/{id} — Removes organization descriptor
GET /organization/{id} — Returns an organization entry
GET /organizations — Lists all organization descriptors [blocked, billing]
POST /organization/{id} — Modifies organization descriptor [name, emailDomain, inheritEnvironment, opened, id_defaultUserRole, limited, blocked, blockedPassword, blockedAPI, allowAdminPassword, enable2FA, environments, trialExpiry, autoCharge, userSyncSettings, autoChargeNoLimit]
POST /organization/{id}/logins — Sets up organization login settings [name, blockedPassword, blockedAPI, allowAdminPassword, enable2FA, opened, id_defaultUserRole, userSyncSettings]
POST /organizations — Creates new organization descriptor [name, emailDomain, inheritEnvironment, opened, id_defaultUserRole, limited, blocked, blockedPassword, blockedAPI, allowAdminPassword, enable2FA, environments, trialExpiry, autoCharge, userSyncSettings]
