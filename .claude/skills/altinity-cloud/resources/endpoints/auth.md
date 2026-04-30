# Auth

GET /auth0/connections — Returns available Auth0 connections for authentication
GET /login/reset — Checks whether password reset token is active [code]
GET /logout — Provides an URL to be used a logout target
GET /probe — System probe
GET /signup/confirm — Checks whether signup token is active [code]
GET /singleauth — Provides with a target URL for Auth0 [type]
POST /aws-marketplace-gateway — Handles landing for AWS Marketplace [x-amzn-marketplace-token, confirm]
POST /aws-marketplace-sub — Handles AWS Marketplace SNS subscriptions [Type, MessageId, Token, TopicArn, Message, SubscribeURL, Timestamp, SignatureVersion, Signature, SigningCertURL]
POST /gcp-marketplace-gateway — Handles landing for Google Marketplace [x-gcp-marketplace-token]
POST /login — Provides a token for a pair of users login and password [login, password]
POST /login/recover — Sends out an email with password reset link [login]
POST /login/reset — Checks whether password reset token is active [code, password]
POST /login/verify — Login with 2FA verification code [code, user]
POST /signup — Creates a trial account in Altinity.Cloud [email, name, company, envName, deployment, provider, region, location, captcha]
POST /signup-email — Creates a trial account in Altinity.Cloud using only Email [name, email, captcha]
POST /signup/confirm — Finishes up signup procedure [code, password]
POST /singleauth — Authenticate user with Auth0 oAuth token [code, state, token]
