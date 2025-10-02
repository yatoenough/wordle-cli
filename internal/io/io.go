package io

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetUserInput() string {
	var input string

	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Println("Input cannot be empty!")
		return GetUserInput()
	}

	return strings.ToLower(input)
}

func ClearScreen() {
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
