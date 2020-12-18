package mdparse_test

import (
	"fmt"
	"testing"

	"github.com/go-ushio/ushio/utils/mdparse"
)

func TestParse(t *testing.T) {
	testStr :=
		`
# H1
## H2
### H3
#### H4
* *Italian* 
* **Bold** 
* ***Italian Bold***

![](https://fff.com)
` + "```" + "go\n" +
			`<a href="javascript:alert('XSS1')" onmouseover="alert('XSS2')">XSS<a>

` + "```"

	html, err := mdparse.Parse(testStr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*html)
}
