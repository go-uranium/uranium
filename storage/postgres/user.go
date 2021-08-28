package postgres

import (
	"strings"

	"github.com/go-uranium/uranium/model/user"
	"github.com/go-uranium/uranium/utils/clean"
	"github.com/go-uranium/uranium/utils/sqlnull"
)

var SQLUserInsertUser = `INSERT INTO "user" VALUES 
                          (username = $1, lowercase = $2, electrons = $3, admin = $4, 
                           created = $5, deleted = $6) RETURNING uid;`

func (pg *Postgres) UserInsertUser(u *user.User) (int32, error) {
	var uid int32
	err := pg.db.QueryRow(SQLUserInsertUser, u.Username, strings.ToLower(u.Username),
		u.Electrons, u.Admin, u.Created, u.Deleted).Scan(&uid)
	return uid, err
}

var SQLUserInsertUserAuth = `INSERT INTO user_auth VALUES (uid = $1, email = $2, password = $3, 
                              security_email = $4, two_factor = $5, locked = $6, locked_till = $7);`

func (pg *Postgres) UserInsertUserAuth(auth *user.Auth) error {
	_, err := pg.db.Exec(SQLUserInsertUserAuth, &auth.UID, &auth.Email, &auth.Password,
		&auth.SecurityEmail, &auth.TwoFactor, &auth.Locked, &auth.LockedTill)
	return err
}

var SQLUserInsertUserProfile = `INSERT INTO user_profile VALUES (uid = $1, name = $2, 
                                 bio = $3, location = $4, birthday = $5, email = $6, social = $7);`

func (pg *Postgres) UserInsertUserProfile(profile *user.Profile) error {
	_, err := pg.db.Exec(SQLUserInsertUserProfile, &profile.UID, &profile.Name,
		&profile.Bio, &profile.Location, &profile.Birthday, &profile.Email, &profile.Social)
	return err
}

var SQLUserByUID = `SELECT uid, username, electrons, admin, created, deleted FROM "user" WHERE uid = $1;`

func (pg *Postgres) UserByUID(uid int32) (*user.User, error) {
	u := &user.User{}
	err := pg.db.QueryRow(SQLUserByUID, uid).
		Scan(&u.UID, &u.Username, &u.Electrons, &u.Admin, &u.Created, &u.Deleted)
	return u, err
}

var SQLUserBasicByUID = `SELECT uid, username, admin FROM "user" WHERE uid = $1;`

func (pg *Postgres) UserBasicByUID(uid int32) (*user.Basic, error) {
	bc := &user.BasicCore{}
	err := pg.db.QueryRow(SQLUserBasicByUID, uid).
		Scan(&bc.UID, &bc.Username, &bc.Admin)
	return user.NewBasicFromCore(bc), err
}

var SQLUserProfileByUID = `SELECT uid, name, bio, location, birthday, email, social FROM user_profile WHERE uid = $1;`

func (pg *Postgres) UserProfileByUID(uid int32) (*user.Profile, error) {
	pf := &user.Profile{}
	err := pg.db.QueryRow(SQLUserProfileByUID, uid).
		Scan(&pf.UID, &pf.Name, &pf.Bio, &pf.Location, &pf.Birthday,
			&pf.Email, &pf.Social)
	return pf, err
}

var SQLUserAuthByUID = `SELECT uid, email, password, security_email, two_factor, locked, locked_till, disabled FROM user_auth WHERE uid = $1;`

func (pg *Postgres) UserAuthByUID(uid int32) (*user.Auth, error) {
	au := &user.Auth{}
	err := pg.db.QueryRow(SQLUserAuthByUID, uid).
		Scan(&au.UID, &au.Email, &au.Password, &au.SecurityEmail, &au.TwoFactor,
			&au.Locked, &au.LockedTill, &au.Disabled)
	return au, err
}

var SQLUserByUsername = `SELECT uid, username, electrons, admin, created, deleted FROM "user" WHERE lowercase = $1;`

func (pg *Postgres) UserByUsername(username string) (*user.User, error) {
	u := &user.User{}
	err := pg.db.QueryRow(SQLUserByUsername, clean.Lowercase(username)).
		Scan(&u.UID, &u.Username, &u.Electrons, &u.Admin, &u.Created, &u.Deleted)
	return u, err
}

var SQLUserByEmail = `SELECT uid, email, password, security_email, two_factor, locked, locked_till, disabled FROM user_auth WHERE email = $1;`

func (pg *Postgres) UserAuthByEmail(email string) (*user.Auth, error) {
	au := &user.Auth{}
	err := pg.db.QueryRow(SQLUserByEmail, clean.Email(email)).
		Scan(&au.UID, &au.Email, &au.Password, &au.SecurityEmail, &au.TwoFactor,
			&au.Locked, &au.LockedTill, &au.Disabled)
	return au, err
}

