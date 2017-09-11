package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/hackebrot/turtle"
)

// writeEmoiji writes attributes of an Emoji to a io.Writer
func writeEmoji(w io.Writer, f string, e *turtle.Emoji, name, char, category, keywords bool) {
	// By default only write the emoji Char to w
	if !(name || char || category || keywords) {
		fmt.Fprintf(w, f, e.Char)
		return
	}

	// Write attributes based on the CLI flags to w
	var out []string

	if name {
		out = append(out, fmt.Sprintf("Name: %q", e.Name))
	}

	if char {
		out = append(out, fmt.Sprintf("Char: %s", e.Char))
	}

	if category {
		out = append(out, fmt.Sprintf("Category: %q", e.Category))
	}

	if keywords {
		out = append(out, fmt.Sprintf("Keywords: %q", e.Keywords))
	}

	fmt.Fprintf(w, f, strings.Join(out, "\n"))
}
