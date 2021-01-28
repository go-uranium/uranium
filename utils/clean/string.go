package clean

import "strings"

// String cleans string by removing white space at leading and trailing
// and lowercase the string
func String(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}
