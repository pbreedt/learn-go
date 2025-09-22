package types

import "testing"

func TestConvertToString(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{"hello", "hello"},
		{42, "42"},
		{-27, "-27"},
		{3.1473, "3.15"}, // NOTE: Rounding to two decimal places results in rounding up
		{true, "unknown"},
	}

	for _, test := range tests {
		result := convertToString(test.input)
		if result != test.expected {
			t.Errorf("convertToString(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestCaseInsensitiveEquals(t *testing.T) {
	tests := []struct {
		a, b     string
		expected bool
	}{
		{"hello", "HELLO", true},
		{"GoLang", "golang", true},
		{"Test", "test1", false},
		{"Case", "case", true},
		{"Different", "DIFFERENT!", false},
	}

	for _, test := range tests {
		result := caseInsensitiveEquals(test.a, test.b)
		if result != test.expected {
			t.Errorf("caseInsensitiveEquals(%v, %v) = %v; want %v", test.a, test.b, result, test.expected)
		}
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		input         string
		expectedBytes int
		expectedRunes int
	}{
		{"hello", 5, 5},
		{"こんにちは", 15, 5}, // Each Japanese character is 3 bytes in UTF-8
		{"😊😊", 8, 2},     // Each emoji is 4 bytes in UTF-8
		{"Go语言", 8, 4},   // Mixed ASCII and Chinese characters
	}

	for _, test := range tests {
		resultBytes := lengthBytes(test.input)
		if resultBytes != test.expectedBytes {
			t.Errorf("lengthBytes(%v) = %v; want %v", test.input, resultBytes, test.expectedBytes)
		}

		resultRunes := lengthRunes(test.input)
		if resultRunes != test.expectedRunes {
			t.Errorf("lengthRunes(%v) = %v; want %v", test.input, resultRunes, test.expectedRunes)
		}
	}
}

func TestIterate(t *testing.T) {
	tests := []struct {
		input         string
		expectedBytes string
		expectedRunes string
	}{
		{"hello", "0:'h' 1:'e' 2:'l' 3:'l' 4:'o' ", "0:'h' 1:'e' 2:'l' 3:'l' 4:'o' "},
		{"こんにちは", `0:'ã' 1:'\u0081' 2:'\u0093' 3:'ã' 4:'\u0082' 5:'\u0093' 6:'ã' 7:'\u0081' 8:'«' 9:'ã' 10:'\u0081' 11:'¡' 12:'ã' 13:'\u0081' 14:'¯' `, "0:'こ' 3:'ん' 6:'に' 9:'ち' 12:'は' "},
		{"😊😊", `0:'ð' 1:'\u009f' 2:'\u0098' 3:'\u008a' 4:'ð' 5:'\u009f' 6:'\u0098' 7:'\u008a' `, "0:'😊' 4:'😊' "},
		{"Go语言", `0:'G' 1:'o' 2:'è' 3:'¯' 4:'\u00ad' 5:'è' 6:'¨' 7:'\u0080' `, "0:'G' 1:'o' 2:'语' 5:'言' "},
	}

	for _, test := range tests {
		resultBytes := iterateBytes(test.input)
		if resultBytes != test.expectedBytes {
			t.Errorf("iterateBytes(%v) = %v; want %v", test.input, resultBytes, test.expectedBytes)
		}

		resultRunes := iterateRunes(test.input)
		if resultRunes != test.expectedRunes {
			t.Errorf("iterateRunes(%v) = %v; want %v", test.input, resultRunes, test.expectedRunes)
		}
	}
}

func TestStrMap(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abc", "bcd"},
		{"xyz", "yz{"},
		{"Hello, World!", "Ifmmp-!Xpsme\""},
		{"Go语言", "Hp诮訁"},
	}

	for _, test := range tests {
		result := strMap(test.input, nextChar)
		if result != test.expected {
			t.Errorf("strMap(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
