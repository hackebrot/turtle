package turtle

// Emojis maps a name to an Emoji
var Emojis = make(map[string]*Emoji)

func init() {
	for _, e := range emojis {
		Emojis[e.Name] = e
	}
}

// filter a given slice of Emoji by f
func filter(emojis []*Emoji, f func(e *Emoji) bool) []*Emoji {
	var r []*Emoji
	for _, e := range emojis {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}
