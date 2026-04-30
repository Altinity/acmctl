# Billing

DELETE /billing/credit-card — Removes all attached credit cards [organization]
DELETE /billing/customer/balance/{id} — Removes a transaction from the customer's balance history
DELETE /billing/discount/{id} — Removes the discount record from a given customer
DELETE /billing/payment-methods — Removes assigned billing payment method [organization]
GET /billing — Returns billing info for a specified environment [period, environment, organization, usage]
GET /billing/address — Returns billing address for an organization [organization]
GET /billing/customer/{organization} — Returns a customer [balance]
GET /billing/customer/{organization}/discounts — Returns the list of available discounts for the customer
GET /billing/customers — Returns list of all customers [blocked, balance]
GET /billing/invoice-terms — Returns the list of available sales terms
GET /billing/invoice/{id} — Downloads given invoice as PDF
GET /billing/invoices — Returns the list of available invoices [organization]
GET /billing/process — Returns the post-flight report of the latest process invoices task (Admin ONLY)
GET /billing/report — Prepares invoices report for all organizations (Admin ONLY) [period]
GET /billing/taxes — Returns the list of available taxes from QB
PATCH /billing/customer/{organization} — Edits a Customer (incl. Organization properties) [trialExpiry, autoCharge, supportPlan, deposit, invoiceInfo, taxInfo, consolidateBilling, autoChargeNoLimit]
PATCH /billing/discount/{id} — Updates a discount record for a given customer [description, type, value, minAmount, maxAmount, active, products]
POST /billing/address — Updates billing address info [email, emailCC, country, city, state, line1, line2, postal_code, organization, companyName]
POST /billing/customer/{organization}/balance — Adds a customer balance transaction [type, description, amount]
POST /billing/customer/{organization}/discounts — Creates a discount record for a given customer [description, type, value, minAmount, maxAmount, active, preset, products]
POST /billing/invoices — Prepares an invoice for a given period (Admin ONLY) [period, organization]
POST /billing/payment-methods — Adds billing payment method. Accepts Stripe payment method payload [type, id, card, organization]
POST /billing/subscription — Adds or Updates Altinity.Cloud subscription [supportPlan]
