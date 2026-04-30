package cmd

import (
	"fmt"
	"os"

	"github.com/altinity/acmctl/pkg/api"
	"github.com/altinity/acmctl/pkg/config"
	"github.com/spf13/cobra"
)

var (
	cfgFile   string
	tokenFlag string
	urlFlag   string
	verbose   bool
	cfg       *config.Config
	apiClient *api.Client
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

		// Flag/env overrides
		if tokenFlag != "" {
			cfg.Token = tokenFlag
		} else if envToken := os.Getenv("ACM_API_KEY"); envToken != "" {
			cfg.Token = envToken
		}
		if urlFlag != "" {
			cfg.URL = urlFlag
		} else if envURL := os.Getenv("ACMCTL_URL"); envURL != "" {
			cfg.URL = envURL
		}

		apiClient = api.NewClient(cfg.URL, cfg.Token)
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
	// Default is ~/.acmctl.yaml; degrade to "" if HOME isn't set so
	// the user gets a clearer "config not found" message at load time
	// than a startup panic from init.
	defaultCfg, _ := config.DefaultPath()
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", defaultCfg, "config file path")
	rootCmd.PersistentFlags().StringVar(&tokenFlag, "token", "", "API token (overrides config)")
	rootCmd.PersistentFlags().StringVar(&urlFlag, "url", "", "API base URL (overrides config)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
