package main

import (
	"fmt"
	"io"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

func newKeywordCmd(w io.Writer) *cobra.Command {

	keywordCmd := &cobra.Command{
		Use:   "keyword",
		Short: "Print all emojis with the keyword",
		Long:  "Print all emojis with the keyword",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runKeyword(w, args[0])
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("require one keyword")
			}
			return nil
		},
	}
	return keywordCmd
}

func runKeyword(w io.Writer, k string) error {

	emojis := turtle.Keyword(k)

	if emojis == nil {
		return fmt.Errorf("cannot find emojis for keyword %q", k)
	}

	j, err := NewJSONWriter(w, WithIndent(prefix, indent))

	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.WriteEmojis(emojis)
}
