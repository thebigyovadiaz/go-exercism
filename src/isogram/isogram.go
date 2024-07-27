package isogram

import (
	"fmt"
	"strings"
)

func IsIsogram(word string) bool {
	if len(word) == 0 {
		return true
	}

	word = strings.ReplaceAll(word, "-", "")
	word = strings.ReplaceAll(word, " ", "")

	for k, v := range word {
		newWord := word

		if k == 0 {
			newWord = word[k+1:]
		} else if k == len(word)-1 {
			newWord = word[:k]
		} else {
			newWord = word[:k] + word[k+1:]
		}

		for _, b := range newWord {
			vU := strings.ToUpper(fmt.Sprintf("%c", v))
			bU := strings.ToUpper(fmt.Sprintf("%c", b))

			if vU == bU {
				return false
			}
		}
	}

	return true
}
