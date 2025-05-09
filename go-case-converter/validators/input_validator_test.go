package validators

import "testing"

func TestValidateSingleWord(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
	}{
		{"validWord", false},      // camelCase
		{"valid_word", false},     // snake_case
		{"valid-word", false},     // kebab-case
		{"Invalid Word", true},    // contains space
		{"invalid@word", true},    // contains special character
		{"", true},                // empty string
		{"123valid", false},       // starts with numbers
		{"valid123", false},       // ends with numbers
		{"valid_word-123", false}, // mixed snake and kebab
	}

	for _, test := range tests {
		err := ValidateSingleWord(test.input)
		if (err != nil) != test.expectError {
			t.Errorf("ValidateSingleWord(%q) = %v; want error: %v", test.input, err, test.expectError)
		}
	}
}
