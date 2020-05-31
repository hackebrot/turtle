package turtle

import (
	"fmt"
	"os"
)

// Example for using the Emojis map to find an
// emoji for the specified name.
func ExampleEmojis() {
	name := "turtle"
	emoji, ok := Emojis[name]

	if !ok {
		fmt.Fprintf(os.Stderr, "no emoji found for name: %v\n", name)
		os.Exit(1)
	}

	fmt.Printf("Name: %q\n", emoji.Name)
	fmt.Printf("Char: %s\n", emoji.Char)
	fmt.Printf("Category: %q\n", emoji.Category)
	fmt.Printf("Keywords: %q\n", emoji.Keywords)

	// Output:
	// Name: "turtle"
	// Char: 🐢
	// Category: "animals_and_nature"
	// Keywords: ["animal" "slow" "nature" "tortoise"]
}

// Example for using the EmojisByChar map to find an
// emoji for the specified emoji character.
func ExampleEmojisByChar() {
	char := "🐢"
	emoji, ok := EmojisByChar[char]

	if !ok {
		fmt.Fprintf(os.Stderr, "no emoji found for char: %v\n", char)
		os.Exit(1)
	}

	fmt.Printf("Name: %q\n", emoji.Name)
	fmt.Printf("Char: %s\n", emoji.Char)
	fmt.Printf("Category: %q\n", emoji.Category)
	fmt.Printf("Keywords: %q\n", emoji.Keywords)

	// Output:
	// Name: "turtle"
	// Char: 🐢
	// Category: "animals_and_nature"
	// Keywords: ["animal" "slow" "nature" "tortoise"]
}

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

// Example for using the Filter function to find all
// emojis in a category with a specific keyword
func ExampleFilter() {
	emojis := Filter(func(e *Emoji) bool {
		if e.Category == "animals_and_nature" {
			for _, keyword := range e.Keywords {
				if keyword == "world" {
					return true
				}
			}
		}
		return false
	})

	if emojis == nil {
		fmt.Fprintf(os.Stderr, "no emojis found in animals_and_nature with keyword world\n")
		os.Exit(1)
	}

	fmt.Printf("emojis: %s", emojis)

	// Output:
	// emojis: [🌍 🌎 🌏]
}
