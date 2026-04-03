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

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
}
