package validate

import (
	"net"
	"regexp"
	"strings"
)

var EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var EmailCheckMX = true

func Email(e string) bool {
	// the shortest example: i@ai
	// which is still a valid email address
	if len(e) < 4 || len(e) > 255 {
		return false
	}
	if !EmailRegex.MatchString(e) {
		return false
	}
	if EmailCheckMX {
		parts := strings.Split(e, "@")
		mx, err := net.LookupMX(parts[1])
		if err != nil || len(mx) == 0 {
			return false
		}
	}
	return true
}
