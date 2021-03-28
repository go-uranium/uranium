package ushio

import (
	"github.com/gofiber/fiber/v2"
)

func (ushio *Ushio) HandleHome(ctx *fiber.Ctx) error {
	// no database writing operations,
	// lock is unnecessary
	nav, err := ushio.NavFromCtx(ctx)
	if err != nil {
		return err
	}

	indexPosts := ushio.Cache.IndexPostInfo()

	return ctx.Render("home", fiber.Map{
		"Meta": Meta{
			Config:      *ushio.Config,
			CurrentPage: "Home",
		},
		"Categories": ushio.Cache.Categories(),
		"Nav":        nav,
		"Data":       indexPosts,
	})
}
