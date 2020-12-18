package ushio

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	"github.com/go-ushio/ushio/cache"
	"github.com/go-ushio/ushio/user"
)

type Nav struct {
	User     *user.SimpleUser
	LoggedIn bool
}

func NavFromCtx(c *fiber.Ctx) (*Nav, error) {
	sessionToken := c.Cookies("token")
	nav := &Nav{
		LoggedIn: false,
	}
	ss, err := cache.SessionByToken(sessionToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nav, nil
		}
		return nav, err
	}
	if !ss.IsValid() {
		return nav, nil
	}

	u, err := cache.UserByUID(ss.UID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &Nav{}, nil
		}
		return &Nav{}, err
	}
	nav.User = u
	nav.LoggedIn = true
	return nav, nil
}
