package mdparse

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

func Parse(in string) (*template.HTML, error) {
	body := struct {
		Text string `json:"text"`
		Mode string `json:"mode"`
	}{
		Text: in,
		Mode: "gfm",
	}
	ma, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(ma)
	resp, err := http.Post("https://api.github.com/markdown", "application/json", r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	html := template.HTML(all)
	return &html, nil
}
