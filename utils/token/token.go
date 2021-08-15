package token

import (
	"crypto/rand"
	"encoding/base64"
)

func New() string {
	r, err := random(24)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(r)
}

func NewInt(l int) string {
	r, err := random(l)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(r)
}

func random(length int) ([]byte, error) {
	r := make([]byte, length)
	_, err := rand.Read(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
