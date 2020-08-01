package view

import (
	"github.com/gofiber/fiber"

	"github.com/go-ushio/ushio/controllor"
)

func HandleUserAPI(c *fiber.Ctx) {
	user, err := controllor.QueryUser(c.Params("uid"))
	if err != nil {
		c.SendStatus(404)
		return
	}
	if err := c.JSON(user); err != nil {
		c.SendStatus(500)
		c.Next(err)
	}
}
