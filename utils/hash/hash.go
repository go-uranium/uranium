package hash

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
)

func SHA256(b []byte) []byte {
	h := sha256.Sum256(b)
	return h[:]
}

func SHA256Compare(hash []byte, b []byte) bool {
	h := SHA256(b)
	if bytes.Compare(hash, h) == 0 {
		return true
	}
	return false
}

func MD5(b []byte) []byte {
	h := md5.Sum(b)
	return h[:]
}
