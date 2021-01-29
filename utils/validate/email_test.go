package validate_test

import (
	"fmt"
	"testing"

	"github.com/go-ushio/ushio/utils/validate"
)

func TestEmail(t *testing.T) {
	fmt.Println(validate.Email("i@ai"))
	fmt.Println(validate.Email("i@iochen.com"))
	fmt.Println(validate.Email("iochen.com@gmail.com"))
	fmt.Println(validate.Email("i+ushio@iochen.com"))
	fmt.Println(validate.Email("i@nomx.iochen.com"))
	fmt.Println(validate.Email("我@iochen.com"))
	fmt.Println(validate.Email("thesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseen@iochen.com"))
}
