package cmd

import (
	"encoding/json"
	"os"

	"golang.org/x/term"
)

// printJSON writes value as indented JSON to stdout. value can be json.RawMessage,
// any encodable Go type, or nil (no-op).
func printJSON(v interface{}) error {
	if v == nil {
		return nil
	}
	if raw, ok := v.(json.RawMessage); ok && len(raw) == 0 {
		return nil
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

// isatty reports whether the file descriptor refers to a terminal.
func isatty(f *os.File) bool {
	return term.IsTerminal(int(f.Fd()))
}

// stdinHasBody reports whether stdin could carry a request body: a pipe or a
// regular file. Character devices — terminals AND /dev/null — return false, so
// bodyless calls never block on io.ReadAll waiting for an EOF that won't come.
func stdinHasBody() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice == 0
}
