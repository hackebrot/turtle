package main

import (
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

func newRandomCmd(w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "random",
		Short: "Print a random emoji",
		Long:  "Print a random emoji",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRandom(w)
		},
	}
}

func runRandom(w io.Writer) error {

	emojis := make([]*turtle.Emoji, 0, len(turtle.Emojis))

	for _, value := range turtle.Emojis {
		emojis = append(emojis, value)
	}

	j, err := NewJSONWriter(w, WithIndent(prefix, indent))

	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.WriteEmoji(emojis[rand.Intn(len(emojis))])
}

func init() {
	rand.Seed(time.Now().Unix())
}
