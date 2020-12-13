package ushio

import "github.com/gofiber/fiber/v2"

func HomeHandler(c *fiber.Ctx) error {
	return c.Render("_base", fiber.Map{
		"Config": config,
	}, "_base", "body/home", "head", "nav", "footer")
}
