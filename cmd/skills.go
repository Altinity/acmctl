package cmd

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/altinity/acmctl/internal/skillbundle"
	"github.com/spf13/cobra"
)

// Skills shipped with acmctl. Each becomes <skill>.zip in the
// GitHub release; the install/update commands iterate over this
// list. To add another skill: drop a directory under skills/ in
// this repo, append its name here.
var bundledSkills = []string{"altinity-cloud"}

var skillsCmd = &cobra.Command{
	Use:   "skills",
	Short: "Manage AI agent skills bundled with acmctl",
	Long: `Install or refresh the Claude Code / Codex skills published from
this repo to per-agent skill directories under your home or current
directory.

Skill content is downloaded from the rolling "skill-bundle" GitHub
release on every run (no embedded copy in the binary), so updates
ship without rebuilding acmctl.`,
}

var skillsInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Download and install skills (errors on conflicts unless --force)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runSkills(cmd, false)
	},
}

var skillsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Refresh already-installed skills (always overwrites)",
	Long: `Refresh skills that are already installed. Differs from "install"
in two ways: (1) by default it targets every agent that has the
skill installed (not just claude); (2) it always overwrites differing
files, since the goal is to bring local content back in line with
the published bundle.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runSkills(cmd, true)
	},
}

func init() {
	for _, sub := range []*cobra.Command{skillsInstallCmd, skillsUpdateCmd} {
		sub.Flags().StringSlice("agent", nil,
			"target agent (repeatable; default: claude for install, all detected for update)")
		sub.Flags().Bool("all", false,
			"install for every supported agent (claude, codex)")
		sub.Flags().String("scope", "global",
			"install scope: 'global' (under $HOME) or 'project' (under cwd)")
		sub.Flags().String("ref", skillbundle.DefaultRef,
			"GitHub release tag to fetch from")
		sub.Flags().Bool("dry-run", false,
			"print what would happen, no writes")
	}
	// Only install accepts --force; for update, force is implicit.
	skillsInstallCmd.Flags().Bool("force", false,
		"overwrite local edits to skill files")

	skillsCmd.AddCommand(skillsInstallCmd, skillsUpdateCmd)
	rootCmd.AddCommand(skillsCmd)
}

// runSkills implements both `install` and `update`. update=true
// means the caller expects the skill to already be installed for at
// least one agent and wants overwrite-on-diff semantics.
func runSkills(cmd *cobra.Command, update bool) error {
	agentNames, _ := cmd.Flags().GetStringSlice("agent")
	all, _ := cmd.Flags().GetBool("all")
	scope, _ := cmd.Flags().GetString("scope")
	ref, _ := cmd.Flags().GetString("ref")
	dryRun, _ := cmd.Flags().GetBool("dry-run")
	// `update` doesn't expose --force; treat it as implicitly forced.
	force := update
	if !update {
		force, _ = cmd.Flags().GetBool("force")
	}

	scopeRoot, err := skillbundle.ScopeRoot(scope)
	if err != nil {
		return err
	}

	// Resolve which agents to operate on.
	var agents []skillbundle.Agent
	if update && len(agentNames) == 0 && !all {
		// Default for update: every agent that already has any of
		// the bundled skills installed.
		for _, a := range skillbundle.Agents {
			for _, sk := range bundledSkills {
				if skillbundle.IsInstalled(scopeRoot, a, sk) {
					agents = append(agents, a)
					break
				}
			}
		}
		if len(agents) == 0 {
			return fmt.Errorf("no skills are installed for any known agent under %s; run `acmctl skills install` first", scopeRoot)
		}
	} else {
		agents, err = skillbundle.SelectAgents(agentNames, all)
		if err != nil {
			return err
		}
	}

	ctx := context.Background()

	// Cache: fetch each unique skill at most once even if multiple
	// agents reuse it.
	bundles := map[string]*skillbundle.Bundle{}
	for _, sk := range bundledSkills {
		fmt.Fprintf(os.Stderr, "fetching %s.zip from ref %q...\n", sk, ref)
		b, err := skillbundle.Fetch(ctx, sk, ref)
		if err != nil {
			return err
		}
		bundles[sk] = b
	}

	totalCreated, totalUpdated, totalUnchanged := 0, 0, 0
	var conflicts []string

	for _, a := range agents {
		for _, sk := range bundledSkills {
			b := bundles[sk]
			summary, err := skillbundle.Install(b, skillbundle.InstallOpts{
				ScopeRoot: scopeRoot,
				Agent:     a,
				Force:     force,
				DryRun:    dryRun,
			})
			if err != nil {
				return err
			}
			totalCreated += summary.Created
			totalUpdated += summary.Updated
			totalUnchanged += summary.Unchanged
			conflicts = append(conflicts, summary.Conflicts...)

			label := fmt.Sprintf("%s -> %s", sk, a.Name)
			if dryRun {
				label = "[dry-run] " + label
			}
			fmt.Fprintf(os.Stderr, "  %s: %d created, %d updated, %d unchanged",
				label, summary.Created, summary.Updated, summary.Unchanged)
			if len(summary.Conflicts) > 0 {
				fmt.Fprintf(os.Stderr, ", %d would-overwrite", len(summary.Conflicts))
			}
			fmt.Fprintln(os.Stderr)
		}
	}

	if len(conflicts) > 0 {
		// install (without --force) aborts before reporting "done".
		// update can't reach this branch because force==true.
		sort.Strings(conflicts)
		fmt.Fprintf(os.Stderr, "\n%d file(s) would be overwritten:\n", len(conflicts))
		for _, c := range conflicts {
			fmt.Fprintf(os.Stderr, "  %s\n", c)
		}
		return fmt.Errorf("refusing to overwrite local edits; rerun with --force, or use `acmctl skills update` if this is an existing install")
	}

	verb := "installed"
	if update {
		verb = "updated"
	}
	if dryRun {
		verb = "(dry-run) " + verb
	}
	fmt.Fprintf(os.Stderr, "\n%s: %d created, %d updated, %d unchanged\n",
		verb, totalCreated, totalUpdated, totalUnchanged)
	return nil
}
