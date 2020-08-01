package controllor

import (
	"fmt"
	"testing"
)

func TestQueryUser(t *testing.T) {
	user, err := QueryUser("1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(user)

	user, err = QueryUser("admin")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(user)
}
