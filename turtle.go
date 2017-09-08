package turtle

// Emojis maps a name to an Emoji
var Emojis = make(map[string]*Emoji)

func init() {
	for _, e := range emojis {
		Emojis[e.Name] = e
	}
}
