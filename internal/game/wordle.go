package game

import (
	"fmt"

	"github.com/yatoenough/wordle-cli/internal/dictionary"
	"github.com/yatoenough/wordle-cli/internal/io"
)

type WordleGame struct {
	wordToGuess string
	userGuess   string
	dict        *dictionary.Dictionary
	guesses     []string
	attempts    int
	maxAttempts int
}

func NewWordleGame(wordToGuess string, dict *dictionary.Dictionary, maxAttempts int) *WordleGame {
	return &WordleGame{
		wordToGuess: wordToGuess,
		dict:        dict,
		userGuess:   "",
		guesses:     make([]string, 0, 6),
		attempts:    0,
		maxAttempts: maxAttempts,
	}
}

func (g *WordleGame) Run() {
	g.runWithError(Ok)
}

func (g *WordleGame) runWithError(errorMsg string) {
	g.attempts++

	io.ClearScreen()

	for _, guess := range g.guesses {
		fmt.Println(guess)
	}

	if errorMsg != Ok {
		fmt.Println(errorMsg)
	}

	fmt.Printf("Attempt %d/%d: ", g.attempts, g.maxAttempts)

	g.userGuess = io.GetUserInput()

	if len(g.userGuess) != 5 {
		g.attempts--
		g.runWithError(IsNotFiveCharsLongErr)
		return
	}

	if !g.dict.Contains(g.userGuess) {
		g.attempts--
		g.runWithError(NotInDictionaryErr)
		return
	}

	result := g.compare()

	if g.parseResult(result) {
		return
	}

	if g.attempts >= g.maxAttempts {
		fmt.Printf("Game over! The word was: %s\n", g.wordToGuess)
		return
	}

	g.runWithError(Ok)
}

func (g *WordleGame) compare() []byte {
	result := make([]byte, 5)
	letterCount := make(map[rune]int)

	for _, char := range g.wordToGuess {
		letterCount[rune(char)]++
	}

	for i := range 5 {
		if g.userGuess[i] == g.wordToGuess[i] {
			result[i] = 'g'
			letterCount[rune(g.userGuess[i])]--
		}
	}

	for i := range 5 {
		if result[i] == 'g' {
			continue
		}
		char := rune(g.userGuess[i])
		if count, exists := letterCount[char]; exists && count > 0 {
			result[i] = 'y'
			letterCount[char]--
		} else {
			result[i] = 'x'
		}
	}

	return result
}

func (g *WordleGame) parseResult(result []byte) bool {
	greenCount := 0

	parsedWord := ""

	for i, char := range result {
		letter := string(g.userGuess[i])
		switch char {
		case 'g':
			parsedWord += green + letter + reset
			greenCount++
		case 'y':
			parsedWord += yellow + letter + reset
		case 'x':
			parsedWord += gray + letter + reset
		}
	}

	g.guesses = append(g.guesses, parsedWord)

	fmt.Println()

	if greenCount == 5 {
		fmt.Println(green + "You guessed!" + reset)
		return true
	}

	return false
}
