package cmd

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/altinity/acmctl/pkg/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage acmctl profiles (prod/dev/stage/...)",
	Long: `acmctl supports multiple ACM environments via profiles. Each
profile holds a URL and an optional token. Switch between them with
` + "`acmctl --profile <name>`" + ` for one command, or
` + "`acmctl config use-profile <name>`" + ` to change the default.`,
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cfg == nil || len(cfg.Profiles) == 0 {
			fmt.Fprintln(os.Stderr, "no profiles configured (try `acmctl config add-profile <name> --url <url>`)")
			return nil
		}
		names := make([]string, 0, len(cfg.Profiles))
		for k := range cfg.Profiles {
			names = append(names, k)
		}
		sort.Strings(names)

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "DEFAULT\tNAME\tURL\tTOKEN")
		for _, n := range names {
			p := cfg.Profiles[n]
			def := ""
			if n == cfg.DefaultProfile {
				def = "*"
			}
			tok := "(unset)"
			if p.Token != "" {
				tok = "set"
			}
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", def, n, p.URL, tok)
		}
		return w.Flush()
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Show the active profile's name and URL",
	RunE: func(cmd *cobra.Command, args []string) error {
		if activeProfile == "" {
			fmt.Fprintln(os.Stderr, "no profile selected; pass --profile or set default with `acmctl config use-profile <name>`")
			return nil
		}
		p := cfg.Profiles[activeProfile]
		fmt.Printf("name:  %s\n", activeProfile)
		fmt.Printf("url:   %s\n", p.URL)
		if p.Token == "" {
			fmt.Println("token: (unset — run `acmctl oauth` or `acmctl login`)")
		} else {
			fmt.Println("token: set")
		}
		return nil
	},
}

var configUseProfileCmd = &cobra.Command{
	Use:   "use-profile <name>",
	Short: "Set the default profile for future commands",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		if _, ok := cfg.Profiles[name]; !ok {
			return fmt.Errorf("profile %q not found; create it first with `acmctl config add-profile %s --url <url>`", name, name)
		}
		cfg.DefaultProfile = name
		if err := config.Save(cfgFile, cfg); err != nil {
			return fmt.Errorf("save config: %w", err)
		}
		fmt.Printf("default profile set to %q\n", name)
		return nil
	},
}

var configAddProfileCmd = &cobra.Command{
	Use:   "add-profile <name>",
	Short: "Create a new profile (token unset; populate via `acmctl --profile <name> oauth`)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		urlVal, _ := cmd.Flags().GetString("url")
		if urlVal == "" {
			return fmt.Errorf("--url is required (e.g., --url https://acm.altinity.cloud/api)")
		}
		if _, exists := cfg.Profiles[name]; exists {
			return fmt.Errorf("profile %q already exists; remove it first or use a different name", name)
		}
		cfg.SetProfile(name, config.Profile{URL: urlVal})
		// First profile becomes the default automatically.
		if cfg.DefaultProfile == "" {
			cfg.DefaultProfile = name
		}
		if err := config.Save(cfgFile, cfg); err != nil {
			return fmt.Errorf("save config: %w", err)
		}
		fmt.Printf("added profile %q (url=%s)", name, urlVal)
		if cfg.DefaultProfile == name {
			fmt.Printf(", set as default")
		}
		fmt.Println()
		fmt.Printf("next: acmctl --profile %s oauth\n", name)
		return nil
	},
}

var configRemoveProfileCmd = &cobra.Command{
	Use:   "remove-profile <name>",
	Short: "Delete a profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		if _, ok := cfg.Profiles[name]; !ok {
			return fmt.Errorf("profile %q not found", name)
		}
		cfg.RemoveProfile(name)
		if err := config.Save(cfgFile, cfg); err != nil {
			return fmt.Errorf("save config: %w", err)
		}
		fmt.Printf("removed profile %q\n", name)
		if cfg.DefaultProfile == "" && len(cfg.Profiles) > 0 {
			fmt.Fprintln(os.Stderr, "note: no default profile is set; choose one with `acmctl config use-profile <name>`")
		}
		return nil
	},
}

func init() {
	configAddProfileCmd.Flags().String("url", "", "ACM API base URL for this profile (required)")
	_ = configAddProfileCmd.MarkFlagRequired("url")

	configCmd.AddCommand(configListCmd, configGetCmd, configUseProfileCmd, configAddProfileCmd, configRemoveProfileCmd)
	rootCmd.AddCommand(configCmd)
}
