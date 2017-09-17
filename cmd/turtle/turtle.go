package main

import (
	"fmt"
	"os"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

var (
	indent string
	prefix string

	cmdTurtle = &cobra.Command{
		Use:   "turtle",
		Short: "Print the unicode character for a given emoji name",
		Long:  "Print the unicode character for a given emoji name",
		RunE:  runTurtle,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("require emoji name")
			}
			return nil
		},
	}
)

func init() {
	cmdTurtle.AddCommand(cmdCategory)
	cmdTurtle.AddCommand(cmdKeyword)
	cmdTurtle.AddCommand(cmdSearch)
	cmdTurtle.AddCommand(cmdVersion)

	cmdTurtle.PersistentFlags().StringVarP(&indent, "indent", "i", "", "indent for JSON output")
	cmdTurtle.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "prefix for JSON output")
}

func runTurtle(cmd *cobra.Command, args []string) error {
	name := args[0]

	e, ok := turtle.Emojis[name]

	if !ok {
		return fmt.Errorf("cannot find emoji with name %q", name)
	}

	j, err := NewJSONWriter(os.Stdout, WithIndent(prefix, indent))

	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.WriteEmoji(e)
}
