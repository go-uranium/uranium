package view

import (
	"github.com/gofiber/fiber"

	"github.com/go-ushio/ushio/controllor/bind"
)

func HandleIndex(c *fiber.Ctx) {
	data, err := bind.IndexData(c)
	if err != nil {
		c.Next(err)
		return
	}

	err = c.Render("index", data)
	if err != nil {
		c.Next(err)
		return
	}
}