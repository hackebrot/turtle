package main

import (
	"fmt"
	"os"
)

func main() {
	if err := cmdTurtle.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
