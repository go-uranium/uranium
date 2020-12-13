package ushio_test

import (
	"os"
	"testing"

	"github.com/go-ushio/ushio"
)

func TestStart(t *testing.T) {
	ushio.DefaultConfig.SQL = os.Getenv("DATA_SOURCE_NAME")
	err := ushio.Start(":8888", ushio.DefaultConfig)
	if err != nil {
		return
	}
}
