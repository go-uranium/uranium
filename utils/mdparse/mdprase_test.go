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

<a href="javascript:alert('XSS1')" onmouseover="alert('XSS2')">XSS<a>
![](https://fff.com)
` + "```" + "go" +
			`<a href="javascript:alert('XSS1')" onmouseover="alert('XSS2')">XSS<a>

` + "```"

	fmt.Println(string(mdparse.Parse(testStr)))
}
