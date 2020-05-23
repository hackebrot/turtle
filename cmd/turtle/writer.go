package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hackebrot/turtle"
)

// JSONWriter writes Emojis as JSON
type JSONWriter struct {
	e *json.Encoder
}

// NewJSONWriter creates a new JSONWriter
func NewJSONWriter(w io.Writer, options ...func(*JSONWriter) error) (*JSONWriter, error) {
	j := &JSONWriter{e: json.NewEncoder(w)}

	for _, option := range options {
		if err := option(j); err != nil {
			return nil, fmt.Errorf("error applying option: %v", err)
		}
	}

	return j, nil
}

// WithIndent sets an indent on  adds a separator to a thread
func WithIndent(prefix, indent string) func(*JSONWriter) error {
	return func(j *JSONWriter) error {
		j.e.SetIndent(prefix, indent)
		return nil
	}
}

// WriteEmoji to an io.Writer as JSON
func (j *JSONWriter) WriteEmoji(emoji *turtle.Emoji) error {
	return j.e.Encode(emoji)
}

// WriteEmojis to an io.Writer as JSON
func (j *JSONWriter) WriteEmojis(emojis []*turtle.Emoji) error {
	return j.e.Encode(emojis)
}

// Write a value to an io.Writer as JSON
func (j *JSONWriter) Write(v interface{}) error {
	return j.e.Encode(v)
}
