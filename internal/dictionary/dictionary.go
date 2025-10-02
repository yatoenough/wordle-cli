package dictionary

import (
	_ "embed"
	"math/rand"
	"strings"
	"time"

	wordle_embed "github.com/yatoenough/wordle-cli"
)

type Dictionary struct {
	Words *[]string
}

func MustLoadDictionary() *Dictionary {
	dict := wordle_embed.Dictionary

	words := parseDictionary(dict)

	return &Dictionary{
		Words: &words,
	}
}

func (d *Dictionary) GetRandomWord() string {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	i := r.Intn(len(*d.Words))

	words := *d.Words

	return words[i]
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
