package validate

import (
	"regexp"
	"strings"
)

var UsernameRegex = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9-]{0,18}[a-zA-Z0-9]$`)

func Username(u string) bool {
	u = strings.TrimSpace(u)
	if len(u) < 1 || len(u) > 20 {
		return false
	}
	if !UsernameRegex.MatchString(u) {
		return false
	}
	return true
}
