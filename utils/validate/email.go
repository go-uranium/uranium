package validate

import (
	"net"
	"regexp"
	"strings"
)

var EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var EmailCheckMX = false

func Email(e string) bool {
	// the shortest example: i@ai
	// which is still a valid email address
	if len(e) < 4 || len(e) > 320 {
		return false
	}
	if !EmailRegex.MatchString(e) {
		return false
	}
	if EmailCheckMX {
		parts := strings.Split(e, "@")
		if len(parts) != 2 {
			return false
		}
		mx, err := net.LookupMX(parts[1])
		if err != nil || len(mx) == 0 {
			return false
		}
	}
	return true
}
