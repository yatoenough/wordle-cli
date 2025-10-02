package game

import (
	"fmt"
	"log"
	"strings"

	"github.com/yatoenough/wordle-cli/internal/dictionary"
)

const (
	green  = "\033[32m"
	yellow = "\033[33m"
	gray   = "\033[90m"
	reset  = "\033[0m"
)

type WordleGame struct {
	wordToGuess string
	userGuess   string
	dict        *dictionary.Dictionary
	attempts    int
	maxAttempts int
}

func NewWordleGame(wordToGuess string, dict *dictionary.Dictionary) *WordleGame {
	return &WordleGame{
		wordToGuess: wordToGuess,
		dict:        dict,
		userGuess:   "",
		attempts:    0,
		maxAttempts: 6,
	}
}

func (g *WordleGame) Run() {
	g.attempts++
	fmt.Printf("Attempt %d/%d: ", g.attempts, g.maxAttempts)

	g.userGuess = getUserInput()

	if len(g.userGuess) != 5 {
		fmt.Println("Word must be 5 chars long!")
		g.attempts--
		g.Run()
		return
	}

	if !g.dict.Contains(g.userGuess) {
		fmt.Println("Word not in dictionary!")
		g.attempts--
		g.Run()
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

	g.Run()
}

func getUserInput() string {
	var input string

	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Println("Input cannot be empty!")
		return getUserInput()
	}

	return strings.ToLower(input)
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

	for i, char := range result {
		letter := string(g.userGuess[i])
		switch char {
		case 'g':
			fmt.Print(green + letter + reset)
			greenCount++
		case 'y':
			fmt.Print(yellow + letter + reset)
		case 'x':
			fmt.Print(gray + letter + reset)
		}
	}

	fmt.Println()

	if greenCount == 5 {
		fmt.Println(green + "You guessed!" + reset)
		return true
	}

	return false
}
