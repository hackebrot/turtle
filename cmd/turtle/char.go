package main

import (
	"fmt"
	"io"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

func newCharCmd(w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "char",
		Short: "Print the emoji for the emoji character",
		Long:  "Print the emoji for the emoji character",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runChar(w, args[0])
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("require one emoji character")
			}
			return nil
		},
	}
}

func runChar(w io.Writer, char string) error {

	e, ok := turtle.EmojisByChar[char]

	if !ok {
		return fmt.Errorf("cannot find emoji with emoji character %q", char)
	}

	j, err := NewJSONWriter(w, WithIndent(prefix, indent))

	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.WriteEmoji(e)
}
