package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Hangman struct {
	word     string
	guesses  []string
	maxGuess int
}

func (h *Hangman) Init(words []string, maxGuess int) {
	rand.Seed(42)
	h.word = words[rand.Intn(len(words))]
	h.guesses = make([]string, 0)
	h.maxGuess = maxGuess
}

func (h *Hangman) Display() {
	fmt.Println(strings.Repeat("-", len(h.word)))
	for _, letter := range h.word {
		if contains(h.guesses, string(letter)) {
			fmt.Printf("%c ", letter)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
	fmt.Printf("Guesses remaining: %d\n", h.maxGuess-len(h.guesses))
}

func (h *Hangman) Guess(letter string) {
	h.guesses = append(h.guesses, letter)
}

func (h *Hangman) Won() bool {
	for _, letter := range h.word {
		if !contains(h.guesses, string(letter)) {
			return false
		}
	}
	return true
}

func (h *Hangman) Lost() bool {
	return len(h.guesses) >= h.maxGuess
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	// initialize game
	words := []string{"apple", "banana", "cherry", "date", "elderberry"}
	maxGuess := 6
	hangman := Hangman{}
	hangman.Init(words, maxGuess)

	// game loop
	for !hangman.Won() && !hangman.Lost() {
		hangman.Display()
		var letter string
		fmt.Print("Guess a letter: ")
		fmt.Scanln(&letter)
		hangman.Guess(letter)
	}

	// display final result
	hangman.Display()
	if hangman.Won() {
		fmt.Println("Congratulations, you won!")
	} else {
		fmt.Printf("Sorry, the word was %s. Better luck next time!\n", hangman.word)
	}
}
