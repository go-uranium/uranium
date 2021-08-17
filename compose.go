package uranium

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/go-uranium/uranium/model/markdown"
	"github.com/go-uranium/uranium/model/post"
	"github.com/go-uranium/uranium/model/user"
	"github.com/go-uranium/uranium/utils/mdparse"
)

func (ushio *Ushio) HandlePOSTCompose(ctx *fiber.Ctx) error {
	ushio.Lock.RLock()
	defer ushio.Lock.RUnlock()
	nav, err := ushio.NavFromCtx(ctx)
	if err != nil {
		return err
	}

	if !nav.LoggedIn {
		return ctx.Redirect("/", http.StatusSeeOther)
	}

	content, err := mdparse.Parse(ctx.FormValue("compose-content"))
	if err != nil {
		return err
	}

	now := time.Now()
	p := &post.Post{
		Info: &post.Info{
			Title:     ctx.FormValue("title"),
			Creator:   user.Basic{UID: nav.User.UID},
			CreatedAt: now,
			LastMod:   now,
			Activity:  now,
		},
		Content: *content,
		Markdown: &markdown.Markdown{
			Type:    markdown.TYPE_POST,
			Content: ctx.FormValue("compose-content"),
		},
	}
	pid, err := ushio.Data.InsertPost(p)
	if err != nil {
		return err
	}

	err = ushio.Cache.IndexPostInfoRefresh()
	if err != nil {
		return err
	}

	return ctx.Redirect("/p/"+strconv.Itoa(int(pid)), http.StatusSeeOther)
}

func (ushio *Ushio) HandleCompose(ctx *fiber.Ctx) error {
	// no database writing operations,
	// lock is unnecessary
	nav, err := ushio.NavFromCtx(ctx)
	if err != nil {
		return err
	}

	if !nav.LoggedIn {
		return ctx.Redirect("/", http.StatusSeeOther)
	}

	return ctx.Render("compose", fiber.Map{
		"Meta": Meta{
			Config:      *ushio.Config,
			CurrentPage: "compose",
		},
		"Nav": nav,
	})
}
