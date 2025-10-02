package game

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
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
	guesses     []string
}

func NewWordleGame(wordToGuess string, dict *dictionary.Dictionary) *WordleGame {
	return &WordleGame{
		wordToGuess: wordToGuess,
		dict:        dict,
		userGuess:   "",
		attempts:    0,
		maxAttempts: 6,
		guesses:     make([]string, 0, 6),
	}
}

func (g *WordleGame) Run() {
	g.runWithError("")
}

func (g *WordleGame) runWithError(errorMsg string) {
	g.attempts++

	clearScreen()

	for _, guess := range g.guesses {
		fmt.Println(guess)
	}

	if errorMsg != "" {
		fmt.Println(errorMsg)
	}

	fmt.Printf("Attempt %d/%d: ", g.attempts, g.maxAttempts)

	g.userGuess = getUserInput()

	if len(g.userGuess) != 5 {
		g.attempts--
		g.runWithError("Word must be 5 chars long!")
		return
	}

	if !g.dict.Contains(g.userGuess) {
		g.attempts--
		g.runWithError("Word is not in dictionary!")
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

	g.runWithError("")
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

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
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
