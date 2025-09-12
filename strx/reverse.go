package strx

import "strings"

func Reverse(s string) string {
	if len(s) == 0 {
		return s
	}
	var sb strings.Builder

	srunes := []rune(s)
	sb.Grow(len(srunes))

	for i := len(srunes) - 1; i >= 0; i-- {
		sb.WriteRune(srunes[i])
	}
	return sb.String()
}
