package types

import (
	"fmt"
	"strings"
)

func convertToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%.2f", v)
	default:
		return "unknown"
	}
}

func caseInsensitiveEquals(a, b string) bool {
	return strings.EqualFold(a, b)
}

func lengthBytes(s string) int {
	return len(s)
}

func lengthRunes(s string) int {
	return len([]rune(s))
	// or
	// return utf8.RuneCountInString(s)
}

// Iterate over bytes could return unicode code points for non-ASCII characters
func iterateBytes(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		result += fmt.Sprintf("%d:%q ", i, s[i])
	}
	return result
}

// Iterate over runes returns actual characters
func iterateRunes(s string) string {
	result := ""

	for i, r := range s {
		result += fmt.Sprintf("%d:%q ", i, r)
	}
	return result
}

func nextChar(r rune) rune {
	return r + 1
}

func strMap(s string, f func(rune) rune) string {
	return strings.Map(f, s)
}
