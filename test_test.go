package ushio_test

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	a := []byte{'z', 'x', 'c', 'v', 'n', 'm', 't', 'r', 'w',
		'e', 'u', 'o', 'i', 's', 'a'}
	for _, i := range a {
		for _, j := range a {
			fmt.Printf("%v%v%v%v.net\n", string(i), string(j), string(i), string(i))
		}
	}
}
