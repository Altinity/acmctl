package cmd

import (
	"fmt"
	"os"

	"github.com/altinity/acmctl/pkg/api"
	"github.com/altinity/acmctl/pkg/config"
	"github.com/spf13/cobra"
)

var (
	cfgFile        string
	tokenFlag      string
	urlFlag        string
	profileFlag    string
	verbose        bool
	cfg            *config.Config
	apiClient      *api.Client
	activeProfile  string // resolved profile name; "" before PersistentPreRunE
)

var rootCmd = &cobra.Command{
	Use:   "acmctl",
	Short: "CLI for Altinity Cloud Manager",
	Long:  "acmctl is a command-line tool for managing Altinity Cloud Manager resources.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Skip client init for commands that don't need it.
		if cmd.Name() == "version" || cmd.Name() == "completion" || cmd.Name() == "help" {
			return nil
		}
		// `acmctl skills install` / `update` fetch from GitHub —
		// no ACM API client needed. Match any subcommand under
		// the `skills` parent.
		for c := cmd; c != nil; c = c.Parent() {
			if c.Name() == "skills" {
				return nil
			}
		}

		var err error
		cfg, err = config.Load(cfgFile)
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		// Resolve the active profile (flag > env > default_profile >
		// single-profile fallback).
		profile, name, perr := cfg.ActiveProfile(profileFlag)

		// If the user EXPLICITLY selected a profile that doesn't
		// exist (--profile <name> or ACMCTL_PROFILE=<name>), fail
		// here with a clear error. If they made no selection and
		// there's no fallback, defer the error to commands that need
		// a profile via requireProfile() — `config list/add-profile`
		// must keep working on empty configs.
		if perr != nil {
			explicit := profileFlag != "" || os.Getenv("ACMCTL_PROFILE") != ""
			if explicit {
				return perr
			}
			activeProfile = ""
		} else {
			activeProfile = name
		}

		// Compose effective URL/token: flag > env > active profile.
		url := profile.URL
		if urlFlag != "" {
			url = urlFlag
		} else if envURL := os.Getenv("ACMCTL_URL"); envURL != "" {
			url = envURL
		}
		if url == "" {
			url = config.DefaultURL
		}

		token := profile.Token
		if tokenFlag != "" {
			token = tokenFlag
		} else if envToken := os.Getenv("ACM_API_KEY"); envToken != "" {
			token = envToken
		}

		apiClient = api.NewClient(url, token)
		apiClient.Verbose = verbose
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	defaultCfg, _ := config.DefaultPath()
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", defaultCfg, "config file path")
	rootCmd.PersistentFlags().StringVar(&profileFlag, "profile", "",
		"profile name to use (default: $ACMCTL_PROFILE or default_profile in config)")
	rootCmd.PersistentFlags().StringVar(&tokenFlag, "token", "", "API token (overrides config + env)")
	rootCmd.PersistentFlags().StringVar(&urlFlag, "url", "", "API base URL (overrides config + env)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

// requireProfile returns the active profile name or a clear error
// for commands that strictly require one (cluster list, env get,
// etc.). Use this inside a RunE before touching apiClient on a path
// that needs a real URL+token.
func requireProfile() (string, error) {
	if activeProfile == "" {
		return "", fmt.Errorf("no profile selected — pass --profile, set ACMCTL_PROFILE, or run `acmctl config use-profile <name>`")
	}
	// A flag/env URL override is enough — don't insist on a profile URL
	// when the caller has supplied one out-of-band.
	if urlFlag != "" || os.Getenv("ACMCTL_URL") != "" {
		return activeProfile, nil
	}
	if cfg == nil || cfg.Profiles[activeProfile].URL == "" {
		return activeProfile, fmt.Errorf("profile %q has no url set; edit %s or run `acmctl config add-profile %s --url <url>`", activeProfile, cfgFile, activeProfile)
	}
	return activeProfile, nil
}
