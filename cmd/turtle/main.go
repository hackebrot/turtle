package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := newTurtleCmd(os.Stdout)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
