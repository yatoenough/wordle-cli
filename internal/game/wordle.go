package game

import (
	"fmt"
	"log"
	"strings"

	"github.com/yatoenough/wordle-cli/internal/dictionary"
)

type WordleGame struct {
	wordToGuess string
	userGuess   string
	dict        *dictionary.Dictionary
}

func NewWordleGame(wordToGuess string, dict *dictionary.Dictionary) *WordleGame {
	return &WordleGame{
		wordToGuess: wordToGuess,
		dict:        dict,
		userGuess:   "",
	}
}

func (g *WordleGame) Run() {
	g.userGuess = getUserInput()

	if len(g.userGuess) != 5 {
		fmt.Println("Word must be 5 chars long!")
		g.Run()
	}

	result := g.compare()

	if !g.parseResult(result) {
		g.Run()
	}
}

func getUserInput() string {
	var input string

	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(input)
}

func (g *WordleGame) compare() []byte {
	result := make([]byte, 0, 5)

	for i := range 5 {
		if strings.Contains(g.wordToGuess, string(g.userGuess[i])) {
			if g.wordToGuess[i] == g.userGuess[i] {
				result = append(result, 'g')
			} else {
				result = append(result, 'y')
			}
			continue
		}

		result = append(result, 'x')
	}

	return result
}

func (g *WordleGame) parseResult(result []byte) bool {
	greenCount := 0

	for _, char := range result {
		if char == 'g' {
			greenCount++
		}

		fmt.Print(string(char))
	}

	fmt.Println()

	if greenCount == 5 {
		fmt.Println("You guessed!")
		return true
	}

	return false
}
