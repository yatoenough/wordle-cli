package game

import (
	"fmt"

	"github.com/yatoenough/wordle-cli/internal/dictionary"
)

func Run() {
	dict := dictionary.MustLoadDictionary()

	fmt.Println(len(*dict.Words))
	fmt.Println(dict.GetRandomWord())
}
