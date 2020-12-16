package ushio

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/go-ushio/ushio/cache"
	"github.com/go-ushio/ushio/data"
	"github.com/go-ushio/ushio/session"
	"github.com/go-ushio/ushio/user"
	"github.com/go-ushio/ushio/utils/recaptcha"
)

func LoginHandler(c *fiber.Ctx) error {
	sessionToken := c.Cookies("token")
	ss, err := cache.SessionByToken(sessionToken)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if ss != nil && ss.IsValid() {
		return c.Redirect("/home", 303)
	}

	return c.Render("_login", nil,
		"_login", "head")
}

func LoginPostHandler(c *fiber.Ctx) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	useEmail := c.FormValue("use-email")
	remA := c.FormValue("remember")
	reca := c.FormValue("g-recaptcha-response")

	s, err := recaptcha.Verify(reca)
	if err != nil {
		return err
	}

	if !s {
		return c.Render("_login", fiber.Map{
			"Warn": "reCAPTCHA not passed",
		},
			"_login", "head")
	}

	u := &user.User{}
	if useEmail == "on" {
		u, err = data.UserByEmail(email)
	} else {
		u, err = data.UserByUsername(username)
	}

	if err != nil {
		c.Status(401)
		if err == sql.ErrNoRows {
			return c.Render("_login", fiber.Map{
				"Warn": "user not found",
			},
				"_login", "head")
		}
		return err
	}

	if u.Valid(password) {
		rem := false
		if remA == "on" {
			rem = true
		}

		t := time.Now()
		s := &session.Session{
			UID:    u.UID,
			Token:  uuid.New().String(),
			UA:     string(c.Request().Header.UserAgent()),
			IP:     c.IP(),
			Time:   t,
			Expire: t.Add(24 * time.Hour),
		}

		if rem {
			s.Expire = t.Add(720 * time.Hour)
		}

		err := data.InsertSession(s)
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

		c.Cookie(ck)
		return c.Redirect("/home", 303)
	} else {
		c.Status(401)
		return c.Render("_login", fiber.Map{
			"Warn": "wrong password",
		}, "_login", "head")
	}

}

func SignUpHandler(c *fiber.Ctx) error {

	return nil
}
