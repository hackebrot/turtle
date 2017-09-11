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
			if len(args) == 0 {
				return fmt.Errorf("require keyword")
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
	j := NewJSONWriter(os.Stdout)

	return j.WriteEmojis(emojis)
}
