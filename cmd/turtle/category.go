package main

import (
	"fmt"
	"os"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

var (
	cmdCategory = &cobra.Command{
		Use:   "category",
		Short: "Print all emojis of the category",
		Long:  "Print all emojis of the category",
		RunE:  runCategory,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("require category")
			}
			return nil
		},
	}
)

func runCategory(cmd *cobra.Command, args []string) error {
	c := args[0]

	emojis := turtle.Category(c)

	if emojis == nil {
		return fmt.Errorf("cannot find emojis of category %q", c)
	}

	j, err := NewJSONWriter(os.Stdout, WithIndent(prefix, indent))

	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.WriteEmojis(emojis)
}
