package ushio

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/go-ushio/ushio/core/post"
	"github.com/go-ushio/ushio/utils/mdparse"
)

func (ushio *Ushio) ComposePostHandler(ctx *fiber.Ctx) error {
	ushio.Lock.RLock()
	defer ushio.Lock.RUnlock()
	nav, err := ushio.NavFromCtx(ctx)
	if err != nil {
		return err
	}

	content, err := mdparse.Parse(ctx.FormValue("compose-content"))
	if err != nil {
		return err
	}

	now := time.Now()
	p := &post.Post{
		Info: &post.Info{
			Title:     ctx.FormValue("title"),
			Creator:   nav.User.UID,
			CreatedAt: now,
			LastMod:   now,
			Activity:  now,
		},
		Content:  *content,
		Markdown: ctx.FormValue("compose-content"),
	}
	if err := ushio.Data.InsertPost(p); err != nil {
		return err
	}

	err = ushio.Cache.IndexPostInfoDrop()
	if err != nil {
		return err
	}

	return ctx.Redirect("/", 303)
}

func (ushio *Ushio) ComposeHandler(ctx *fiber.Ctx) error {
	// no database writing operations,
	// lock is unnecessary
	nav, err := ushio.NavFromCtx(ctx)
	if err != nil {
		return err
	}

	return ctx.Render("compose", fiber.Map{
		"Meta": Meta{
			Config:      *ushio.Config,
			CurrentPage: "home",
		},
		"Nav": nav,
	})
}
