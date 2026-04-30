package api

import "fmt"

func (c *Client) GetBilling(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/billing", params, &result)
	return result, err
}

func (c *Client) GetBillingAddress(organization string) (interface{}, error) {
	var result interface{}
	q := map[string]string{}
	if organization != "" {
		q["organization"] = organization
	}
	err := c.Do("GET", "/billing/address", q, &result)
	return result, err
}

func (c *Client) UpdateBillingAddress(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/billing/address", params, &result)
	return result, err
}

func (c *Client) DeleteCreditCard(organization string) error {
	q := map[string]string{}
	if organization != "" {
		q["organization"] = organization
	}
	return c.Do("DELETE", "/billing/credit-card", q, nil)
}

func (c *Client) GetCustomer(organization string, balance string) (interface{}, error) {
	var result interface{}
	q := map[string]string{}
	if balance != "" {
		q["balance"] = balance
	}
	err := c.Do("GET", fmt.Sprintf("/billing/customer/%s", organization), q, &result)
	return result, err
}

func (c *Client) UpdateCustomer(organization string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("PATCH", fmt.Sprintf("/billing/customer/%s", organization), params, &result)
	return result, err
}

func (c *Client) AddCustomerBalance(organization string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/billing/customer/%s/balance", organization), params, &result)
	return result, err
}

func (c *Client) DeleteCustomerBalance(id string) error {
	return c.Do("DELETE", fmt.Sprintf("/billing/customer/balance/%s", id), nil, nil)
}

func (c *Client) ListCustomerDiscounts(organization string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/billing/customer/%s/discounts", organization), nil, &result)
	return result, err
}

func (c *Client) CreateCustomerDiscount(organization string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", fmt.Sprintf("/billing/customer/%s/discounts", organization), params, &result)
	return result, err
}

func (c *Client) UpdateDiscount(id string, params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("PATCH", fmt.Sprintf("/billing/discount/%s", id), params, &result)
	return result, err
}

func (c *Client) DeleteDiscount(id string) error {
	return c.Do("DELETE", fmt.Sprintf("/billing/discount/%s", id), nil, nil)
}

func (c *Client) ListCustomers(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/billing/customers", params, &result)
	return result, err
}

func (c *Client) ListInvoiceTerms() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/billing/invoice-terms", nil, &result)
	return result, err
}

func (c *Client) DownloadInvoice(id string) (interface{}, error) {
	var result interface{}
	err := c.Do("GET", fmt.Sprintf("/billing/invoice/%s", id), nil, &result)
	return result, err
}

func (c *Client) ListInvoices(organization string) (interface{}, error) {
	var result interface{}
	q := map[string]string{}
	if organization != "" {
		q["organization"] = organization
	}
	err := c.Do("GET", "/billing/invoices", q, &result)
	return result, err
}

func (c *Client) CreateInvoice(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/billing/invoices", params, &result)
	return result, err
}

func (c *Client) AddPaymentMethod(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/billing/payment-methods", params, &result)
	return result, err
}

func (c *Client) DeletePaymentMethod(organization string) error {
	q := map[string]string{}
	if organization != "" {
		q["organization"] = organization
	}
	return c.Do("DELETE", "/billing/payment-methods", q, nil)
}

func (c *Client) GetBillingProcess() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/billing/process", nil, &result)
	return result, err
}

func (c *Client) GetBillingReport(period string) (interface{}, error) {
	var result interface{}
	q := map[string]string{}
	if period != "" {
		q["period"] = period
	}
	err := c.Do("GET", "/billing/report", q, &result)
	return result, err
}

func (c *Client) UpdateSubscription(params map[string]string) (interface{}, error) {
	var result interface{}
	err := c.Do("POST", "/billing/subscription", params, &result)
	return result, err
}

func (c *Client) ListTaxes() (interface{}, error) {
	var result interface{}
	err := c.Do("GET", "/billing/taxes", nil, &result)
	return result, err
}
