package hash_test

import (
	"fmt"
	"testing"

	"github.com/go-ushio/ushio/utils/hash"
)

func TestHash(t *testing.T) {
	fmt.Println(hash.Hash("password"))
}
