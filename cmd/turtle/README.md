# turtle

CLI for the [turtle library][library] 😄 🐢 💻

## Installation

``$ go get github.com/hackebrot/turtle/cmd/turtle``

## Usage

```text
Print the emoji for a given name identifier

Usage:
  turtle [flags]
  turtle [command]

Available Commands:
  category    Print all emojis of the category
  help        Help about any command
  keyword     Print all emojis with the keyword
  search      Print emojis with a name that contains the search string
  version     Print the version number of turtle

Flags:
  -h, --help            help for turtle
  -i, --indent string   indent for JSON output
  -p, --prefix string   prefix for JSON output

Use "turtle [command] --help" for more information about a command.
```

### Emoji lookup

```bash
$ turtle -i "  " rocket
```

```json
{
  "name": "rocket",
  "category": "travel_and_places",
  "char": "🚀",
  "keywords": [
    "launch",
    "ship",
    "staffmode",
    "NASA",
    "outer space",
    "outer_space",
    "fly"
  ]
}
```

### Search

```bash
$ turtle -i "  " search computer
```

```json
[
  {
    "name": "computer",
    "category": "objects",
    "char": "💻",
    "keywords": [
      "technology",
      "laptop",
      "screen",
      "display",
      "monitor"
    ]
  },
  {
    "name": "computer_mouse",
    "category": "objects",
    "char": "🖱",
    "keywords": [
      "click"
    ]
  },
  {
    "name": "desktop_computer",
    "category": "objects",
    "char": "🖥",
    "keywords": [
      "technology",
      "computing",
      "screen"
    ]
  }
]
```

### Category

```bash
$ turtle -i "  " category travel_and_places
```

```json
[
  {
    "name": "aerial_tramway",
    "category": "travel_and_places",
    "char": "🚡",
    "keywords": [
      "transportation",
      "vehicle",
      "ski"
    ]
  },
  {
    "name": "airplane",
    "category": "travel_and_places",
    "char": "✈️",
    "keywords": [
      "vehicle",
      "transportation",
      "flight",
      "fly"
    ]
  },
  {
    "name": "ambulance",
    "category": "travel_and_places",
    "char": "🚑",
    "keywords": [
      "health",
      "911",
      "hospital"
    ]
  }
]
```

### Keyword

```bash
$ turtle -i "  " keyword happy
```

```json
[
  {
    "name": "blush",
    "category": "people",
    "char": "😊",
    "keywords": [
      "face",
      "smile",
      "happy",
      "flushed",
      "crush",
      "embarrassed",
      "shy",
      "joy"
    ]
  },
  {
    "name": "grin",
    "category": "people",
    "char": "😁",
    "keywords": [
      "face",
      "happy",
      "smile",
      "joy",
      "kawaii"
    ]
  },
  {
    "name": "grinning",
    "category": "people",
    "char": "😀",
    "keywords": [
      "face",
      "smile",
      "happy",
      "joy",
      ":D",
      "grin"
    ]
  }
]
```

[library]: ../../README.md
