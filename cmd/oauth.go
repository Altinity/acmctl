package cmd

import (
	"context"
	"fmt"

	"github.com/altinity/acmctl/pkg/config"
	"github.com/spf13/cobra"
)

var oauthCmd = &cobra.Command{
	Use:   "oauth",
	Short: "Authenticate via Auth0 (browser-based PKCE flow); save token to active profile",
	Long: `Run the OAuth authorization-code flow against altinity.auth0.com:

  - Opens your default browser to Auth0
  - Listens on a fixed loopback port (49152, fallback 49153/49154) to
    catch the redirect
  - Exchanges the code with Auth0 for an id_token
  - Posts the id_token to ACM /singleauth for an ACM session token
  - Saves the session token into the active profile in your config

Usage with profiles:
  acmctl --profile prod oauth   # log in to prod
  acmctl --profile dev oauth    # log in to dev

The token is short-lived (Auth0 JWT, typically hours). For
long-lived programmatic access, mint an API key in the ACM web UI
and put it in 1Password under op://Employee/ACM_API_KEY.

Auth0 setup required (one-time, by an admin of the altinity.auth0.com
tenant): a Native application with PKCE-only auth and these callback
URLs allowed:

  http://localhost:49152/cb
  http://localhost:49153/cb
  http://localhost:49154/cb`,
	RunE: func(cmd *cobra.Command, args []string) error {
		profileName, err := requireProfile()
		if err != nil {
			return err
		}
		profile := cfg.Profiles[profileName]

		noBrowser, _ := cmd.Flags().GetBool("no-browser")
		ctx := context.Background()
		result, err := apiClient.OAuthLogin(ctx, !noBrowser)
		if err != nil {
			return err
		}
		profile.Token = result.SessionToken
		cfg.SetProfile(profileName, profile)
		if err := config.Save(cfgFile, cfg); err != nil {
			return fmt.Errorf("save config: %w", err)
		}
		who := "Logged in"
		if result.User != nil && result.User.Email != "" {
			who = fmt.Sprintf("Logged in as %s", result.User.Email)
		}
		fmt.Printf("%s. Token saved to profile %q in %s\n", who, profileName, cfgFile)
		return nil
	},
}

func init() {
	oauthCmd.Flags().Bool("no-browser", false,
		"don't try to open the URL automatically; just print it")
	rootCmd.AddCommand(oauthCmd)
}
