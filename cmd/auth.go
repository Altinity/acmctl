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
	Long: `Login with email and password (interactive). Token is saved to the config file.
For a long-lived API key, use --token instead:
  acmctl login --token <api-key>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if tokenFlag != "" {
			cfg.Token = tokenFlag
			if err := config.Save(cfgFile, cfg); err != nil {
				return fmt.Errorf("save config: %w", err)
			}
			fmt.Println("Token saved.")
			return nil
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Email: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)

		fmt.Print("Password: ")
		passBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return fmt.Errorf("read password: %w", err)
		}
		fmt.Println()
		password := string(passBytes)

		user, err := apiClient.Login(email, password)
		if err != nil {
			return fmt.Errorf("login failed: %w", err)
		}
		if user.Token == "" {
			return fmt.Errorf("login succeeded but no token returned (2FA may be required — POST /login/verify with the code)")
		}

		cfg.Token = user.Token
		if err := config.Save(cfgFile, cfg); err != nil {
			return fmt.Errorf("save config: %w", err)
		}
		fmt.Printf("Logged in as %s (%s). Token saved to %s\n", user.Name, user.Email, cfgFile)
		return nil
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Clear the stored authentication token",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg.Token = ""
		if err := config.Save(cfgFile, cfg); err != nil {
			return fmt.Errorf("save config: %w", err)
		}
		fmt.Println("Logged out.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
}
