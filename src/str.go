package src

import "unicode"

func ReverseString(text *string) string {
	length := len(*text)

	if length <= 1 {
		return *text
	}

	runes := make([]rune, length)

	for _, element := range *text {
		length--
		runes[length] = element
	}

	return string(runes)
}

func CountSymbols(text *string) (count int) {
	count = 0

	for _, element := range *text {
		if !unicode.IsLetter(element) && !unicode.IsNumber(element) {
			count++
		}
	}

	return count
}