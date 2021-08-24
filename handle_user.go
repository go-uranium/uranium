package uranium

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrInvalidUID   = NewError(http.StatusBadRequest, "Invalid UID.")
	ErrUserNotFound = NewError(http.StatusNotFound, "User not found.")
)

func (uranium *Uranium) HandleUserInfoByUID(ctx *fiber.Ctx) error {
	if err := uranium.PublicAuth(ctx); err != nil {
		return err
	}

	// Process request
	// get params
	uid, err := strconv.Atoi(ctx.Params("uid"))
	if err != nil {
		return ErrInvalidUID
	}
	user, err := uranium.storage.UserByUID(int32(uid))
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}
	return ctx.JSON(user)
}

func (uranium *Uranium) HandleUserBasicByUID(ctx *fiber.Ctx) error {
	if err := uranium.PublicAuth(ctx); err != nil {
		return err
	}

	// Process request
	uid, err := strconv.Atoi(ctx.Params("uid"))
	if err != nil {
		return ErrInvalidUID
	}
	// query from cache
	userb, err := uranium.cache.UserBasicByUID(int32(uid))
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}
	return ctx.JSON(userb)
}

func (uranium *Uranium) HandleUserProfileByUID(ctx *fiber.Ctx) error {
	if err := uranium.PublicAuth(ctx); err != nil {
		return err
	}

	uid, err := strconv.Atoi(ctx.Params("uid"))
	if err != nil {
		return ErrInvalidUID
	}
	profile, err := uranium.storage.UserProfileByUID(int32(uid))
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}
	return ctx.JSON(profile)
}

func (uranium *Uranium) HandleUserAuthByUID(ctx *fiber.Ctx) error {
	if err := uranium.PublicAuth(ctx); err != nil {
		return err
	}

	uid, err := strconv.Atoi(ctx.Params("uid"))
	if err != nil {
		return ErrInvalidUID
	}
	auth, err := uranium.storage.UserAuthByUID(int32(uid))
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}
	return ctx.JSON(auth)
}

//func (ushio *Ushio) HandleUser(ctx *fiber.Ctx) error {
//	// no database writing operations,
//	// lock is unnecessary
//	name := ctx.Params("name")
//	if len(name) < 1 || len(name) > 20 {
//		return fiber.NewError(http.StatusBadRequest, "Invalid username or uid.")
//	}
//
//	u := &user.User{}
//	uid, err := strconv.Atoi(name)
//	if err == nil {
//		u, err = ushio.Data.UserByUID(int64(uid))
//		if err != nil {
//			if err == sql.ErrNoRows {
//				return fiber.NewError(http.StatusNotFound, "User not found.")
//			}
//			return err
//		}
//		return ctx.Redirect("/u/"+u.Username, http.StatusTemporaryRedirect)
//	} else {
//		u, err = ushio.Data.UserByUsername(name)
//		if err != nil {
//			if err == sql.ErrNoRows {
//				return fiber.NewError(http.StatusNotFound, "User not found.")
//			}
//			return err
//		}
//	}
//
//	nav, err := ushio.NavFromCtx(ctx)
//	if err != nil {
//		return err
//	}
//
//	p := ctx.Query("p", "1")
//	page, err := strconv.Atoi(p)
//	if err != nil || page < 1 {
//		page = 1
//	}
//
//	ps := ushio.Config.PageSize
//	posts, err := ushio.Data.PostsInfoByUID(ps, int64(page-1)*ps, u.UID)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			// do something
//		} else {
//			return err
//		}
//	}
//
//	return ctx.Render("user", fiber.Map{
//		"Meta": &Meta{
//			Config:      *ushio.Config,
//			CurrentPage: fmt.Sprintf("%s (@%s)", u.Name, u.Username),
//		},
//		"Nav":    nav,
//		"User":   u,
//		"Posts":  posts,
//		"Active": 1,
//	})
//}
//
//func (ushio *Ushio) HandleUserComments(ctx *fiber.Ctx) error {
//	// no database writing operations,
//	// lock is unnecessary
//	name := ctx.Params("name")
//	if len(name) < 1 || len(name) > 20 {
//		return fiber.NewError(http.StatusUnauthorized, "Invalid username or uid.")
//	}
//
//	u := &user.User{}
//	uid, err := strconv.Atoi(name)
//	if err == nil {
//		u, err = ushio.Data.UserByUID(int64(uid))
//		if err != nil {
//			if err == sql.ErrNoRows {
//				return fiber.NewError(http.StatusNotFound, "User not found.")
//			}
//			return err
//		}
//		return ctx.Redirect("/u/"+u.Username+"/comments", http.StatusTemporaryRedirect)
//	} else {
//		u, err = ushio.Data.UserByUsername(name)
//		if err != nil {
//			if err == sql.ErrNoRows {
//				return fiber.NewError(http.StatusNotFound, "User not found.")
//			}
//			return err
//		}
//	}
//
//	nav, err := ushio.NavFromCtx(ctx)
//	if err != nil {
//		return err
//	}
//
//	p := ctx.Query("p", "1")
//	page, err := strconv.Atoi(p)
//	if err != nil || page < 1 {
//		page = 1
//	}
//
//	ps := ushio.Config.PageSize
//	posts, err := ushio.Data.PostsInfoByCommentCreator(ps, int64(page-1)*ps, u.UID)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			// do something
//		} else {
//			return err
//		}
//	}
//
//	return ctx.Render("user", fiber.Map{
//		"Meta": &Meta{
//			Config:      *ushio.Config,
//			CurrentPage: fmt.Sprintf("%s (@%s)", u.Name, u.Username),
//		},
//		"Nav":    nav,
//		"User":   u,
//		"Posts":  posts,
//		"Active": 2,
//	})
//}
