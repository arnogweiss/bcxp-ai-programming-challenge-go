package converters

import "testing"

func TestConvertKebabToSnake(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"kebab-case", "kebab_case"},
		{"another-example", "another_example"},
		{"test-case-here", "test_case_here"},
	}

	for _, test := range tests {
		result := ConvertKebabToSnake(test.input)
		if result != test.expected {
			t.Errorf("ConvertKebabToSnake(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
