package main

import (
	"fmt"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

var (
	cmdTurtle = &cobra.Command{
		Use:   "turtle",
		Short: "Print the emoji with name",
		Long:  "Print the emoji with name",
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
	cmdTurtle.AddCommand(cmdVersion)
}

func runTurtle(cmd *cobra.Command, args []string) error {
	name := args[0]

	e, ok := turtle.Emojis[name]

	if !ok {
		return fmt.Errorf("cannot find emoji with name %q", name)
	}

	fmt.Printf("%s\n", e)
	return nil
}
