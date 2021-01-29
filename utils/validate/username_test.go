package validate_test

import (
	"fmt"
	"testing"

	"github.com/go-ushio/ushio/utils/validate"
)

func TestUsername(t *testing.T) {
	fmt.Println(validate.Username("i"))
	fmt.Println(validate.Username("iochen"))
	fmt.Println(validate.Username("longusername"))
	fmt.Println(validate.Username("i1_1li"))
	fmt.Println(validate.Username("___"))
	fmt.Println(validate.Username("_i1l"))
	fmt.Println(validate.Username("_un"))
	fmt.Println(validate.Username("1"))
	fmt.Println(validate.Username("1admin"))
	fmt.Println(validate.Username("0"))

}
