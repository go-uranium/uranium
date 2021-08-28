package uranium

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	"github.com/go-uranium/uranium/model/session"
)

// AuthUser validates the cookie "token",
// it matches the resource which can be accessed by every user.
func (uranium *Uranium) AuthUser(ctx *fiber.Ctx) error {
	// try to find token in cookie
	token := ctx.Cookies("token")
	// no token found
	if len(token) == 0 {
		return ErrTokenRequired
	}
	// try to find token in database(cache)
	cache, _, err := uranium.cache.ValidSessionByToken(token)
	if err != nil {
		// if token not found in db
		if err == sql.ErrNoRows {
			// wipe client token to avoid too many invalid request
			ctx.Cookie(WipeToken)
			return ErrInvalidToken
		}
		// unexpected error
		return err
	}
	// if token has been expired
	if !cache.Valid {
		// wipe client token to avoid too many invalid request
		ctx.Cookie(WipeToken)
		return ErrTokenExpired
	}
	// pass
	return nil
}

// SudoAuth validates the cookie "token_sudo" and "token_admin",
// it matches the resource which can be accessed by user in sudo mode and super admin.
func (uranium *Uranium) SudoAuth(ctx *fiber.Ctx) (int16, error) {
	// try to get sudo token
	sudoToken := ctx.Cookies("token_sudo")
	// try to get admin token
	adminToken := ctx.Cookies("token_admin")
	// neither is found
	if len(sudoToken) == 0 && len(adminToken) == 0 {
		return session.UNKNOWN, ErrSudoTokenRequired
	}

	// if admin token found in cookies
	if len(adminToken) != 0 {
		// try to find admin token in db
		adminSess, _, err := uranium.cache.ValidSessionByToken(adminToken)
		if err != nil {
			// admin token not found in db
			if err == sql.ErrNoRows {
				// wipe admin token to avoid too many invalid requests
				ctx.ClearCookie("token_admin")
				return session.UNKNOWN, ErrInvalidAdminToken
			}
			// unexpected error
			return session.UNKNOWN, err
		}
		// admin token has been expired
		if !adminSess.Valid {
			// wipe admin token to avoid too many invalid requests
			ctx.ClearCookie("token_admin")
			return session.UNKNOWN, ErrAdminTokenExpired
		}
		// pass as admin
		return session.ADMIN, nil
	}

	// if admin token not found, then try to verify sudo token.
	// try to find sudo token in db
	sudoSess, _, err := uranium.cache.ValidSessionByToken(sudoToken)
	if err != nil {
		// sudo token not found in db
		if err == sql.ErrNoRows {
			// wipe sudo token to avoid too much invalid requests
			ctx.ClearCookie("token_sudo")
			return session.UNKNOWN, ErrInvalidSudoToken
		}
		// unexpected error
		return session.UNKNOWN, err
	}
	// sudo token has been expired
	if !sudoSess.Valid {
		// wipe sudo token to avoid too much invalid requests
		ctx.ClearCookie("token_sudo")
		return session.UNKNOWN, ErrSudoTokenExpired
	}
	// pass as user in sudo mode
	return session.SUDO, nil
}
