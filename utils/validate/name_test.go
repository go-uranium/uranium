package validate_test

import (
	"testing"

	"github.com/go-uranium/uranium/utils/validate"
)

func TestName(t *testing.T) {
	if !validate.Name("Richard Chen") {
		t.Error("expected true, got false")
	}

	if validate.Name("Richard Chen - This is too long that not allowed") {
		t.Error("expected false, got true")
	}
	if validate.Name("\xbd\xb2\x3d") {
		t.Error("expected false, got true")
	}
}
