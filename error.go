package uranium

import "net/http"

// ========================= TOKEN =============================
// In normal cases, user shouldn't meet these errors,
// unless they have modified their cookies.
var (
	// When token is need but no token found.
	ErrTokenRequired = NewError(1006, http.StatusUnauthorized, "Token required.")
	// When the action requires sudo permission but no sudo token found.
	ErrSudoTokenRequired = NewError(1007, http.StatusUnauthorized, "Sudo token required.")
	// When the action requires admin permission but no admin token found.
	ErrAdminTokenRequired = NewError(1008, http.StatusUnauthorized, "Admin token required.")
	// When the resource is only available to mod/admin, then return this error.
	ErrNoPermission = NewError(1009, http.StatusForbidden, "No permission to the resource.")

	// When token can be found in user request but cannot be found in database.
	ErrInvalidToken = NewError(1010, http.StatusUnauthorized, "Invalid token.")
	// When token can be found in user request and in database, but has been expired.
	ErrTokenExpired = NewError(1011, http.StatusUnauthorized, "Token has been expired.")

	// When sudo token found in user request but cannot be found in database.
	ErrInvalidSudoToken = NewError(1012, http.StatusUnauthorized, "Invalid sudo token.")
	// When sudo token found in user request and in database, but has been expired.
	ErrSudoTokenExpired = NewError(1013, http.StatusUnauthorized, "Sudo token has been expired.")

	// When mod token found in user request but cannot be found in database.
	ErrInvalidModToken = NewError(1014, http.StatusUnauthorized, "Invalid mod token.")
	// When mod token found in user request and in database, but has been expired.
	ErrModTokenExpired = NewError(1015, http.StatusUnauthorized, "Mod token has been expired.")

	// When admin token found in user request but cannot be found in database.
	ErrInvalidAdminToken = NewError(1016, http.StatusUnauthorized, "Invalid admin token.")
	// When admin token found in user request and in database, but has been expired.
	ErrAdminTokenExpired = NewError(1017, http.StatusUnauthorized, "Admin token has been expired.")
)

// ========================= USER =============================
var (
	ErrInvalidUID   = NewError(1018, http.StatusBadRequest, "Invalid UID.")
	ErrUserNotFound = NewError(1019, http.StatusNotFound, "User not found.")
)
