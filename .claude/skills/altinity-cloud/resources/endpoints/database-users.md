# DatabaseUsers

DELETE /cluster/{cluster}/user/{id} — Remove user by name from a given cluster
DELETE /user/{id} — Remove cluster user from a given cluster
GET /cluster/{cluster}/users — Lists all cluster users from a given cluster
POST /cluster/{cluster}/user/{id} — Modify cluster user of a given cluster [login, password, databases, id_profile, id_quota, accessManagement]
POST /cluster/{cluster}/users — Add cluster user to a given cluster [login, password, databases, id_profile, id_quota, accessManagement]
POST /user/{id} — Modify cluster user of a given cluster [login, password, databases, id_profile, id_quota, accessManagement]
