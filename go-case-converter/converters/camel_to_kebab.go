package converters

import (
	"strings"
	"unicode"
)

// ConvertCamelToKebab converts a camelCase string to kebab-case.
func ConvertCamelToKebab(camel string) string {
	var result []string
	for i, char := range camel {
		if unicode.IsUpper(char) && i > 0 {
			result = append(result, "-")
		}
		result = append(result, string(unicode.ToLower(char)))
	}
	return strings.Join(result, "")
}
