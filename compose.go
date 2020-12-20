package ushio

import (
	"github.com/gofiber/fiber/v2"

	"github.com/go-ushio/ushio/utils/mdparse"
)

func ComposePostHandler(c *fiber.Ctx) error {
	html, e := mdparse.Parse(c.FormValue("compose-content"))
	if e != nil {
		return e
	}
	_, e = c.Write([]byte(string(*html)))
	return e
}

func ComposeHandler(c *fiber.Ctx) error {
	nav, err := NavFromCtx(c)
	if err != nil {
		return err
	}

	return c.Render("_compose", fiber.Map{
		"Nav":    nav,
		"Config": config,
	}, "_compose", "editor_head", "head", "nav", "footer")
}