var SQLUserBasicByUsername = `SELECT uid, username, admin FROM "user" WHERE lowercase = $1;`

func (pg *Postgres) UserBasicByUsername(username string) (*user.Basic, error) {
	bc := &user.BasicCore{}
	err := pg.db.QueryRow(SQLUserBasicByUsername, clean.Lowercase(username)).
		Scan(&bc.UID, &bc.Username, &bc.Admin)
	return user.NewBasicFromCore(bc), err
}

var SQLUserUIDByUsername = `SELECT uid FROM "user" WHERE lowercase = $1;`

func (pg *Postgres) UserUIDByUsername(username string) (int32, error) {
	var uid int32
	err := pg.db.QueryRow(SQLUserUIDByUsername, clean.Lowercase(username)).
		Scan(&uid)
	return uid, err
}

var SQLUserUsernameExists = `SELECT exists(SELECT uid FROM "user" WHERE lowercase = $1);`

func (pg *Postgres) UserUsernameExists(username string) (bool, error) {
	exists := true
	err := pg.db.QueryRow(SQLUserUsernameExists, clean.Lowercase(username)).
		Scan(&exists)
	return exists, err
}

var SQLUserEmailExists = `SELECT exists(SELECT uid FROM user_auth WHERE email = $1);`

func (pg *Postgres) UserEmailExists(email string) (bool, error) {
	exists := true
	err := pg.db.QueryRow(SQLUserEmailExists, email).
		Scan(&exists)
	return exists, err
}

var SQLUserUpdateUsername = `UPDATE "user" SET username = $2, lowercase = $3 WHERE uid = $1;`

func (pg *Postgres) UserUpdateUsername(uid int32, username string) error {
	_, err := pg.db.Exec(SQLUserUpdateUsername, uid, username, clean.Lowercase(username))
	return err
}

var SQLUserUpdateEmail = `UPDATE user_auth SET email = $2 WHERE uid = $1;`

func (pg *Postgres) UserUpdateEmail(uid int32, email string) error {
	_, err := pg.db.Exec(SQLUserUpdateEmail, uid, email)
	return err
}

var SQLUserUpdatePassword = `UPDATE user_auth SET password = $2 WHERE uid = $1;`

func (pg *Postgres) UserUpdatePassword(uid int32, hashed []byte) error {
	_, err := pg.db.Exec(SQLUserUpdatePassword, uid, hashed)
	return err
}

var SQLUserUpdateProfile = `UPDATE user_profile SET name = $2, bio = $3, location = $4, 
                        birthday = $5, email = $6, social = $7 WHERE uid = $1;`

func (pg *Postgres) UserUpdateProfile(uid int32, profile *user.Profile) error {
	_, err := pg.db.Exec(SQLUserUpdateProfile, uid, &profile.Name,
		&profile.Bio, &profile.Location, &profile.Birthday, &profile.Email, &profile.Social)
	return err
}

var SQLUserUpdateSecurityEmail = `UPDATE user_auth SET security_email = $2 WHERE uid = $1;`

func (pg *Postgres) UserUpdateSecurityEmail(uid int32, se sqlnull.String) error {
	_, err := pg.db.Exec(SQLUserUpdateSecurityEmail, uid, se)
	return err
}

var SQLUserUpdateLocked = `UPDATE user_auth SET locked = $2, locked_till = $3 WHERE uid = $1;`

func (pg *Postgres) UserUpdateLocked(uid int32, locked bool, till sqlnull.Time) error {
	_, err := pg.db.Exec(SQLUserUpdateLocked, uid, locked, till)
	return err
}

var SQLUserUpdateDisabled = `UPDATE user_auth SET disabled = $2 WHERE uid = $1;`

func (pg *Postgres) UserUpdateDisabled(uid int32, disabled bool) error {
	_, err := pg.db.Exec(SQLUserUpdateDisabled, uid, disabled)
	return err
}

var SQLUserUpdateElectrons = `UPDATE "user" SET electrons = $2 WHERE uid = $1;`

func (pg *Postgres) UserUpdateElectrons(uid int32, electrons int32) error {
	_, err := pg.db.Exec(SQLUserUpdateElectrons, uid, electrons)
	return err
}

var SQLUserUpdateDeltaElectrons = `UPDATE "user" SET electrons = electrons + $2 WHERE uid = $1;`

func (pg *Postgres) UserUpdateDeltaElectrons(uid int32, delta int32) error {
	_, err := pg.db.Exec(SQLUserUpdateDeltaElectrons, uid, delta)
	return err
}

var SQLUserUpdateAdmin = `UPDATE "user" SET admin = $2 WHERE uid = $1;`

func (pg *Postgres) UserUpdateAdmin(uid int32, admin int16) error {
	_, err := pg.db.Exec(SQLUserUpdateAdmin, uid, admin)
	return err
}
