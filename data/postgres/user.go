package postgres

import (
	"database/sql"
	"strings"

	"github.com/lib/pq"

	"github.com/go-ushio/ushio/core/user"
	"github.com/go-ushio/ushio/utils/clean"
)

var (
	SQLUserByUID      = `SELECT uid, name, username, email, avatar, bio, created_at, artifact, following FROM ushio."user" WHERE uid = $1;`
	SQLUserByEmail    = `SELECT uid, name, username, email, avatar, bio, created_at, artifact, following FROM ushio."user" WHERE email = $1;`
	SQLUserByUsername = `SELECT uid, name, username, email, avatar, bio, created_at, artifact, following FROM ushio."user" WHERE username = $1;`
	SQLUserAuthByUID  = `SELECT uid, password, locked, security_email FROM ushio.user_auth WHERE uid = $1;`
	SQLInsertUser     = `INSERT INTO ushio.user(name, username, email, avatar, bio, created_at, artifact, following) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING uid;`
	SQLInsertUserAuth = `INSERT INTO ushio.user_auth(uid, password, locked, security_email) VALUES ($1, $2, $3, $4);`
	SQLUpdateUser     = `UPDATE ushio."user" SET name = $2, username = $3, email = $4, avatar = $5, bio = $6, created_at = $7, artifact = $8, following = $9 WHERE uid = $1;`
	SQLUpdateUserAuth = `UPDATE ushio.user_auth SET password = $2, locked = $3, security_email = $4 WHERE uid = $1;`
	SQLAddArtifact    = `UPDATE ushio."user" SET artifact = artifact + $2 WHERE uid = $1;`
	SQLDeleteUser     = `DELETE FROM ushio.user_auth WHERE uid = $1;DELETE FROM ushio.session WHERE uid = $1;DELETE FROM ushio."user" WHERE uid = $1;
UPDATE ushio.post_info SET creator = 0 WHERE creator = $1;UPDATE ushio.comment SET creator = 0 WHERE creator = $1;`
	SQLUsernameExists = `SELECT EXISTS(SELECT uid FROM ushio."user" WHERE username = $1);`
	SQLEmailExists    = `SELECT EXISTS(SELECT uid FROM ushio."user" WHERE email = $1);`
)

func (data *Data) UserByUID(uid int) (*user.User, error) {
	row := data.db.QueryRow(SQLUserByUID, uid)
	u, err := user.ScanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &user.User{}, err
	}
	u.Tidy()
	return u, nil
}

func (data *Data) UserByEmail(email string) (*user.User, error) {
	email = clean.String(email)
	row := data.db.QueryRow(SQLUserByEmail, email)
	u, err := user.ScanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &user.User{}, err
	}
	u.Tidy()
	return u, nil
}

func (data *Data) UserByUsername(username string) (*user.User, error) {
	username = clean.String(username)
	row := data.db.QueryRow(SQLUserByUsername, username)
	u, err := user.ScanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &user.User{}, err
	}
	u.Tidy()
	return u, nil
}

func (data *Data) UserAuthByUID(uid int) (*user.Auth, error) {
	row := data.db.QueryRow(SQLUserAuthByUID, uid)
	auth, err := user.ScanAuth(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &user.Auth{}, err
	}
	return auth, nil
}

func (data *Data) InsertUser(u *user.User) (int, error) {
	u.Tidy()
	uid := 0
	err := data.db.QueryRow(SQLInsertUser, u.Name,
		u.Username, u.Email, u.Avatar, u.Bio,
		u.CreatedAt, u.Artifact, pq.Array(u.Following)).Scan(&uid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

func (data *Data) InsertUserAuth(auth *user.Auth) error {
	_, err := data.db.Exec(SQLInsertUserAuth, auth.UID,
		auth.Password, auth.Locked, auth.SecurityEmail)
	return err
}

func (data *Data) UpdateUser(u *user.User) error {
	u.Tidy()
	_, err := data.db.Exec(SQLUpdateUser, u.UID, u.Name,
		u.Username, u.Email, u.Avatar, u.Bio, u.CreatedAt,
		u.Artifact, pq.Array(u.Following))
	return err
}

func (data *Data) UpdateUserAuth(auth *user.Auth) error {
	_, err := data.db.Exec(SQLUpdateUserAuth, auth.UID,
		auth.Password, auth.Locked, auth.SecurityEmail)
	return err
}

func (data *Data) AddArtifact(uid, add int) error {
	_, err := data.db.Exec(SQLAddArtifact, uid, add)
	return err
}

func (data *Data) DeleteUser(uid int) error {
	_, err := data.db.Exec(SQLDeleteUser, uid)
	return err
}

func (data *Data) UsernameExists(username string) (bool, error) {
	username = strings.ToLower(username)
	row := data.db.QueryRow(SQLUsernameExists, username)
	e := true
	err := row.Scan(&e)
	if err != nil {
		return true, err
	}
	return e, nil
}

func (data *Data) EmailExists(email string) (bool, error) {
	email = strings.ToLower(email)
	row := data.db.QueryRow(SQLEmailExists, email)
	e := true
	err := row.Scan(&e)
	if err != nil {
		return true, err
	}
	return e, nil
}
