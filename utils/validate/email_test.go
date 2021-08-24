package validate_test

import (
	"testing"

	"github.com/go-uranium/uranium/utils/validate"
)

func TestEmail(t *testing.T) {
	validate.EmailCheckMX = false

	if !validate.Email("i@ai") {
		t.Error("expected true, got false")
	}
	if !validate.Email("i@iochen.com") {
		t.Error("expected true, got false")
	}
	if !validate.Email("iochen.com@gmail.com") {
		t.Error("expected true, got false")
	}
	if !validate.Email("i+uranium@iochen.com") {
		t.Error("expected true, got false")
	}
	if !validate.Email("i@nomx.iochen.com") {
		t.Error("expected true, got false")
	}
	if !validate.Email("thelongestdomainnameandemailaddress.thatyoucaneverfindintheworld@thelongestdomainnameandemailaddressthatyoucaneverfindintheworld.thelongestdomainnameandemailaddressthatyoucaneverfindintheworld.thelongestdomainnameandaddressthatyoucanfindintheworld.thelongestdomainnameandemailaddressthatyoucaneverfindintheworld.com") {
		t.Error("expected true, got false")
	}

	if validate.Email("ai") {
		t.Error("expected false, got true")
	}
	if validate.Email("i@i") {
		t.Error("expected false, got true")
	}
	if validate.Email("@ai.com") {
		t.Error("expected false, got true")
	}
	if validate.Email("我@iochen.com") {
		t.Error("expected false, got true")
	}
	if validate.Email("thesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseen@iochen.com") {
		t.Error("expected false, got true")
	}
}

func TestEmailWithMX(t *testing.T) {
	validate.EmailCheckMX = true

	if !validate.Email("i@iochen.com") {
		t.Error("expected true, got false")
	}
	if !validate.Email("iochen.com@gmail.com") {
		t.Error("expected true, got false")
	}
	if !validate.Email("thelongestdomainnameandemailaddress.thatyoucaneverfindintheworld@thelongestdomainnameandemailaddressthatyoucaneverfindintheworld.thelongestdomainnameandemailaddressthatyoucaneverfindintheworld.thelongestdomainnameandaddressthatyoucanfindintheworld.thelongestdomainnameandemailaddressthatyoucaneverfindintheworld.com") {
		t.Error("expected true, got false")
	}

	if validate.Email("i@nomx.iochen.com") {
		t.Error("expected false, got true")
	}
	if validate.Email("google@nomx.iochen.com") {
		t.Error("expected false, got true")
	}
	if validate.Email("thesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseenthesuperlongemailaddressihaveeverseen@iochen.com") {
		t.Error("expected false, got true")
	}
}
