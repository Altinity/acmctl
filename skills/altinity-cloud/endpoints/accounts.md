# Accounts

DELETE /account-role/{id} — Removes ACM user account role
DELETE /account/{id} — Removes ACM user account
GET /access-rights — Gets all available access rights settings
GET /account — Displays complete user account information
GET /account-roles — Lists all ACM user account roles
GET /account/anywhere-token — Generates cloud.anywhere API key
GET /account/token — Generates a random API key
GET /account/{user}/log — Displays user account action log [page, limit, filter, order]
GET /accounts — Lists all ACM user accounts [blocked, organization, page, limit, filter, order]
POST /account — Modifies the current user account [name, password, origins, id_environment, tokens, darkTheme]
POST /account-role/{id} — Modifies ACM user account role [name, rights]
POST /account-roles — Creates new ACM user account role [name, rights]
POST /account/{id} — Modifies ACM user account [email, name, password, id_role, id_organization, origins, blocked, environments, clusters, tokens]
POST /account/{id}/access — Changes environment access for ACM user account [environments]
POST /accounts — Creates new ACM user account [email, name, password, id_role, id_organization, origins, blocked, environments, clusters, tokens]
