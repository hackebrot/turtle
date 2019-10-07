package main

import (
	"fmt"
	"os"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

var (
	cmdName = &cobra.Command{
		Use:   "name",
		Short: "Print the name of an emoji",
		Long:  "Print the name of an emoji",
		RunE:  runName,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("require one name")
			}
			return nil
		},
	}
)

func runName(cmd *cobra.Command, args []string) error {
	k := args[0]

	emoji := turtle.Name(k)

	if emoji == nil {
		return fmt.Errorf("cannot find emoji name for %q", k)
	}

	j, err := NewJSONWriter(os.Stdout, WithIndent(prefix, indent))

	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.Write(emoji.Name)
}
