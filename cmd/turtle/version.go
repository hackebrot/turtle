package main

import (
	"fmt"
	"io"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

func newVersionCmd(w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of turtle",
		Long:  "Print the version number of turtle",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runVersion(w)
		},
	}
}

func runVersion(w io.Writer) error {

	if _, err := fmt.Fprintf(w, "turtle %s\n", turtle.Version); err != nil {
		return fmt.Errorf("error writing version number: %v", err)
	}

	return nil
}
