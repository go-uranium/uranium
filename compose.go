package ushio

import (
	"github.com/gofiber/fiber/v2"

	"github.com/go-ushio/ushio/data"
	"github.com/go-ushio/ushio/post"
	"github.com/go-ushio/ushio/utils/mdparse"
)

func ComposePostHandler(c *fiber.Ctx) error {
	nav, err := NavFromCtx(c)
	if err != nil {
		return err
	}

	content, err := mdparse.Parse(c.FormValue("compose-content"))
	if err != nil {
		return err
	}

	p := &post.Post{
		Title:    c.FormValue("title"),
		Creator:  nav.User.UID,
		Content:  *content,
		Markdown: c.FormValue("compose-content"),
	}
	if err := data.InsertPost(p); err != nil {
		return err
	}

	return c.Redirect("/", 303)
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
