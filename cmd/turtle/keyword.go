package main

import (
	"fmt"
	"os"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

var (
	cmdKeyword = &cobra.Command{
		Use:   "keyword",
		Short: "Print all emojis with the keyword",
		Long:  "Print all emojis with the keyword",
		RunE:  runKeyword,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("require one keyword")
			}
			return nil
		},
	}
)

func runKeyword(cmd *cobra.Command, args []string) error {
	k := args[0]

	emojis := turtle.Keyword(k)

	if emojis == nil {
		return fmt.Errorf("cannot find emojis for keyword %q", k)
	}

	var err error
	var w Writer
	if plain {
		w, err = NewPlainWriter(os.Stdout)
	} else {
		w, err = NewJSONWriter(os.Stdout, WithIndent(prefix, indent))
	}
	if err != nil {
		return fmt.Errorf("error creating Writer: %v", err)
	}

	return w.WriteEmojis(emojis)
}
