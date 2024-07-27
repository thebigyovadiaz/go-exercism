package main

import (
	"fmt"
	difference_squares "github.com/thebigyovadiaz/go-exercism/src/difference-squares"
	"github.com/thebigyovadiaz/go-exercism/src/isogram"
	"github.com/thebigyovadiaz/go-exercism/src/leap"
	"github.com/thebigyovadiaz/go-exercism/src/luhn"
	scrabble "github.com/thebigyovadiaz/go-exercism/src/scrabble-score"
)

func main() {
	// Scrabble score exercise
	resultScrabble := scrabble.ScrabbleScore("cabbage")
	fmt.Println("resultScrabble:", resultScrabble)

	// Leap year exercise
	leapYear := leap.IsLeapYear(1996)
	fmt.Println("leap year:", leapYear)

	// Isogram exercise
	isogram := isogram.IsIsogram("issogram")
	fmt.Println("isogram:", isogram)

	// Difference squares exercise
	diffSum := difference_squares.Difference(5)
	fmt.Println("diffSum:", diffSum)

	// Luhn exercise
	isValid := luhn.Valid("055-444-285")
	fmt.Println("isValid:", isValid)
}
