package clean

import "strings"

// Username cleans string by removing white space at leading and trailing
func Username(str string) string {
	return strings.TrimSpace(str)
}

func Lowercase(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

// Name cleans string by removing white space at leading and trailing
func Name(str string) string {
	return strings.TrimSpace(str)
}

// Email cleans string by removing white space at leading and trailing
// and lowercase the string
func Email(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}
