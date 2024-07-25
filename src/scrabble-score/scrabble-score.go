package scrabble_score

import (
	"fmt"
	"strings"
)

func letterValue(letter string) int {
	switch strings.ToUpper(letter) {
	case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
		return 1
	case "D", "G":
		return 2
	case "B", "C", "M", "P":
		return 3
	case "F", "H", "V", "W", "Y":
		return 4
	case "K":
		return 5
	case "J", "X":
		return 8
	case "Q", "Z":
		return 10
	default:
		return 0
	}
}

func ScrabbleScore(word string) int {
	result := 0

	for _, v := range word {
		letter := fmt.Sprintf("%c", v)
		value := letterValue(letter)
		result += value
	}

	return result
}
