package mdparse

import (
	"fmt"
	"html"
	"regexp"

	"github.com/russross/blackfriday/v2"
)

func Parse(in string) []byte {
	s := html.EscapeString(in)
	run := blackfriday.Run([]byte(s))
	imgCom := regexp.MustCompile(`<img src="((?:(?!(?:(?:.+\.|)iochen\.com)).)+)" alt="(?:.+|)" />`)
	imgSrcList := imgCom.FindAllStringSubmatch(string(run), 1)
	fmt.Println(len(imgSrcList))
	for i := range imgSrcList {
		fmt.Println(imgSrc[i])
	}

	return run
}
