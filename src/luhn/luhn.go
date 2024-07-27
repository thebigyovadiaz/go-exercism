package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	if len(id) <= 1 {
		return false
	}

	newId := strings.ReplaceAll(id, " ", "")

	if len(newId) <= 1 {
		return false
	}

	sum := 0
	isSecond := false

	for i := len(newId) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(newId[i]))
		if err != nil {
			return false
		}

		if isSecond {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		isSecond = !isSecond
	}

	return sum%10 == 0
}
