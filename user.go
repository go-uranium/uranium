package ushio

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/go-ushio/ushio/core/user"
)

func (ushio *Ushio) UserHandler(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	if len(name) < 1 || len(name) > 10 {
		return fiber.NewError(400, "Invalid username or uid.")
	}

	u := &user.User{}
	uid, err := strconv.Atoi(name)
	if err == nil {
		u, err = ushio.Data.UserByUID(uid)
		if err != nil {
			if err == sql.ErrNoRows {
				return fiber.NewError(404, "User not found.")
			}
			return err
		}
	} else {
		u, err = ushio.Data.UserByUsername(name)
		if err != nil {
			if err == sql.ErrNoRows {
				return fiber.NewError(404, "User not found.")
			}
			return err
		}
	}

	nav, err := ushio.NavFromCtx(ctx)
	if err != nil {
		return err
	}

	return ctx.Render("user", fiber.Map{
		"Meta": &Meta{
			Config:      *ushio.Config,
			CurrentPage: "/u/" + u.Username,
		},
		"Nav":  nav,
		"User": u,
	})
}
