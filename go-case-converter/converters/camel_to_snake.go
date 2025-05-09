package converters

import (
	"strings"
	"unicode"
)

// ConvertCamelToSnake converts a camel case string to a snake case string, ignoring underscores in the input.
func ConvertCamelToSnake(input string) string {
	var result []string
	for i, char := range input {
		if char == '_' {
			// Skip underscores in the input
			continue
		}
		if unicode.IsUpper(char) && i > 0 {
			result = append(result, "_")
		}
		result = append(result, string(unicode.ToLower(char)))
	}
	return strings.Join(result, "")
}
