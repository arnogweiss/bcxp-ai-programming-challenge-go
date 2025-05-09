package converters

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ConvertKebabToCamel converts a kebab-case string to camelCase.
func ConvertKebabToCamel(kebab string) string {
	words := strings.Split(kebab, "-")
	titleCaser := cases.Title(language.Und)
	for i := range words {
		if i == 0 {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = titleCaser.String(words[i])
		}
	}
	return strings.Join(words, "")
}
