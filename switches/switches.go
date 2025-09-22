package switches

import "fmt"

// ExactMatch example
func WithExpression(value string) string {
	switch value {
	case "something":
		return "it's something"
	case "this", "that":
		return "it's this or that"
	default:
		return "it's probably nothing"
	}
}

// ExactMatch example
func WithExtendedExpression(value string) string {
	if len(value) < 3 {
		return "got short"
	}

	switch start := value[0:4]; start {
	case "some":
		return "got some"
	default:
		return "got something else"
	}
}

// Range example
func WithoutExpression(i int) string {
	switch {
	case i > 10:
		return "Ohh, big!"
	case i < 10:
		return "Mmm, small."
	default:
		return "Just right."
	}
}

// Fallthrough example
func WithFallthroughMultipleCases(i int) string {
	result := ""
	switch {
	case i < 5:
		result = "Less than five\n"
		fallthrough
	case i < 8:
		result += "Less than eight\n"
		fallthrough
	case i < 10:
		result += "Less than ten\n"
	default:
		result += "Greater than or equal to ten\n"
	}
	return result
}

func WithBreak(value string) string {
	result := ""

Loop:
	for _, ch := range value {
		switch ch {
		case ' ': // skip/ignore space
			break
		case '\n': // break at newline
			break Loop
		default:
			result += fmt.Sprintf("%c\n", ch)
		}
	}

	return result
}
