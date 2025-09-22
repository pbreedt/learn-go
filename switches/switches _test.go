package switches

import (
	"testing"
)

func TestWithExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"something", "it's something"},
		{"this", "it's this or that"},
		{"that", "it's this or that"},
		{"other", "it's probably nothing"},
	}

	for _, test := range tests {
		result := WithExpression(test.input)
		if result != test.expected {
			t.Errorf("WithExpression(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestWithExtendedExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"so", "got short"},
		{"somevalue", "got some"},
		{"other", "got something else"},
	}

	for _, test := range tests {
		result := WithExtendedExpression(test.input)
		if result != test.expected {
			t.Errorf("WithExtendedExpression(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestWithoutExpression(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{5, "Mmm, small."},
		{10, "Just right."},
		{15, "Ohh, big!"},
	}

	for _, test := range tests {
		result := WithoutExpression(test.input)
		if result != test.expected {
			t.Errorf("WithoutExpression(%d) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestWithFallthroughMultipleCases(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{2, "Less than five\nLess than eight\nLess than ten\n"},
		{5, "Less than eight\nLess than ten\n"},
		{6, "Less than eight\nLess than ten\n"},
		{8, "Less than ten\n"},
		{9, "Less than ten\n"},
		{10, "Greater than or equal to ten\n"},
		{11, "Greater than or equal to ten\n"},
	}

	for _, test := range tests {
		result := WithFallthroughMultipleCases(test.input)
		if result != test.expected {
			t.Errorf("WithFallthroughMultipleCases(%d) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestWithBreak(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a b\nc", "a\nb\n"},
		{"hello world", "h\ne\nl\nl\no\nw\no\nr\nl\nd\n"},
		{"go\nlang", "g\no\n"},
	}
	for _, test := range tests {
		result := WithBreak(test.input)
		if result != test.expected {
			t.Errorf("WithBreak(%s) = %q; want %q", test.input, result, test.expected)
		}
	}

}
