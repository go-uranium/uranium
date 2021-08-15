package validate_test

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(Name("Richard Chen"))
	fmt.Println(Name("Richard Chen - Long Version"))
	fmt.Println(Name("\xbd\xb2\x3d"))
}
