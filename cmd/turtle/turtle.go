package main

import (
	"fmt"
	"io"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

var indent string
var prefix string

// Create a new turtle command with the given output
func newTurtleCmd(w io.Writer) *cobra.Command {
	turtleCmd := &cobra.Command{
		Use:   "turtle [NAME]",
		Short: "Print the emoji with the specified name identifier",
		Long:  "Print the emoji with the specified name identifier",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTurtle(w, args[0])
		},
		Version: turtle.Version,
	}

	turtleCmd.PersistentFlags().StringVarP(&indent, "indent", "i", "", "indent for JSON output")
	turtleCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "prefix for JSON output")

	turtleCmd.SetVersionTemplate("{{with .Name}}{{printf \"%s \" .}}{{end}}{{printf \"%s\" .Version}}\n")

	turtleCmd.AddCommand(newCharCmd(w))
	turtleCmd.AddCommand(newCategoryCmd(w))
	turtleCmd.AddCommand(newKeywordCmd(w))
	turtleCmd.AddCommand(newSearchCmd(w))
	turtleCmd.AddCommand(newListCmd(w))
	turtleCmd.AddCommand(newRandomCmd(w))

	return turtleCmd
}

func runTurtle(w io.Writer, name string) error {

	e, ok := turtle.Emojis[name]

	if !ok {
		return fmt.Errorf("cannot find emoji with name %q", name)
	}

	j, err := NewJSONWriter(w, WithIndent(prefix, indent))

	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.WriteEmoji(e)
}
