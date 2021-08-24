package validate

import (
	"strings"
	"unicode/utf8"
)

func Name(str string) bool {
	str = strings.TrimSpace(str)
	if len(str) < 1 || len(str) > 30 {
		return false
	}
	return utf8.ValidString(str)
}
