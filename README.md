# Wordle Game CLI

A command-line Wordle game written in Go

## Installation

```bash
git clone https://github.com/yourusername/wordle-cli.git
cd wordle-cli
```

## Usage

### Make

```bash
make run
```

or

```bash
make build
./dist/wordle
```

### Go

```bash
go run cmd/wordle/main.go
```

## How to Play

- Guess a 5-letter word
- 🟩 Green = correct letter, correct position
- 🟨 Yellow = correct letter, wrong position
- ⬜ Gray = letter not in word
- Win in 6 tries or less
