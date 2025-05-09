package converters

import "testing"

func TestSnakeToCamel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"snake_case", "snakeCase"},
		{"another_example", "anotherExample"},
		{"test_case_here", "testCaseHere"},
	}

	for _, test := range tests {
		result := ConvertSnakeToCamel(test.input)
		if result != test.expected {
			t.Errorf("ConvertSnakeToCamel(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
