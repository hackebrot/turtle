package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

var (
	cmdList = &cobra.Command{
		Use:   "list",
		Short: "Print a list of values from the turtle library",
		Long:  "Print a list of values from the turtle library",
	}

	cmdKeywords = &cobra.Command{
		Use:   "keywords",
		Short: "Print all keywords from the turtle library",
		Long:  "Print all keywords from the turtle library",
		Args:  cobra.NoArgs,
		RunE:  runKeywords,
	}

	cmdCategories = &cobra.Command{
		Use:   "categories",
		Short: "Print all categories from the turtle library",
		Long:  "Print all categories from the turtle library",
		Args:  cobra.NoArgs,
		RunE:  runCategories,
	}

	cmdNames = &cobra.Command{
		Use:   "names",
		Short: "Print all names from the turtle library",
		Long:  "Print all names from the turtle library",
		Args:  cobra.NoArgs,
		RunE:  runNames,
	}
)

func init() {
	cmdList.AddCommand(cmdCategories)
	cmdList.AddCommand(cmdKeywords)
	cmdList.AddCommand(cmdNames)
}

// addIfUnique adds strings to a given slice, if it
// cannot be found in the slice and then sorts the slice
func addIfUnique(r []string, sItems ...string) []string {
	for _, s := range sItems {
		i := sort.SearchStrings(r, s)

		if i >= len(r) || r[i] != s {
			r = append(r, s)
			sort.Strings(r)
		}
	}
	return r
}

func runKeywords(cmd *cobra.Command, args []string) error {
	var keywords []string

	for _, e := range turtle.Emojis {
		keywords = addIfUnique(keywords, e.Keywords...)
	}

	j, err := NewJSONWriter(os.Stdout, WithIndent(prefix, indent))
	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.Write(keywords)
}

func runCategories(cmd *cobra.Command, args []string) error {
	var categories []string

	for _, e := range turtle.Emojis {
		categories = addIfUnique(categories, e.Category)
	}

	j, err := NewJSONWriter(os.Stdout, WithIndent(prefix, indent))
	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.Write(categories)
}

func runNames(cmd *cobra.Command, args []string) error {
	var names []string

	for _, e := range turtle.Emojis {
		names = addIfUnique(names, e.Name)
	}

	j, err := NewJSONWriter(os.Stdout, WithIndent(prefix, indent))
	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.Write(names)
}
