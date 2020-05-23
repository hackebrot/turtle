package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

var (
	cmdRandom = &cobra.Command{
		Use:   "random",
		Short: "Print a random emoji",
		Long:  "Print a random emoji",
		RunE:  runRandom,
		Args:  nil,
	}
)

func runRandom(cmd *cobra.Command, args []string) error {

	emojis := make([]*turtle.Emoji, 0, len(turtle.Emojis))

	for _, value := range turtle.Emojis {
		emojis = append(emojis, value)
	}

	j, err := NewJSONWriter(os.Stdout, WithIndent(prefix, indent))

	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.WriteEmoji(emojis[rand.Intn(len(emojis))])
}

func init() {
	rand.Seed(time.Now().Unix())
}
