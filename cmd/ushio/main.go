package main

import (
	"fmt"

	"github.com/gofiber/fiber"

	"github.com/go-ushio/ushio"
	"github.com/go-ushio/ushio/config"
	"github.com/go-ushio/ushio/controllor"
)

func main() {
	err := controllor.Init("root:QuyWlWOmH2K8t2wDzzK3Ttsg21Agn5vr5YNK0Nbp7N3uBN3xAu@tcp(localhost:3306)/ushio?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}

	u := ushio.NewUshio(fiber.New())
	site := &config.Site{}
	site.Title = "Ushio"
	site.Icon = "https://iochen.com/images/favicon-32x32-mrchen.png"
	site.Logo = "https://iochen.com/images/avantar.png"
	config.SiteConf = site
	println(u.Start())
}