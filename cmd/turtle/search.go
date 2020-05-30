package main

import (
	"fmt"
	"io"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

func newSearchCmd(w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "search",
		Short: "Print emojis with a name that contains the search string",
		Long:  "Print emojis with a name that contains the search string",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSearch(w, args[0])
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("require one search string")
			}
			return nil
		},
	}
}

func runSearch(w io.Writer, s string) error {

	emojis := turtle.Search(s)

	if emojis == nil {
		return fmt.Errorf("cannot find emojis for search %q", s)
	}

	j, err := NewJSONWriter(w, WithIndent(prefix, indent))

	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.WriteEmojis(emojis)
}
