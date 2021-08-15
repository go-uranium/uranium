package clean

import "strings"

// Username cleans string by removing white space at leading and trailing
// and lowercase the string
func Username(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

// Email cleans string by removing white space at leading and trailing
// and lowercase the string
func Email(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}