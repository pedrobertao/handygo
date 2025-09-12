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
		if strings.Contains(s, "_") {
			ssplit := strings.Split(s, "_")
			isScreaming := true
			for _, subs := range ssplit {
				if subs != strings.ToLower(subs) {
					isScreaming = false
					break
				}
			}
			if isScreaming {
				return s
			} else {
				sjoined := strings.Join(ssplit, "_")
				return strings.ToLower(sjoined)
			}
		}
		var sb strings.Builder
		sb.Grow(len(s))

		for i, r := range strFields[0] {
			if i == 0 {
				sb.WriteRune(unicode.ToLower(r))
				continue
			}
			if unicode.IsUpper(r) {
				sb.WriteRune('_')
			}
			sb.WriteRune(unicode.ToLower(r))
		}
		return sb.String()
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
		if strings.Contains(s, "_") {
			ssplit := strings.Split(s, "_")
			isScreaming := true
			for _, subs := range ssplit {
				if subs != strings.ToUpper(subs) {
					isScreaming = false
					break
				}
			}
			if isScreaming {
				return s
			} else {
				sjoined := strings.Join(ssplit, "_")
				return strings.ToUpper(sjoined)
			}
		}

		var sb strings.Builder
		sb.Grow(len(s))

		for i, r := range strFields[0] {
			if i == 0 {
				sb.WriteRune(unicode.ToUpper(r))
				continue
			}
			if unicode.IsUpper(r) {
				sb.WriteRune('_')
			}
			sb.WriteRune(unicode.ToUpper(r))
		}
		return sb.String()
	}

	for i, r := range strFields {
		strFields[i] = strings.ToUpper(r)
	}
	return strings.Join(strFields, "_")
}

func ToCamel(s string) string {
	if len(s) == 0 {
		return s
	}

	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}

	if strings.Contains(s, "_") {
		parts := strings.Split(s, "_")
		var sb strings.Builder
		for i, part := range parts {
			if len(part) == 0 {
				continue
			}
			if i == 0 {
				sb.WriteString(strings.ToLower(part))
			} else {
				sb.WriteString(strings.Title(part))
			}
		}
		return sb.String()
	}

	if strings.Contains(s, "-") {
		parts := strings.Split(s, "-")
		var sb strings.Builder
		for i, part := range parts {
			if len(part) == 0 {
				continue
			}
			if i == 0 {
				sb.WriteString(strings.ToLower(part))
			} else {
				sb.WriteString(strings.Title(part))
			}
		}
		return sb.String()
	}

	strFields := strings.Fields(s)
	if len(strFields) == 0 {
		return s
	}

	if len(strFields) == 1 {
		runes := []rune(s)
		if len(runes) > 0 {
			runes[0] = unicode.ToLower(runes[0])
		}
		return string(runes)
	}

	var sb strings.Builder
	for i, field := range strFields {
		if i == 0 {
			sb.WriteString(strings.ToLower(field))
		} else {
			sb.WriteString(strings.Title(field))
		}
	}
	return sb.String()
}

func ToKebab(s string) string {
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
				ss = append(ss, '-')
			}
			ss = append(ss, unicode.ToLower(r))
		}
		return string(ss)
	}

	for i, r := range strFields {
		strFields[i] = strings.ToLower(r)
	}
	return strings.Join(strFields, "-")
}
