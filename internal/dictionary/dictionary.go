package dictionary

import (
	_ "embed"
	"strings"

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

func parseDictionary(dict string) []string {
	parsed := strings.Split(dict, "\n")

	return parsed
}
