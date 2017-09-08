package turtle

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hackebrot/go-repr/repr"
)

func Test_filter(t *testing.T) {
	emojis := []*Emoji{
		{
			Name:     "turtle",
			Category: "animals_and_nature",
			Char:     "🐢",
			Keywords: []string{"animal", "slow", "nature", "tortoise"},
		},
		{
			Name:     "coffee",
			Category: "food_and_drink",
			Char:     "☕",
			Keywords: []string{"beverage", "caffeine", "latte", "espresso"},
		},
		{
			Name:     "woman_technologist",
			Category: "people",
			Char:     "👩‍💻",
			Keywords: []string{"coder", "developer", "engineer", "programmer", "software", "woman", "human"},
		},
		{
			Name:     "dog",
			Category: "animals_and_nature",
			Char:     "🐶",
			Keywords: []string{"animal", "friend", "nature", "woof", "puppy", "pet", "faithful"},
		},
	}

	tests := []struct {
		name string
		f    func(e *Emoji) bool
		want []*Emoji
	}{
		{
			name: "no match",
			f: func(e *Emoji) bool {
				return e.Name == "nope"
			},
			want: nil,
		},
		{
			name: "single match",
			f: func(e *Emoji) bool {
				for _, keyword := range e.Keywords {
					if keyword == "developer" {
						return true
					}
				}
				return false
			},
			want: []*Emoji{
				{
					Name:     "woman_technologist",
					Category: "people",
					Char:     "👩‍💻",
					Keywords: []string{"coder", "developer", "engineer", "programmer", "software", "woman", "human"},
				},
			},
		},
		{
			name: "multiple matches",
			f: func(e *Emoji) bool {
				return e.Category == "animals_and_nature"
			},
			want: []*Emoji{
				{
					Name:     "turtle",
					Category: "animals_and_nature",
					Char:     "🐢",
					Keywords: []string{"animal", "slow", "nature", "tortoise"},
				},
				{
					Name:     "dog",
					Category: "animals_and_nature",
					Char:     "🐶",
					Keywords: []string{"animal", "friend", "nature", "woof", "puppy", "pet", "faithful"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filter(emojis, tt.f); !cmp.Equal(got, tt.want) {
				t.Errorf("filter() = %v, want %v", repr.Repr(got), repr.Repr(tt.want))
			}
		})
	}
}
