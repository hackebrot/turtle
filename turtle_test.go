package turtle

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hackebrot/go-repr/repr"
)

var testEmojis = []*Emoji{
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

func Test_category(t *testing.T) {
	tests := []struct {
		name string
		c    string
		want []*Emoji
	}{
		{
			name: "no matches",
			c:    "activity",
			want: nil,
		},
		{
			name: "multiple matches",
			c:    "animals_and_nature",
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
			if got := category(testEmojis, tt.c); !cmp.Equal(got, tt.want) {
				t.Errorf("category() = %v, want %v", repr.Repr(got), repr.Repr(tt.want))
			}
		})
	}
}

func Test_keyword(t *testing.T) {
	tests := []struct {
		name string
		k    string
		want []*Emoji
	}{
		{
			name: "no matches",
			k:    "weather",
			want: nil,
		},
		{
			name: "single match",
			k:    "programmer",
			want: []*Emoji{
				{
					Name:     "woman_technologist",
					Category: "people",
					Char:     "👩‍💻",
					Keywords: []string{"coder", "developer", "engineer", "programmer", "software", "woman", "human"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keyword(testEmojis, tt.k); !cmp.Equal(got, tt.want) {
				t.Errorf("keyword() = %v, want %v", repr.Repr(got), repr.Repr(tt.want))
			}
		})
	}
}

func Test_search(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []*Emoji
	}{
		{
			name: "no matches",
			s:    "nope",
			want: nil,
		},
		{
			name: "substring",
			s:    "technologist",
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
			name: "full string",
			s:    "woman_technologist",
			want: []*Emoji{
				{
					Name:     "woman_technologist",
					Category: "people",
					Char:     "👩‍💻",
					Keywords: []string{"coder", "developer", "engineer", "programmer", "software", "woman", "human"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(testEmojis, tt.s); !cmp.Equal(got, tt.want) {
				t.Errorf("search() = %v, want %v", repr.Repr(got), repr.Repr(tt.want))
			}
		})
	}
}

func Test_Emojis(t *testing.T) {
	want := &Emoji{
		Name:     "fox_face",
		Category: "animals_and_nature",
		Char:     "🦊",
		Keywords: []string{"animal", "nature", "face"},
	}

	got, ok := Emojis[want.Name]

	if !ok {
		t.Fatalf("Emojis does not contain name %v", repr.Repr(want.Name))
	}

	if !cmp.Equal(got, want) {
		t.Errorf("Emojis[] = %v, want %v", repr.Repr(got), repr.Repr(want))
	}
}

func Test_EmojisByChar(t *testing.T) {
	want := &Emoji{
		Name:     "fox_face",
		Category: "animals_and_nature",
		Char:     "🦊",
		Keywords: []string{"animal", "nature", "face"},
	}

	got, ok := EmojisByChar[want.Char]

	if !ok {
		t.Fatalf("Emojis does not contain char %v", repr.Repr(want.Char))
	}

	if !cmp.Equal(got, want) {
		t.Errorf("Emojis[] = %v, want %v", repr.Repr(got), repr.Repr(want))
	}
}
