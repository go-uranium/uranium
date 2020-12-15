package ushio

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	return c.Render("_login", nil,
		"_login", "head")
}

func LoginPostHandler(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("%v", c.FormValue("email")))
}

func SignUpHandler(c *fiber.Ctx) error {

	return nil
}
