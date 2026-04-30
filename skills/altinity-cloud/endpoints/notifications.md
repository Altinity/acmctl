# Notifications

DELETE /notification/{id} — Removes admin notification
GET /n — Lists all user notifications [all, page, limit, filter, order]
GET /notification/{id} — Returns given notification object
GET /notifications — Lists all admin notifications [sticky, page, limit, filter, order]
POST /notification/{id} — Modifies a given admin notification [message, level, recipients, expiry, sticky, send, channelEmail, channelPopup]
POST /notifications — Creates new admin notification [message, level, recipients, expiry, sticky, send, channelEmail, channelPopup]
PUT /notification/{id}/ack — Acknowledges a given notification
