package validate_test

import (
	"fmt"
	"testing"
)

func TestEmail(t *testing.T) {
	fmt.Println(Email("i@ai"))
	fmt.Println(Email("i@iochen.com"))
	fmt.Println(Email("iochen.com@gmail.com"))
	fmt.Println(Email("i+uranium@iochen.com"))
	fmt.Println(Email("i@nomx.iochen.com"))
	fmt.Println(Email("我@iochen.com"))
	fmt.Println(Email("thesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseen@iochen.com"))
}
