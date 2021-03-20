package ushio

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/go-ushio/ushio/core/user"
)

func (ushio *Ushio) HandleUser(ctx *fiber.Ctx) error {
	// no database writing operations,
	// lock is unnecessary
	name := ctx.Params("name")
	if len(name) < 1 || len(name) > 10 {
		return fiber.NewError(400, "Invalid username or uid.")
	}

	u := &user.User{}
	uid, err := strconv.Atoi(name)
	if err == nil {
		u, err = ushio.Data.UserByUID(int64(uid))
		if err != nil {
			if err == sql.ErrNoRows {
				return fiber.NewError(404, "User not found.")
			}
			return err
		}
		return ctx.Redirect("/u/"+u.Username, 307)
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

	posts, err := ushio.Data.PostedBy(u.UID)
	if err != nil {
		if err == sql.ErrNoRows {
			// do something
		} else {
			return err
		}
	}

	return ctx.Render("user", fiber.Map{
		"Meta": &Meta{
			Config:      *ushio.Config,
			CurrentPage: fmt.Sprintf("%s (@%s)", u.Name, u.Username),
		},
		"Nav":   nav,
		"User":  u,
		"Posts": posts,
	})
}
