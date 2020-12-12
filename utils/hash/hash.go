package hash

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

var (
	SaltFormat = "!!1 iMp0R7AnT :%s)^36E#7+@5*$89^# ||!"
)

func Hash(str string) []byte {
	h := sha256.Sum256([]byte(fmt.Sprintf(SaltFormat, str)))
	return h[:]
}

func Compare(hash []byte, str string) bool {
	h := Hash(str)
	if bytes.Compare(hash, h) == 0 {
		return true
	}
	return false
}
