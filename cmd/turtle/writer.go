package main

import (
	"encoding/json"
	"io"

	"github.com/hackebrot/turtle"
)

// JSONWriter writes Emojis as JSON
type JSONWriter struct {
	e *json.Encoder
}

// NewJSONWriter creates a new JSONWriter
func NewJSONWriter(w io.Writer) *JSONWriter {
	return &JSONWriter{e: json.NewEncoder(w)}
}

// WriteEmoji to an io.Writer
func (j *JSONWriter) WriteEmoji(emoji *turtle.Emoji) error {
	return j.e.Encode(emoji)
}

// WriteEmojis to an io.Writer
func (j *JSONWriter) WriteEmojis(emojis []*turtle.Emoji) error {
	return j.e.Encode(emojis)
}
