package converters

import "testing"

func TestCamelToSnake(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"CamelCase", "camel_case"},
		{"camelCase", "camel_case"},
		{"", ""},
		{"Already_Snake", "already_snake"},
	}

	for _, test := range tests {
		result := ConvertCamelToSnake(test.input)
		if result != test.expected {
			t.Errorf("ConvertCamelToSnake(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
