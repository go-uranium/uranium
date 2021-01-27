package ushio

import (
	"github.com/gofiber/fiber/v2"

	"github.com/go-ushio/ushio/core/post"
	"github.com/go-ushio/ushio/core/user"
)

type Present struct {
	Post *post.Info
	User *user.SimpleUser
}

func (ushio *Ushio) HomeHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		nav, err := ushio.NavFromCtx(c)
		if err != nil {
			return err
		}

		sps, err := ushio.Data.SimplePosts(0, 25)
		if err != nil {
			return err
		}

		var ps []Present

		for i := range sps {
			sp := sps[i]
			simpleUser, err := ushio.Cache.UserByUID(sp.Creator)
			if err != nil {
				return err
			}
			ps = append(ps,
				Present{
					sps[i],
					simpleUser,
				})
		}

		return c.Render("home", fiber.Map{
			"Nav":    nav,
			"Config": config,
			"Data":   ps,
		})
	}
}
