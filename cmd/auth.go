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
	Short: "Authenticate with ACM (email/password) and store the token in the active profile",
	Long: `Login with email and password (interactive). Token is saved into
the active profile in the config file (default ~/.acmctl.yaml).

For long-lived API keys, use --token <api-key> instead.

For OAuth (browser-based) login, use ` + "`acmctl oauth`" + ` instead — most
ACM tenants are SSO-only and don't accept email/password directly.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		profileName, err := requireProfile()
		if err != nil {
			return err
		}
		profile := cfg.Profiles[profileName]

		if tokenFlag != "" {
			profile.Token = tokenFlag
			cfg.SetProfile(profileName, profile)
			if err := config.Save(cfgFile, cfg); err != nil {
				return fmt.Errorf("save config: %w", err)
			}
			fmt.Printf("Token saved to profile %q.\n", profileName)
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
			return fmt.Errorf("login succeeded but no token returned (server may require 2FA — not supported by this CLI; use 'acmctl oauth' instead)")
		}

		profile.Token = user.Token
		cfg.SetProfile(profileName, profile)
		if err := config.Save(cfgFile, cfg); err != nil {
			return fmt.Errorf("save config: %w", err)
		}
		fmt.Printf("Logged in as %s (%s). Token saved to profile %q in %s\n", user.Name, user.Email, profileName, cfgFile)
		return nil
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Clear the stored token (active profile by default; --all clears every profile)",
	RunE: func(cmd *cobra.Command, args []string) error {
		all, _ := cmd.Flags().GetBool("all")

		if all {
			cleared := 0
			for n, p := range cfg.Profiles {
				if p.Token == "" {
					continue
				}
				p.Token = ""
				cfg.Profiles[n] = p
				cleared++
			}
			if cleared == 0 {
				fmt.Println("No profiles had a token set.")
				return nil
			}
			if err := config.Save(cfgFile, cfg); err != nil {
				return fmt.Errorf("save config: %w", err)
			}
			fmt.Printf("Cleared tokens from %d profile(s).\n", cleared)
			return nil
		}

		profileName, err := requireProfile()
		if err != nil {
			return err
		}
		profile := cfg.Profiles[profileName]
		if profile.Token == "" {
			fmt.Printf("Profile %q has no token (already logged out).\n", profileName)
			return nil
		}
		profile.Token = ""
		cfg.SetProfile(profileName, profile)
		if err := config.Save(cfgFile, cfg); err != nil {
			return fmt.Errorf("save config: %w", err)
		}
		fmt.Printf("Logged out of profile %q.\n", profileName)
		return nil
	},
}

func init() {
	logoutCmd.Flags().Bool("all", false, "clear tokens from every profile, not just the active one")
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
}
