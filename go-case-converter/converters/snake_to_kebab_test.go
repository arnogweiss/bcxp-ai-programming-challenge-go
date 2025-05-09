package converters

import "testing"

func TestConvertSnakeToKebab(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"snake_case", "snake-case"},
		{"another_example", "another-example"},
		{"test_case_here", "test-case-here"},
	}

	for _, test := range tests {
		result := ConvertSnakeToKebab(test.input)
		if result != test.expected {
			t.Errorf("ConvertSnakeToKebab(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
