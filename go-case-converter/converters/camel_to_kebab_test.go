package converters

import "testing"

func TestConvertCamelToKebab(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"camelCase", "camel-case"},
		{"CamelCase", "camel-case"},
		{"testCaseHere", "test-case-here"},
	}

	for _, test := range tests {
		result := ConvertCamelToKebab(test.input)
		if result != test.expected {
			t.Errorf("ConvertCamelToKebab(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
