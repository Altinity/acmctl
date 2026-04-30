package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/altinity/acmctl/pkg/config"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with ACM and store the token",
	Long: `Login with email and password. The returned token is saved to the config file.
Alternatively, set a long-lived API key directly:
  acmctl login --token <api-key>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If --token flag was provided, just save it directly
		if tokenFlag != "" {
			cfg.Token = tokenFlag
			if err := config.Save(cfgFile, cfg); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}
			fmt.Println("Token saved successfully.")
			return nil
		}

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Email: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)

		fmt.Print("Password: ")
		passBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return fmt.Errorf("failed to read password: %w", err)
		}
		fmt.Println()
		password := string(passBytes)

		user, err := apiClient.Login(email, password)
		if err != nil {
			return fmt.Errorf("login failed: %w", err)
		}

		if user.Token == "" {
			return fmt.Errorf("login succeeded but no token returned (2FA may be required)")
		}

		cfg.Token = user.Token
		if err := config.Save(cfgFile, cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Printf("Logged in as %s (%s). Token saved to %s\n", user.Name, user.Email, cfgFile)
		return nil
	},
}

var loginVerifyCmd = &cobra.Command{
	Use:   "login-verify",
	Short: "Complete 2FA verification with code and user ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		code, _ := cmd.Flags().GetString("code")
		userID, _ := cmd.Flags().GetString("user")
		if code == "" || userID == "" {
			return fmt.Errorf("--code and --user are required")
		}
		user, err := apiClient.LoginVerify(code, userID)
		if err != nil {
			return fmt.Errorf("verify failed: %w", err)
		}
		if user.Token == "" {
			return fmt.Errorf("verification succeeded but no token returned")
		}
		cfg.Token = user.Token
		if err := config.Save(cfgFile, cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}
		fmt.Printf("Verified. Token saved to %s\n", cfgFile)
		return nil
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Clear stored authentication token",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg.Token = ""
		if err := config.Save(cfgFile, cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}
		fmt.Println("Logged out. Token cleared.")
		return nil
	},
}

var probeCmd = &cobra.Command{
	Use:   "probe",
	Short: "Health-check the API",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.Probe()
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var auth0ConnectionsCmd = &cobra.Command{
	Use:   "auth0-connections",
	Short: "List available Auth0 connections",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.Auth0Connections()
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var loginRecoverCmd = &cobra.Command{
	Use:   "login-recover <email>",
	Short: "Send a password reset email",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.LoginRecover(args[0]); err != nil {
			return err
		}
		fmt.Println("Recovery email sent.")
		return nil
	},
}

var loginResetCheckCmd = &cobra.Command{
	Use:   "login-reset-check <code>",
	Short: "Check whether a password reset code is still active",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.CheckResetToken(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var loginResetCmd = &cobra.Command{
	Use:   "login-reset",
	Short: "Reset password using a recovery code",
	RunE: func(cmd *cobra.Command, args []string) error {
		code, _ := cmd.Flags().GetString("code")
		password, _ := cmd.Flags().GetString("password")
		if code == "" || password == "" {
			return fmt.Errorf("--code and --password are required")
		}
		if err := apiClient.ResetPassword(code, password); err != nil {
			return err
		}
		fmt.Println("Password reset.")
		return nil
	},
}

var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "Create a trial account",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"email": "email", "name": "name", "company": "company",
			"env-name": "envName", "deployment": "deployment",
			"provider": "provider", "region": "region", "location": "location",
			"captcha": "captcha",
		})
		result, err := apiClient.Signup(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var signupEmailCmd = &cobra.Command{
	Use:   "signup-email",
	Short: "Create a trial account using only email",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"email": "email", "name": "name", "captcha": "captcha",
		})
		result, err := apiClient.SignupEmail(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var signupConfirmCheckCmd = &cobra.Command{
	Use:   "signup-confirm-check <code>",
	Short: "Check whether a signup confirmation code is still active",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.CheckSignupToken(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var signupConfirmCmd = &cobra.Command{
	Use:   "signup-confirm",
	Short: "Confirm signup with code and chosen password",
	RunE: func(cmd *cobra.Command, args []string) error {
		code, _ := cmd.Flags().GetString("code")
		password, _ := cmd.Flags().GetString("password")
		if code == "" || password == "" {
			return fmt.Errorf("--code and --password are required")
		}
		result, err := apiClient.SignupConfirm(code, password)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var singleAuthURLCmd = &cobra.Command{
	Use:   "singleauth-url",
	Short: "Get the Auth0 single-sign-on URL",
	RunE: func(cmd *cobra.Command, args []string) error {
		typ, _ := cmd.Flags().GetString("type")
		result, err := apiClient.SingleAuthURL(typ)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var singleAuthCmd = &cobra.Command{
	Use:   "singleauth",
	Short: "Authenticate with an Auth0 oAuth code or token",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"code": "code", "state": "state", "token": "token",
		})
		result, err := apiClient.SingleAuth(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var awsMarketplaceGatewayCmd = &cobra.Command{
	Use:   "aws-marketplace-gateway",
	Short: "AWS Marketplace landing handler",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"token": "x-amzn-marketplace-token", "confirm": "confirm",
		})
		result, err := apiClient.AWSMarketplaceGateway(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var awsMarketplaceSubCmd = &cobra.Command{
	Use:   "aws-marketplace-sub",
	Short: "AWS Marketplace SNS subscription handler",
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectFieldFlags(cmd)
		if err != nil {
			return err
		}
		result, err := apiClient.AWSMarketplaceSub(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var gcpMarketplaceGatewayCmd = &cobra.Command{
	Use:   "gcp-marketplace-gateway",
	Short: "Google Marketplace landing handler",
	RunE: func(cmd *cobra.Command, args []string) error {
		token, _ := cmd.Flags().GetString("token")
		params := map[string]string{}
		if token != "" {
			params["x-gcp-marketplace-token"] = token
		}
		result, err := apiClient.GCPMarketplaceGateway(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

func init() {
	loginVerifyCmd.Flags().String("code", "", "2FA code")
	loginVerifyCmd.Flags().String("user", "", "user ID returned by initial login")

	loginResetCmd.Flags().String("code", "", "reset code")
	loginResetCmd.Flags().String("password", "", "new password")

	for _, c := range []*cobra.Command{signupCmd} {
		c.Flags().String("email", "", "email")
		c.Flags().String("name", "", "name")
		c.Flags().String("company", "", "company")
		c.Flags().String("env-name", "", "environment name")
		c.Flags().String("deployment", "", "deployment type")
		c.Flags().String("provider", "", "cloud provider")
		c.Flags().String("region", "", "region")
		c.Flags().String("location", "", "location")
		c.Flags().String("captcha", "", "captcha token")
	}

	signupEmailCmd.Flags().String("email", "", "email")
	signupEmailCmd.Flags().String("name", "", "name")
	signupEmailCmd.Flags().String("captcha", "", "captcha token")

	signupConfirmCmd.Flags().String("code", "", "confirmation code")
	signupConfirmCmd.Flags().String("password", "", "chosen password")

	singleAuthURLCmd.Flags().String("type", "", "auth type")

	singleAuthCmd.Flags().String("code", "", "oAuth code")
	singleAuthCmd.Flags().String("state", "", "oAuth state")
	singleAuthCmd.Flags().String("token", "", "oAuth token")

	awsMarketplaceGatewayCmd.Flags().String("token", "", "x-amzn-marketplace-token")
	awsMarketplaceGatewayCmd.Flags().String("confirm", "", "confirm flag")

	awsMarketplaceSubCmd.Flags().StringSliceP("field", "F", nil, "key=value (repeatable)")

	gcpMarketplaceGatewayCmd.Flags().String("token", "", "x-gcp-marketplace-token")

	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(loginVerifyCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(probeCmd)
	rootCmd.AddCommand(auth0ConnectionsCmd)
	rootCmd.AddCommand(loginRecoverCmd)
	rootCmd.AddCommand(loginResetCheckCmd)
	rootCmd.AddCommand(loginResetCmd)
	rootCmd.AddCommand(signupCmd)
	rootCmd.AddCommand(signupEmailCmd)
	rootCmd.AddCommand(signupConfirmCheckCmd)
	rootCmd.AddCommand(signupConfirmCmd)
	rootCmd.AddCommand(singleAuthURLCmd)
	rootCmd.AddCommand(singleAuthCmd)
	rootCmd.AddCommand(awsMarketplaceGatewayCmd)
	rootCmd.AddCommand(awsMarketplaceSubCmd)
	rootCmd.AddCommand(gcpMarketplaceGatewayCmd)
}
