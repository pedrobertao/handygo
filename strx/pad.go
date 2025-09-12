package strx

func PadLeft(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}
	padLength := length - len(s)
	padding := make([]rune, padLength)
	for i := range padding {
		padding[i] = padChar
	}
	return string(padding) + s
}

func PadRight(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}
	padLength := length - len(s)
	padding := make([]rune, padLength)
	for i := range padding {
		padding[i] = padChar
	}
	return s + string(padding)
}
