package ushio

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/go-ushio/ushio/core/session"
	"github.com/go-ushio/ushio/core/user"
	"github.com/go-ushio/ushio/utils/recaptcha"
)

func (ushio *Ushio) LoginHandler(ctx *fiber.Ctx) error {
	sessionToken := ctx.Cookies("token")
	ss, err := ushio.Cache.SessionByToken(sessionToken)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if ss != nil && ss.IsValid() {
		return ctx.Redirect("/home?msg=Already Logged in!", 303)
	}

	return ctx.Render("login", nil)
}

func (ushio *Ushio) LoginPostHandler(ctx *fiber.Ctx) error {
	username := ctx.FormValue("username")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	useEmail := ctx.FormValue("use-email")
	remA := ctx.FormValue("remember")
	reca := ctx.FormValue("g-recaptcha-response")

	s, err := recaptcha.Verify(reca)
	if err != nil {
		return err
	}

	if !s {
		return ctx.Render("login", fiber.Map{
			"Warn": "reCAPTCHA not passed",
		})
	}

	u := &user.User{}
	if useEmail == "on" {
		u, err = ushio.Data.UserByEmail(email)
	} else {
		u, err = ushio.Data.UserByUsername(username)
	}

	if err != nil {
		ctx.Status(401)
		if err == sql.ErrNoRows {
			return ctx.Render("_login", fiber.Map{
				"Warn": "user not found",
			},
				"_login", "head")
		}
		return err
	}

	auth, err := ushio.Data.UserAuthByUID(u.UID)
	if err != nil {
		return err
	}

	if auth.Valid([]byte(password)) {
		rem := false
		if remA == "on" {
			rem = true
		}

		t := time.Now()
		s := &session.Session{
			UID:    u.UID,
			Token:  uuid.New().String(),
			UA:     string(ctx.Request().Header.UserAgent()),
			IP:     ctx.IP(),
			Time:   t,
			Expire: t.Add(36 * time.Hour),
		}

		if rem {
			s.Expire = t.Add(720 * time.Hour)
		}

		err := ushio.Data.InsertSession(s)
		if err != nil {
			return err
		}

		ck := &fiber.Cookie{
			Name:  "token",
			Value: s.Token,
			Path:  "/",
		}

		if rem {
			ck.Expires = s.Expire
		}

		ctx.Cookie(ck)
		return ctx.Redirect("/home", 303)
	} else {
		ctx.Status(401)
		return ctx.Render("login", fiber.Map{
			"Warn": "wrong password",
		})
	}
}

func (ushio *Ushio) SignUpHandler(ctx *fiber.Ctx) error {

	return nil
}
