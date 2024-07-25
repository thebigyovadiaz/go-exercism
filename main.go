package main

import (
	"fmt"
	"github.com/thebigyovadiaz/go-exercism/src/leap"
	scrabble "github.com/thebigyovadiaz/go-exercism/src/scrabble-score"
)

func main() {
	// Scrabble score exercise
	resultScrabble := scrabble.ScrabbleScore("cabbage")
	fmt.Println("resultScrabble:", resultScrabble)

	// Leap year exercise
	leapYear := leap.IsLeapYear(1996)
	fmt.Println("leap year:", leapYear)
}
