package config

import (
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

type Site struct {
	Title string `toml:"name",json:"name"`
	Logo  string `toml:"logo",json:"logo"`
	Icon  string `toml:"icon",json:"icon"`
}

var SiteConf *Site

func NewSite(r io.Reader) (*Site,error) {
	site := &Site{}
	_, err := toml.DecodeReader(r, site)
	return site,err
}

func NewSiteFromFile(path string) (*Site,error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return &Site{}, err
	}
	defer file.Close()
	site, err := NewSite(file)
	return site,err
}