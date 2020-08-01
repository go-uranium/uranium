package controllor

import (
	"fmt"
	"os"
)

func init() {
	err := Init(os.Getenv("SQL"))
	if err != nil {
		fmt.Println(err)
	}
}
