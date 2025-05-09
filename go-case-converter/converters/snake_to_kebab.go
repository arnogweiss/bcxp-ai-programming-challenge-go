package converters

import "strings"

// ConvertSnakeToKebab converts a snake_case string to kebab-case.
func ConvertSnakeToKebab(snake string) string {
	return strings.ReplaceAll(snake, "_", "-")
}
