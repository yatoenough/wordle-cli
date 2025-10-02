package main

import (
	"github.com/yatoenough/wordle-cli/internal/dictionary"
	"github.com/yatoenough/wordle-cli/internal/game"
)

func main() {
	dict := dictionary.MustLoadDictionary()

	wordToGuess := dict.GetRandomWord()

	game := game.NewWordleGame(wordToGuess, dict, 6)
	game.Run()
}
