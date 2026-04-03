package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"gopkg.in/yaml.v3"
)

// Tabulable allows a model to be rendered as a table row.
type Tabulable interface {
	TableHeaders() []string
	TableRow() []string
}

func Print(format string, data interface{}) error {
	switch format {
	case "json":
		return printJSON(os.Stdout, data)
	case "yaml":
		return printYAML(os.Stdout, data)
	default:
		return printTable(os.Stdout, data)
	}
}

func printJSON(w io.Writer, data interface{}) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	return enc.Encode(data)
}

func printYAML(w io.Writer, data interface{}) error {
	return yaml.NewEncoder(w).Encode(data)
}

func printTable(w io.Writer, data interface{}) error {
	switch v := data.(type) {
	case Tabulable:
		return printTableRows(w, v.TableHeaders(), []Tabulable{v})
	case []Tabulable:
		if len(v) == 0 {
			fmt.Fprintln(w, "No resources found.")
			return nil
		}
		return printTableRows(w, v[0].TableHeaders(), v)
	default:
		return printJSON(w, data)
	}
}

func printTableRows(w io.Writer, headers []string, rows []Tabulable) error {
	tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)
	for i, h := range headers {
		if i > 0 {
			fmt.Fprint(tw, "\t")
		}
		fmt.Fprint(tw, h)
	}
	fmt.Fprintln(tw)
	for _, row := range rows {
		vals := row.TableRow()
		for i, v := range vals {
			if i > 0 {
				fmt.Fprint(tw, "\t")
			}
			fmt.Fprint(tw, v)
		}
		fmt.Fprintln(tw)
	}
	return tw.Flush()
}

// PrintTabulableList is a helper for printing typed slices.
func PrintTabulableList[T Tabulable](format string, items []T) error {
	if format == "json" || format == "yaml" {
		return Print(format, items)
	}
	if len(items) == 0 {
		fmt.Println("No resources found.")
		return nil
	}
	tbl := make([]Tabulable, len(items))
	for i := range items {
		tbl[i] = items[i]
	}
	return Print(format, tbl)
}
