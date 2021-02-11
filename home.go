package ushio

import (
	"database/sql"

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

	indexInfos := ushio.Cache.IndexPostInfo()
	var indexPosts []IndexPosts

	for i := range indexInfos {
		sp := indexInfos[i]
		u, err := ushio.Cache.User(sp.Creator)
		if err != nil || u == nil {
			if err != sql.ErrNoRows {
				return err
			}
			u, err = ushio.Cache.User(0)
			if err != nil {
				return err
			}
		}
		indexPosts = append(indexPosts,
			IndexPosts{
				Info:     indexInfos[i],
				User:     u,
				Category: ushio.Cache.Category(sp.Category),
			})
	}

	return c.Render("home", fiber.Map{
		"Meta": Meta{
			Config:      *ushio.Config,
			CurrentPage: "Home",
		},
		"Nav":  nav,
		"Data": indexPosts,
	})
}
