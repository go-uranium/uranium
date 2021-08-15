package token

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestNew(t *testing.T) {
	fmt.Println(New())
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = uuid.New().String()
	}
}
