package converters

import "testing"

func TestConvertKebabToCamel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"kebab-case", "kebabCase"},
		{"another-example", "anotherExample"},
		{"test-case-here", "testCaseHere"},
	}

	for _, test := range tests {
		result := ConvertKebabToCamel(test.input)
		if result != test.expected {
			t.Errorf("ConvertKebabToCamel(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
