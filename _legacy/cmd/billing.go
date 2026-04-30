package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var billingCmd = &cobra.Command{
	Use:   "billing",
	Short: "Manage billing, customers, invoices, and payment methods",
}

var billingGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get billing info for an environment/organization",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"period": "period", "environment": "environment",
			"organization": "organization", "usage": "usage",
		})
		result, err := apiClient.GetBilling(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingAddressGetCmd = &cobra.Command{
	Use:   "address-get",
	Short: "Get billing address",
	RunE: func(cmd *cobra.Command, args []string) error {
		org, _ := cmd.Flags().GetString("organization")
		result, err := apiClient.GetBillingAddress(org)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingAddressUpdateCmd = &cobra.Command{
	Use:   "address-update",
	Short: "Update billing address",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"email": "email", "email-cc": "emailCC", "country": "country",
			"city": "city", "state": "state", "line1": "line1", "line2": "line2",
			"postal-code": "postal_code", "organization": "organization",
			"company-name": "companyName",
		})
		result, err := apiClient.UpdateBillingAddress(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingDeleteCardCmd = &cobra.Command{
	Use:   "delete-card",
	Short: "Remove all attached credit cards",
	RunE: func(cmd *cobra.Command, args []string) error {
		org, _ := cmd.Flags().GetString("organization")
		if err := apiClient.DeleteCreditCard(org); err != nil {
			return err
		}
		fmt.Println("Credit cards removed.")
		return nil
	},
}

var billingCustomerGetCmd = &cobra.Command{
	Use:   "customer-get <organization>",
	Short: "Get customer record",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		balance, _ := cmd.Flags().GetString("balance")
		result, err := apiClient.GetCustomer(args[0], balance)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingCustomerUpdateCmd = &cobra.Command{
	Use:   "customer-update <organization>",
	Short: "Update customer (incl. organization properties)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"trial-expiry": "trialExpiry", "auto-charge": "autoCharge",
			"support-plan": "supportPlan", "deposit": "deposit",
			"invoice-info": "invoiceInfo", "tax-info": "taxInfo",
			"consolidate-billing": "consolidateBilling",
			"auto-charge-no-limit": "autoChargeNoLimit",
		})
		result, err := apiClient.UpdateCustomer(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingBalanceAddCmd = &cobra.Command{
	Use:   "balance-add <organization>",
	Short: "Add a customer balance transaction",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"type": "type", "description": "description", "amount": "amount",
		})
		result, err := apiClient.AddCustomerBalance(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingBalanceDeleteCmd = &cobra.Command{
	Use:   "balance-delete <transaction-id>",
	Short: "Remove a transaction from balance history",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteCustomerBalance(args[0]); err != nil {
			return err
		}
		fmt.Printf("Transaction %s removed.\n", args[0])
		return nil
	},
}

var billingDiscountListCmd = &cobra.Command{
	Use:   "discounts <organization>",
	Short: "List discounts for a customer",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ListCustomerDiscounts(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingDiscountCreateCmd = &cobra.Command{
	Use:   "discount-create <organization>",
	Short: "Create a discount for a customer",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"description": "description", "type": "type", "value": "value",
			"min-amount": "minAmount", "max-amount": "maxAmount",
			"active": "active", "preset": "preset", "products": "products",
		})
		result, err := apiClient.CreateCustomerDiscount(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingDiscountUpdateCmd = &cobra.Command{
	Use:   "discount-update <discount-id>",
	Short: "Update a discount",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"description": "description", "type": "type", "value": "value",
			"min-amount": "minAmount", "max-amount": "maxAmount",
			"active": "active", "products": "products",
		})
		result, err := apiClient.UpdateDiscount(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingDiscountDeleteCmd = &cobra.Command{
	Use:   "discount-delete <discount-id>",
	Short: "Remove a discount",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteDiscount(args[0]); err != nil {
			return err
		}
		fmt.Printf("Discount %s removed.\n", args[0])
		return nil
	},
}

var billingCustomersCmd = &cobra.Command{
	Use:   "customers",
	Short: "List all customers",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"blocked": "blocked", "balance": "balance",
		})
		result, err := apiClient.ListCustomers(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingInvoiceTermsCmd = &cobra.Command{
	Use:   "invoice-terms",
	Short: "List available sales terms",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ListInvoiceTerms()
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingInvoiceGetCmd = &cobra.Command{
	Use:   "invoice <id>",
	Short: "Download an invoice (returns PDF data or location)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.DownloadInvoice(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingInvoicesListCmd = &cobra.Command{
	Use:   "invoices",
	Short: "List invoices for an organization",
	RunE: func(cmd *cobra.Command, args []string) error {
		org, _ := cmd.Flags().GetString("organization")
		result, err := apiClient.ListInvoices(org)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingInvoicesCreateCmd = &cobra.Command{
	Use:   "invoices-create",
	Short: "Prepare an invoice for a period (Admin)",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"period": "period", "organization": "organization",
		})
		result, err := apiClient.CreateInvoice(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingPaymentMethodAddCmd = &cobra.Command{
	Use:   "payment-method-add",
	Short: "Add a billing payment method (Stripe payload)",
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectFieldFlags(cmd)
		if err != nil {
			return err
		}
		for k, v := range flagsToParams(cmd, map[string]string{
			"type": "type", "id": "id", "card": "card", "organization": "organization",
		}) {
			params[k] = v
		}
		result, err := apiClient.AddPaymentMethod(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingPaymentMethodDeleteCmd = &cobra.Command{
	Use:   "payment-method-delete",
	Short: "Remove the assigned billing payment method",
	RunE: func(cmd *cobra.Command, args []string) error {
		org, _ := cmd.Flags().GetString("organization")
		if err := apiClient.DeletePaymentMethod(org); err != nil {
			return err
		}
		fmt.Println("Payment method removed.")
		return nil
	},
}

var billingProcessCmd = &cobra.Command{
	Use:   "process",
	Short: "Get post-flight report of latest process-invoices task (Admin)",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetBillingProcess()
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingReportCmd = &cobra.Command{
	Use:   "report",
	Short: "Prepare invoices report for all organizations (Admin)",
	RunE: func(cmd *cobra.Command, args []string) error {
		period, _ := cmd.Flags().GetString("period")
		result, err := apiClient.GetBillingReport(period)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingSubscriptionCmd = &cobra.Command{
	Use:   "subscription",
	Short: "Add or update Altinity.Cloud subscription",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{"support-plan": "supportPlan"})
		result, err := apiClient.UpdateSubscription(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var billingTaxesCmd = &cobra.Command{
	Use:   "taxes",
	Short: "List available taxes (QuickBooks)",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ListTaxes()
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

func init() {
	for _, c := range []*cobra.Command{billingGetCmd} {
		c.Flags().String("period", "", "billing period")
		c.Flags().String("environment", "", "environment ID")
		c.Flags().String("organization", "", "organization ID")
		c.Flags().String("usage", "", "include usage")
	}

	billingAddressGetCmd.Flags().String("organization", "", "organization ID")
	for _, name := range []string{"email", "email-cc", "country", "city", "state",
		"line1", "line2", "postal-code", "organization", "company-name"} {
		billingAddressUpdateCmd.Flags().String(name, "", "")
	}

	billingDeleteCardCmd.Flags().String("organization", "", "organization ID")

	billingCustomerGetCmd.Flags().String("balance", "", "include balance")
	for _, name := range []string{"trial-expiry", "auto-charge", "support-plan", "deposit",
		"invoice-info", "tax-info", "consolidate-billing", "auto-charge-no-limit"} {
		billingCustomerUpdateCmd.Flags().String(name, "", "")
	}

	for _, name := range []string{"type", "description", "amount"} {
		billingBalanceAddCmd.Flags().String(name, "", "")
	}

	for _, c := range []*cobra.Command{billingDiscountCreateCmd, billingDiscountUpdateCmd} {
		for _, name := range []string{"description", "type", "value", "min-amount", "max-amount",
			"active", "preset", "products"} {
			c.Flags().String(name, "", "")
		}
	}

	billingCustomersCmd.Flags().String("blocked", "", "filter by blocked")
	billingCustomersCmd.Flags().String("balance", "", "include balance")

	billingInvoicesListCmd.Flags().String("organization", "", "organization ID")
	billingInvoicesCreateCmd.Flags().String("period", "", "period")
	billingInvoicesCreateCmd.Flags().String("organization", "", "organization ID")

	for _, name := range []string{"type", "id", "card", "organization"} {
		billingPaymentMethodAddCmd.Flags().String(name, "", "")
	}
	billingPaymentMethodAddCmd.Flags().StringSliceP("field", "F", nil, "key=value (repeatable)")
	billingPaymentMethodDeleteCmd.Flags().String("organization", "", "organization ID")

	billingReportCmd.Flags().String("period", "", "period")
	billingSubscriptionCmd.Flags().String("support-plan", "", "support plan")

	for _, c := range []*cobra.Command{
		billingGetCmd, billingAddressGetCmd, billingAddressUpdateCmd, billingDeleteCardCmd,
		billingCustomerGetCmd, billingCustomerUpdateCmd, billingBalanceAddCmd, billingBalanceDeleteCmd,
		billingDiscountListCmd, billingDiscountCreateCmd, billingDiscountUpdateCmd, billingDiscountDeleteCmd,
		billingCustomersCmd, billingInvoiceTermsCmd, billingInvoiceGetCmd, billingInvoicesListCmd,
		billingInvoicesCreateCmd, billingPaymentMethodAddCmd, billingPaymentMethodDeleteCmd,
		billingProcessCmd, billingReportCmd, billingSubscriptionCmd, billingTaxesCmd,
	} {
		billingCmd.AddCommand(c)
	}
	rootCmd.AddCommand(billingCmd)
}
