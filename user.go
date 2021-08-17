package uranium

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/go-uranium/uranium/model/user"
)

func (ushio *Ushio) HandleUser(ctx *fiber.Ctx) error {
	// no database writing operations,
	// lock is unnecessary
	name := ctx.Params("name")
	if len(name) < 1 || len(name) > 20 {
		return fiber.NewError(http.StatusBadRequest, "Invalid username or uid.")
	}

	u := &user.User{}
	uid, err := strconv.Atoi(name)
	if err == nil {
		u, err = ushio.Data.UserByUID(int64(uid))
		if err != nil {
			if err == sql.ErrNoRows {
				return fiber.NewError(http.StatusNotFound, "User not found.")
			}
			return err
		}
		return ctx.Redirect("/u/"+u.Username, http.StatusTemporaryRedirect)
	} else {
		u, err = ushio.Data.UserByUsername(name)
		if err != nil {
			if err == sql.ErrNoRows {
				return fiber.NewError(http.StatusNotFound, "User not found.")
			}
			return err
		}
	}

	nav, err := ushio.NavFromCtx(ctx)
	if err != nil {
		return err
	}

	p := ctx.Query("p", "1")
	page, err := strconv.Atoi(p)
	if err != nil || page < 1 {
		page = 1
	}

	ps := ushio.Config.PageSize
	posts, err := ushio.Data.PostsInfoByUID(ps, int64(page-1)*ps, u.UID)
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
		"Nav":    nav,
		"User":   u,
		"Posts":  posts,
		"Active": 1,
	})
}

func (ushio *Ushio) HandleUserComments(ctx *fiber.Ctx) error {
	// no database writing operations,
	// lock is unnecessary
	name := ctx.Params("name")
	if len(name) < 1 || len(name) > 20 {
		return fiber.NewError(http.StatusUnauthorized, "Invalid username or uid.")
	}

	u := &user.User{}
	uid, err := strconv.Atoi(name)
	if err == nil {
		u, err = ushio.Data.UserByUID(int64(uid))
		if err != nil {
			if err == sql.ErrNoRows {
				return fiber.NewError(http.StatusNotFound, "User not found.")
			}
			return err
		}
		return ctx.Redirect("/u/"+u.Username+"/comments", http.StatusTemporaryRedirect)
	} else {
		u, err = ushio.Data.UserByUsername(name)
		if err != nil {
			if err == sql.ErrNoRows {
				return fiber.NewError(http.StatusNotFound, "User not found.")
			}
			return err
		}
	}

	nav, err := ushio.NavFromCtx(ctx)
	if err != nil {
		return err
	}

	p := ctx.Query("p", "1")
	page, err := strconv.Atoi(p)
	if err != nil || page < 1 {
		page = 1
	}

	ps := ushio.Config.PageSize
	posts, err := ushio.Data.PostsInfoByCommentCreator(ps, int64(page-1)*ps, u.UID)
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
		"Nav":    nav,
		"User":   u,
		"Posts":  posts,
		"Active": 2,
	})
}
