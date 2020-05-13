package main

import (
	"fmt"
	"os"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

var (
	cmdChar = &cobra.Command{
		Use:   "char",
		Short: "Print the emoji for the emoji character",
		Long:  "Print the emoji for the emoji character",
		RunE:  runChar,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("require one emoji character")
			}
			return nil
		},
	}
)

func runChar(cmd *cobra.Command, args []string) error {
	char := args[0]

	e, ok := turtle.EmojisByChar[char]

	if !ok {
		return fmt.Errorf("cannot find emoji with emoji character %q", char)
	}

	j, err := NewJSONWriter(os.Stdout, WithIndent(prefix, indent))

	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.WriteEmoji(e)
}
