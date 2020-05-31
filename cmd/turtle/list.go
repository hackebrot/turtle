package main

import (
	"fmt"
	"io"
	"sort"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

func newListCmd(w io.Writer) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Print a list of values from the turtle library",
		Long: `Print a list of values from the turtle library.

List requires a subcommand, e.g. turte list categories`,
		Args: cobra.NoArgs,
	}
	categoriesCmd := &cobra.Command{
		Use:   "categories",
		Short: "Print all categories from the turtle library",
		Long:  "Print all categories from the turtle library",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCategories(w)
		},
	}
	keywordsCmd := &cobra.Command{
		Use:   "keywords",
		Short: "Print all keywords from the turtle library",
		Long:  "Print all keywords from the turtle library",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runKeywords(w)
		},
	}
	namesCmd := &cobra.Command{
		Use:   "names",
		Short: "Print all names from the turtle library",
		Long:  "Print all names from the turtle library",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNames(w)
		},
	}
	listCmd.AddCommand(categoriesCmd)
	listCmd.AddCommand(keywordsCmd)
	listCmd.AddCommand(namesCmd)
	return listCmd
}

// addIfUnique adds strings to a given slice, if it
// cannot be found in the slice and then sorts the slice
func addIfUnique(r []string, sItems ...string) []string {

	// The slice must be sorted in ascending order
	// before we use it in sort.SearchStrings
	sort.Strings(r)

	for _, s := range sItems {
		i := sort.SearchStrings(r, s)

		if i >= len(r) || r[i] != s {
			// r does not contain s, add it and sort
			r = append(r, s)
			sort.Strings(r)
		}
	}

	return r
}

func runKeywords(w io.Writer) error {
	var keywords []string

	for _, e := range turtle.Emojis {
		keywords = addIfUnique(keywords, e.Keywords...)
	}

	j, err := NewJSONWriter(w, WithIndent(prefix, indent))
	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.Write(keywords)
}

func runCategories(w io.Writer) error {
	var categories []string

	for _, e := range turtle.Emojis {
		categories = addIfUnique(categories, e.Category)
	}

	j, err := NewJSONWriter(w, WithIndent(prefix, indent))
	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.Write(categories)
}

func runNames(w io.Writer) error {
	var names []string

	for _, e := range turtle.Emojis {
		names = addIfUnique(names, e.Name)
	}

	j, err := NewJSONWriter(w, WithIndent(prefix, indent))
	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.Write(names)
}
