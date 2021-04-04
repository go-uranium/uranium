package ushio

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	"github.com/go-ushio/ushio/core/user"
)

type Nav struct {
	User     *user.User
	LoggedIn bool
}

func (ushio *Ushio) NavFromCtx(c *fiber.Ctx) (*Nav, error) {
	sessionToken := c.Cookies("token")
	nav := &Nav{
		LoggedIn: false,
	}
	ss, err := ushio.Cache.Session(sessionToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nav, nil
		}
		return nav, err
	}
	if !ss.Valid() {
		return nav, nil
	}

	u, err := ushio.Cache.User(ss.UID)
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

func (ushio *Ushio) AddArtifact(uid, add int64, reason string) error {
	err := ushio.Data.AddArtifact(uid, add)
	if err != nil {
		return err
	}
	return nil
}
