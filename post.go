package ushio

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/go-ushio/ushio/data"
)

func PostHandler(c *fiber.Ctx) error {
	nav, err := NavFromCtx(c)
	if err != nil {
		return err
	}

	postID := c.Params("post")
	if len(postID) < 1 || len(postID) > 10 {
		return fiber.NewError(400, "Invalid post id.")
	}

	pid, err := strconv.Atoi(postID)
	if err != nil {
		return fiber.NewError(400, "Invalid post id.")
	}

	post, err := data.PostByPID(pid)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(404, "Post not found.")
		}
		return err
	}

	err = c.Render("_base", fiber.Map{
		"Config": config,
		"Nav":    nav,
		"Post":   post,
	}, "_base", "body/post", "head", "nav", "footer")
	if err != nil {
		return err
	}

	return nil
}
