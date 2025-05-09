package validators

import "testing"

func TestValidateSingleWord(t *testing.T) {
	tests := []struct {
		input    string
		expected string // Expected error message, or empty string if valid
	}{
		{"validInput", ""},
		{"valid_input", ""},
		{"Valid123", ""},
		{"123", ""},
		{"", "input must be a single word containing only letters, numbers, or underscores"},
		{"invalid input", "input must be a single word containing only letters, numbers, or underscores"},
		{"invalid-input", "input must be a single word containing only letters, numbers, or underscores"},
		{"invalid@input", "input must be a single word containing only letters, numbers, or underscores"},
		{"invalid#input", "input must be a single word containing only letters, numbers, or underscores"},
	}

	for _, test := range tests {
		err := ValidateSingleWord(test.input)
		if err != nil && err.Error() != test.expected {
			t.Errorf("ValidateSingleWord(%q) returned error %q; expected %q", test.input, err.Error(), test.expected)
		}
		if err == nil && test.expected != "" {
			t.Errorf("ValidateSingleWord(%q) returned no error; expected %q", test.input, test.expected)
		}
	}
}
