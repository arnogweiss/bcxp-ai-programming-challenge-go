package converters

import "strings"

// ConvertKebabToSnake converts a kebab-case string to snake_case.
func ConvertKebabToSnake(kebab string) string {
	return strings.ReplaceAll(kebab, "-", "_")
}
