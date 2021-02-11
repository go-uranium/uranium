package validate

import "unicode/utf8"

func Name(str string) bool {
	if len(str) < 1 || len(str) > 30 {
		return false
	}
	return utf8.ValidString(str)
}
