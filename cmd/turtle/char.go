package main

import (
	"fmt"
	"io"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

func newCharCmd(w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "char [CHAR]",
		Short: "Print the emoji for the emoji character",
		Long:  "Print the emoji for the emoji character",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runChar(w, args[0])
		},
	}
}

func runChar(w io.Writer, c string) error {

	e, ok := turtle.EmojisByChar[c]

	if !ok {
		return fmt.Errorf("cannot find emoji with emoji character %q", c)
	}

	j, err := NewJSONWriter(w, WithIndent(prefix, indent))

	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.WriteEmoji(e)
}
