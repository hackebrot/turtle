package turtle

import (
	"fmt"
	"os"
)

// Example for using the Category function to find all
// emojis of the specified category.
func ExampleCategory() {
	c := "travel_and_places"
	emojis := Category(c)

	if emojis == nil {
		fmt.Fprintf(os.Stderr, "no emojis found for category: %v\n", c)
		os.Exit(1)
	}

	fmt.Printf("%s: %s\n", c, emojis[:3])

	// Output:
	// travel_and_places: [🚡 ✈️ 🚑]
}

// Example for using the Keyword function to find all
// emojis with the specified keyword.
func ExampleKeyword() {
	k := "happy"
	emojis := Keyword(k)

	if emojis == nil {
		fmt.Fprintf(os.Stderr, "no emoji found for keyword: %v\n", k)
		os.Exit(1)
	}

	fmt.Printf("%s: %s", k, emojis[:4])

	// Output:
	// happy: [😊 😁 😀 😂]
}

// Example for using the Search function to find all
// emojis with a name that contains the search string.
func ExampleSearch() {
	s := "computer"
	emojis := Search(s)

	if emojis == nil {
		fmt.Fprintf(os.Stderr, "no emojis found for search: %v\n", s)
		os.Exit(1)
	}

	fmt.Printf("%s: %s", s, emojis)

	// Output:
	// computer: [💻 🖱 🖥]
}
