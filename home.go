package ushio

import (
	"github.com/gofiber/fiber/v2"
)

func HomeHandler(c *fiber.Ctx) error {
	nav, err := NavFromCtx(c)
	if err != nil {
		return err
	}

	return c.Render("_base", fiber.Map{
		"Nav":    nav,
		"Config": config,
	}, "_base", "body/home", "head", "nav", "footer")
}
