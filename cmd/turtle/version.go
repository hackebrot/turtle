package main

import (
	"fmt"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

var (
	cmdVersion = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of turtle",
		Long:  "Print the version number of turtle",
		Run:   runVersion,
	}
)

func runVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("turtle %s\n", turtle.Version)
}
