package post

import (
	"fmt"
	"testing"
	"time"
)

func TestInfo_Copy(t *testing.T) {
	info := &Info{
		PID:       64,
		Title:     "Test Title",
		Creator:   0,
		CreatedAt: time.Now(),
		LastMod:   time.Now(),
		Replies:   12,
		Views:     2300,
		Activity:  time.Now(),
		Hidden:    false,
		Anonymous: false,
	}
	fmt.Println(*info)
	info2 := info.Copy()
	info2.Title = "Edited"
	info2.Activity = time.Now()
	fmt.Println(*info2)
	fmt.Println(*info)
}
