package converters

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ConvertSnakeToCamel converts a snake_case string to camelCase.
func ConvertSnakeToCamel(snake string) string {
	words := strings.Split(snake, "_")
	titleCaser := cases.Title(language.Und) // Create a title caser for Unicode handling
	for i := range words {
		if i == 0 {
			// Lowercase the first word
			words[i] = strings.ToLower(words[i])
		} else {
			// Capitalize the first letter of subsequent words
			words[i] = titleCaser.String(words[i])
		}
	}
	return strings.Join(words, "")
}
