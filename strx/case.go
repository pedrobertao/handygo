package strx

import (
	"strings"
	"unicode"
)

func ToSnake(s string) string {
	if len(s) == 0 {
		return s
	}
	strFields := strings.Fields(s)
	lenFields := len(strFields)

	if lenFields == 0 {
		return s
	}

	if lenFields == 1 {
		ss := make([]rune, 0)
		for i, r := range strFields[0] {
			if i == 0 {
				ss = append(ss, unicode.ToLower(r))
				continue
			}
			if unicode.IsUpper(r) {
				ss = append(ss, '_')
			}
			ss = append(ss, unicode.ToLower(r))
		}
		return string(ss)
	}

	for i, r := range strFields {
		strFields[i] = strings.ToLower(r)
	}
	return strings.Join(strFields, "_")
}

func ToScreamingSnakeCase(s string) string {
	if len(s) == 0 {
		return s
	}
	strFields := strings.Fields(s)
	lenFields := len(strFields)

	if lenFields == 0 {
		return s
	}

	if lenFields == 1 {
		ss := make([]rune, 0)
		for i, r := range strFields[0] {
			if i == 0 {
				ss = append(ss, unicode.ToUpper(r))
				continue
			}
			if unicode.IsUpper(r) {
				ss = append(ss, '_')
			}
			ss = append(ss, unicode.ToUpper(r))
		}
		return string(ss)
	}

	for i, r := range strFields {
		strFields[i] = strings.ToLower(r)
	}
	return strings.Join(strFields, "_")
}

func ToCamel() string {
	return ""
}
func ToKebab() string {
	return ""
}

func Reverse() (string, error) {
	return "", nil
}
