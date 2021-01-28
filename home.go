package ushio

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	"github.com/go-ushio/ushio/core/post"
	"github.com/go-ushio/ushio/core/user"
)

type IndexPosts struct {
	Info *post.Info
	User *user.User
}

func (ushio *Ushio) HomeHandler(c *fiber.Ctx) error {
	nav, err := ushio.NavFromCtx(c)
	if err != nil {
		return err
	}

	sps, err := ushio.Cache.IndexPostInfo(25)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	var ps []IndexPosts

	for i := range sps {
		sp := sps[i]
		user, err := ushio.Cache.UserByUID(sp.Creator)
		if err != nil {
			if err != sql.ErrNoRows {
				return err
			}
			user, err = ushio.Cache.UserByUID(0)
			if err != nil {
				return err
			}
		}
		ps = append(ps,
			IndexPosts{
				sps[i],
				user,
			})
	}

	return c.Render("home", fiber.Map{
		"Meta": Meta{
			Config:      *ushio.Config,
			CurrentPage: "home",
		},
		"Nav":  nav,
		"Data": ps,
	})
}
