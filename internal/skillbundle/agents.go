// Package skillbundle implements the `acmctl skills install` and
// `acmctl skills update` commands: download a published skill zip
// from a GitHub release, extract it, and copy into the appropriate
// per-agent directory under a chosen scope (global = $HOME, project
// = current working directory).
package skillbundle

import (
	"fmt"
	"sort"
	"strings"
)

// Agent describes a known consumer of skill bundles. The install
// path is rooted at the chosen scope (e.g., "$HOME/.claude/skills"
// for --scope global, "$CWD/.claude/skills" for --scope project).
type Agent struct {
	// Name is the user-facing identifier (e.g., "claude").
	Name string
	// InstallSubdir is the path relative to the scope root where
	// skills land for this agent (e.g., ".claude/skills").
	InstallSubdir string
}

// Agents lists every agent acmctl knows how to install skills for.
// To add another agent, append one entry here.
var Agents = []Agent{
	{Name: "claude", InstallSubdir: ".claude/skills"},
	{Name: "codex", InstallSubdir: ".codex/skills"},
}

// AgentByName returns the Agent matching the given name, or an
// error listing supported agents.
func AgentByName(name string) (Agent, error) {
	for _, a := range Agents {
		if a.Name == name {
			return a, nil
		}
	}
	return Agent{}, fmt.Errorf("unknown agent %q (supported: %s)", name, supportedNames())
}

// SelectAgents resolves the user's flag combination into a concrete
// list of agents to install for.
//
//   - all=true returns every agent in Agents (--all)
//   - len(names)>0 returns the named subset (--agent claude --agent codex)
//   - both unset defaults to claude (the most common case)
func SelectAgents(names []string, all bool) ([]Agent, error) {
	if all {
		out := make([]Agent, len(Agents))
		copy(out, Agents)
		return out, nil
	}
	if len(names) == 0 {
		a, err := AgentByName("claude")
		if err != nil {
			return nil, err
		}
		return []Agent{a}, nil
	}
	seen := map[string]bool{}
	out := make([]Agent, 0, len(names))
	for _, n := range names {
		if seen[n] {
			continue
		}
		seen[n] = true
		a, err := AgentByName(n)
		if err != nil {
			return nil, err
		}
		out = append(out, a)
	}
	return out, nil
}

func supportedNames() string {
	names := make([]string, len(Agents))
	for i, a := range Agents {
		names[i] = a.Name
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}
