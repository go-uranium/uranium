package ushio

import (
	"github.com/gofiber/fiber/v2"

	"github.com/go-ushio/ushio/core/category"
	"github.com/go-ushio/ushio/core/post"
	"github.com/go-ushio/ushio/core/user"
)

type IndexPosts struct {
	Info     *post.Info
	User     *user.User
	Category *category.Category
}

func (ushio *Ushio) HomeHandler(c *fiber.Ctx) error {
	// no database writing operations,
	// lock is unnecessary
	nav, err := ushio.NavFromCtx(c)
	if err != nil {
		return err
	}

	indexPosts := ushio.Cache.IndexPostInfo()

	return c.Render("home", fiber.Map{
		"Meta": Meta{
			Config:      *ushio.Config,
			CurrentPage: "Home",
		},
		"Categories": ushio.Cache.Categories(),
		"Nav":        nav,
		"Data":       indexPosts,
	})
}
