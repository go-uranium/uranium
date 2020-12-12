package ushio

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/go-ushio/ushio/data"
	"github.com/go-ushio/ushio/user"
)

func UserHandler(c *fiber.Ctx) error {
	name := c.Params("name")
	if len(name) < 1 || len(name) > 10 {
		return fiber.NewError(400, "Invalid username or uid.")
	}

	u := &user.User{}
	uid, err := strconv.Atoi(name)
	if err == nil {
		u, err = data.UserByUID(uid)
		if err != nil {
			if err == sql.ErrNoRows {
				return fiber.NewError(404, "User not found.")
			}
			return err
		}
	} else {
		u, err = data.UserByUsername(name)
		if err != nil {
			if err == sql.ErrNoRows {
				return fiber.NewError(404, "User not found.")
			}
			return err
		}
	}

	err = c.JSON(u)
	if err != nil {
		return err
	}

	c.Render()

	return nil
}
