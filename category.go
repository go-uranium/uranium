package ushio

import (
	"github.com/gofiber/fiber/v2"
)

func (ushio *Ushio) HandleCategory(ctx *fiber.Ctx) error {
	// no database writing operations,
	// lock is unnecessary
	nav, err := ushio.NavFromCtx(ctx)
	if err != nil {
		return err
	}

	tname := ctx.Params("tname")
	category := ushio.Cache.Category(tname)
	if category == nil {
		return fiber.NewError(404, "Category not found.")
	}

	infos, err := ushio.Data.PostInfoCategory(ushio.Cache.IndexSize(), category.TID)
	if err != nil {
		return err
	}

	return ctx.Render("category", fiber.Map{
		"Meta": Meta{
			Config:      *ushio.Config,
			CurrentPage: "Home",
		},
		"Categories": ushio.Cache.Categories(),
		"Nav":        nav,
		"Data":       infos,
		"Category":   category,
	})
}
