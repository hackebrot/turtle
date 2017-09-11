package main

import (
	"bytes"
	"testing"

	"github.com/hackebrot/turtle"
)

func Test_writeEmoji(t *testing.T) {
	emoji := &turtle.Emoji{
		Name:     "smile",
		Category: "people",
		Char:     "😄",
		Keywords: []string{"face", "happy", "joy", "funny", "haha", "laugh", "like", ":D", ":)"},
	}

	type args struct {
		f        string
		e        *turtle.Emoji
		name     bool
		char     bool
		category bool
		keywords bool
	}

	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "name",
			args: args{
				f:    "%v",
				e:    emoji,
				name: true,
			},
			wantW: `Name: "smile"`,
		},
		{
			name: "category",
			args: args{
				f:        "%v",
				e:        emoji,
				category: true,
			},
			wantW: `Category: "people"`,
		},
		{
			name: "keywords",
			args: args{
				f:        "%v",
				e:        emoji,
				keywords: true,
			},
			wantW: `Keywords: ["face" "happy" "joy" "funny" "haha" "laugh" "like" ":D" ":)"]`,
		},
		{
			name: "char",
			args: args{
				f:    "%v",
				e:    emoji,
				char: true,
			},
			wantW: `Char: 😄`,
		},
		{
			name: "multiple attributes",
			args: args{
				f:        "%v",
				e:        emoji,
				name:     true,
				category: true,
			},
			wantW: "Name: \"smile\"\nCategory: \"people\"",
		},
		{
			name: "format",
			args: args{
				f:    "# %v #",
				e:    emoji,
				char: true,
			},
			wantW: `# Char: 😄 #`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			writeEmoji(w, tt.args.f, tt.args.e, tt.args.name, tt.args.char, tt.args.category, tt.args.keywords)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("writeEmoji() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
