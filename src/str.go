package src

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
