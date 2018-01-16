package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hackebrot/turtle"
)

// Writer writes Emojis to output
type Writer interface {
	WriteEmoji(emoji *turtle.Emoji) error
	WriteEmojis(emojis []*turtle.Emoji) error
}

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

// PlainWriter writes Emojis as Plain text (without any metadata)
type PlainWriter struct {
	w io.Writer
}

// NewPlainWriter creates a new PlainWriter
func NewPlainWriter(w io.Writer) (*PlainWriter, error) {
	return &PlainWriter{w: w}, nil
}

// WriteEmoji to an io.Writer as Plaintext
func (p *PlainWriter) WriteEmoji(emoji *turtle.Emoji) error {
	_, err := io.WriteString(p.w, emoji.Char)
	return err
}

// WriteEmojis to an io.Writer as Plaintext
func (p *PlainWriter) WriteEmojis(emojis []*turtle.Emoji) error {
	for i := 0; i < len(emojis); i++ {
		if err := p.WriteEmoji(emojis[i]); err != nil {
			return err
		}
	}
	return nil
}
