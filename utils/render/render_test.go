package render_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-ushio/ushio/utils/render"
)

func TestEngine_Load(t *testing.T) {
	engine := render.New("test", "test/partials",".html")
	err := engine.Load()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(engine)

	err = engine.Render(os.Stdout, "index", nil)
	if err != nil {
		t.Error(err)
	}
}