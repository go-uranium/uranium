package mdparse

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func Parse(in string) (*template.HTML, error) {
	r := strings.NewReader(in)
	resp, err := http.Post("https://api.github.com/markdown/raw", "text/plain", r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	html := template.HTML(bytes)
	return &html, nil
}
