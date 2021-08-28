package uranium

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/go-uranium/uranium/model/session"
	"github.com/go-uranium/uranium/model/user"
	"github.com/go-uranium/uranium/utils/recaptcha"
	"github.com/go-uranium/uranium/utils/token"
)

type UserLoginReq struct {
	Username  string `json:"username"`
	Password  []byte `json:"password"`
	IsEmail   bool   `json:"is_email"`
	Remember  bool   `json:"remember"`
	Recaptcha string `json:"recaptcha"`
	Type      int16  `json:"type"`
}

func (uranium *Uranium) HandleUserLogin(ctx *fiber.Ctx) error {
	req := &UserLoginReq{}
	if err := ctx.BodyParser(req); err != nil {
		return err
	}
	if uranium.config.LoginRecaptcha {
		passed, err := recaptcha.Verify(req.Recaptcha)
		if err != nil {
			return err
		}
		if !passed {
			return ErrRecaptchaFailed
		}
	}
	var auth *user.Auth
	var err error
	if req.IsEmail {
		auth, err = uranium.storage.UserAuthByEmail(req.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				return ErrUsernameNotFound
			}
			return err
		}
	} else {
		uid, err := uranium.storage.UserUIDByUsername(req.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				return ErrUsernameNotFound
			}
			return err
		}
		auth, err = uranium.storage.UserAuthByUID(uid)
		if err != nil {
			if err == sql.ErrNoRows {
				return ErrUsernameNotFound
			}
			return err
		}
	}
	if auth.LockedOrDisabled() {
		return ErrUserLocked
	}
	if !auth.PasswordValid(req.Password) {
		return ErrWrongPassword
	}
	req.Type = session.CleanTypeWithDefault(req.Type)
	b, err := uranium.storage.UserBasicByUID(auth.UID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUsernameNotFound
		}
		return err
	}
	bc := b.Core()
	if !session.ValidSessionType(bc.Admin, req.Type) {
		return ErrInvalidTokenType
	}
	now := time.Now()
	sess := &session.Session{
		Token:   token.New(),
		UID:     bc.UID,
		Mode:    req.Type,
		UA:      ctx.Get("User-Agent"),
		IP:      ctx.IP(),
		Created: now,
	}
	cookie := &fiber.Cookie{
		Value: sess.Token,
	}
	switch req.Type {
	case session.USER:
		cookie.Name = "token"
		if req.Remember {
			cookie.Expires = now.Add(30 * 24 * time.Hour)
			sess.Expire = cookie.Expires.Add(5 * time.Minute)
		} else {
			sess.Expire = now.Add(24 * time.Hour)
		}
	case session.SUDO:
		cookie.Name = "token_sudo"
		sess.Expire = cookie.Expires.Add(20 * time.Minute)
	case session.MODERATOR:
		cookie.Name = "token_mod"
		sess.Expire = cookie.Expires.Add(20 * time.Minute)
	case session.ADMIN:
		cookie.Name = "token_admin"
		sess.Expire = cookie.Expires.Add(20 * time.Minute)
	}
	err = uranium.storage.SessionInsertSession(sess)
	if err != nil {
		return err
	}
	ctx.Cookie(cookie)
	return ctx.JSON(&SuccessResp{
		Success: true,
	})
}

func (uranium *Uranium) HandleUserInfoByUID(ctx *fiber.Ctx) error {
	if err := uranium.AuthUser(ctx); err != nil {
		return err
	}

	// Process request
	// get params
	uid, err := strconv.Atoi(ctx.Params("uid"))
	if err != nil {
		return ErrInvalidUID
	}
	u, err := uranium.storage.UserByUID(int32(uid))
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}
	return ctx.JSON(u)
}

func (uranium *Uranium) HandleUserBasicByUID(ctx *fiber.Ctx) error {
	if err := uranium.AuthUser(ctx); err != nil {
		return err
	}

	// Process request
	uid, err := strconv.Atoi(ctx.Params("uid"))
	if err != nil {
		return ErrInvalidUID
	}
	// query from cache
	userb, _, err := uranium.cache.UserBasicByUID(int32(uid))
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}
	return ctx.JSON(userb)
}

func (uranium *Uranium) HandleUserProfileByUID(ctx *fiber.Ctx) error {
	if err := uranium.AuthUser(ctx); err != nil {
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
	if err := uranium.AuthUser(ctx); err != nil {
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

func (uranium *Uranium) HandleUserInfoByUsername(ctx *fiber.Ctx) error {
	if err := uranium.AuthUser(ctx); err != nil {
		return err
	}

	// Process request
	// get params
	username := ctx.Params("username")
	u, err := uranium.storage.UserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}
	return ctx.JSON(u)
}

func (uranium *Uranium) HandleUserProfileByUsername(ctx *fiber.Ctx) error {
	if err := uranium.AuthUser(ctx); err != nil {
		return err
	}

	username := ctx.Params("username")
	uid, _, err := uranium.cache.UserUIDByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}
	profile, err := uranium.storage.UserProfileByUID(uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}
	return ctx.JSON(profile)
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
