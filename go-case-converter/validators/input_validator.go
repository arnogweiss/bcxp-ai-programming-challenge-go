package validators

import (
	"errors"
	"regexp"
)

// ValidateSingleWord checks if the input is a single word in snake_case, kebab-case, or camelCase.
func ValidateSingleWord(input string) error {
	// Regular expression to match a single word in snake_case, kebab-case, or camelCase
	validWordRegex := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	if !validWordRegex.MatchString(input) {
		return errors.New("input must be a single word containing only letters, numbers, underscores, or hyphens")
	}
	return nil
}
