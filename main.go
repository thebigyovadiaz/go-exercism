package main

import (
	"fmt"
	scrabble "github.com/thebigyovadiaz/go-exercism/src/scrabble-score"
)

func main() {
	result := scrabble.ScrabbleScore("cabbage")
	fmt.Println("result", result)
}
