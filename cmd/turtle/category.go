package main

import (
	"fmt"
	"io"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

func newCategoryCmd(w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "category [CATEGORY]",
		Short: "Print all emojis of the category",
		Long:  "Print all emojis of the category",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCategory(w, args[0])
		},
	}
}

func runCategory(w io.Writer, c string) error {

	emojis := turtle.Category(c)

	if emojis == nil {
		return fmt.Errorf("cannot find emojis of category %q", c)
	}

	j, err := NewJSONWriter(w, WithIndent(prefix, indent))

	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.WriteEmojis(emojis)
}
