package validators

import (
	"errors"
	"regexp"
)

// ValidateSingleWord checks if the input is a single word without spaces or special characters.
func ValidateSingleWord(input string) error {
	// Regular expression to match a single word (letters, numbers, underscores only)
	validWordRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !validWordRegex.MatchString(input) {
		return errors.New("input must be a single word containing only letters, numbers, or underscores")
	}
	return nil
}
