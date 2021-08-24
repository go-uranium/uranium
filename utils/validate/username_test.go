package validate_test

import (
	"testing"

	"github.com/go-uranium/uranium/utils/validate"
)

func TestUsername(t *testing.T) {
	if !validate.Username("iochen") {
		t.Error("expected true, got false")
	}
	if !validate.Username("1admin1") {
		t.Error("expected true, got false")
	}
	if !validate.Username("i1-1li") {
		t.Error("expected true, got false")
	}

	if validate.Username("i") {
		t.Error("expected false, got true")
	}
	if validate.Username("long-username-that-not-allowed") {
		t.Error("expected false, got true")
	}
	if validate.Username("---") {
		t.Error("expected false, got true")
	}
	if validate.Username("-i1l") {
		t.Error("expected false, got true")
	}
	if validate.Username("-ness") {
		t.Error("expected false, got true")
	}
	if validate.Username("anti-") {
		t.Error("expected false, got true")
	}
	if validate.Username("r chen") {
		t.Error("expected false, got true")
	}
	if validate.Username("") {
		t.Error("expected false, got true")
	}
}
