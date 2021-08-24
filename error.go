package uranium

import "net/http"

// ========================= TOKEN =============================
// In normal cases, user shouldn't meet these errors,
// unless they have modified their cookies.
var (
	// When token is need but no token found.
	ErrTokenRequired = NewError(http.StatusUnauthorized, "Token required.")
	// When the action requires sudo permission but no sudo token found.
	ErrSudoTokenRequired = NewError(http.StatusUnauthorized, "Sudo token required.")
	// When the resource is only available to mod/admin, then return this error.
	ErrNoPermission = NewError(http.StatusForbidden, "No permission to the resource.")

	// When token can be found in user request but cannot be found in database.
	ErrInvalidToken = NewError(http.StatusUnauthorized, "Invalid token.")
	// When token can be found in user request and in database, but has been expired.
	ErrTokenExpired = NewError(http.StatusUnauthorized, "Token has been expired.")

	// When sudo token found in user request but cannot be found in database.
	ErrInvalidSudoToken = NewError(http.StatusUnauthorized, "Invalid sudo token.")
	// When sudo token found in user request and in database, but has been expired.
	ErrSudoTokenExpired = NewError(http.StatusUnauthorized, "Sudo token has been expired.")

	// When mod token found in user request but cannot be found in database.
	ErrInvalidModToken = NewError(http.StatusUnauthorized, "Invalid mod token.")
	// When mod token found in user request and in database, but has been expired.
	ErrModTokenExpired = NewError(http.StatusUnauthorized, "Mod token has been expired.")

	// When admin token found in user request but cannot be found in database.
	ErrInvalidAdminToken = NewError(http.StatusUnauthorized, "Invalid admin token.")
	// When admin token found in user request and in database, but has been expired.
	ErrAdminTokenExpired = NewError(http.StatusUnauthorized, "Admin token has been expired.")
)
