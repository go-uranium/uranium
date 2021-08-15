package validate_test

import (
	"fmt"
	"testing"
)

func TestUsername(t *testing.T) {
	fmt.Println(Username("i"))
	fmt.Println(Username("iochen"))
	fmt.Println(Username("longusername"))
	fmt.Println(Username("i1_1li"))
	fmt.Println(Username("___"))
	fmt.Println(Username("_i1l"))
	fmt.Println(Username("_un"))
	fmt.Println(Username("1"))
	fmt.Println(Username("1admin"))
	fmt.Println(Username("0"))

}
