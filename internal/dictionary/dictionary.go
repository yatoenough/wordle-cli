package dictionary

import (
	_ "embed"
	"math/rand/v2"
	"strings"

	wordle_embed "github.com/yatoenough/wordle-cli"
)

type Dictionary struct {
	Words []string
}

func MustLoadDictionary() *Dictionary {
	dict := wordle_embed.Dictionary

	words := parseDictionary(dict)

	return &Dictionary{
		Words: words,
	}
}

func (d *Dictionary) GetRandomWord() string {
	i := rand.IntN(len(d.Words))
	return d.Words[i]
}

func (d *Dictionary) Contains(word string) bool {
	word = strings.ToLower(strings.TrimSpace(word))
	for _, w := range d.Words {
		if strings.ToLower(w) == word {
			return true
		}
	}
	return false
}

func parseDictionary(dict string) []string {
	lines := strings.Split(dict, "\n")
	parsed := make([]string, 0, len(lines))

	for _, line := range lines {
		word := strings.TrimSpace(line)
		if word != "" {
			parsed = append(parsed, word)
		}
	}

	return parsed
}
