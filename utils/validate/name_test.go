package validate_test

import (
	"fmt"
	"testing"

	"github.com/go-ushio/ushio/utils/validate"
)

func TestName(t *testing.T) {
	fmt.Println(validate.Name("Richard Chen"))
	fmt.Println(validate.Name("Richard Chen - Long Version"))
	fmt.Println(validate.Name("\xbd\xb2\x3d"))
}
