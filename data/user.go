package data

import (
	"database/sql"
	"strings"

	"github.com/go-ushio/ushio/core/user"
	"github.com/go-ushio/ushio/utils/clean"
)

func (data *Data) UserByUID(uid int) (*user.User, error) {
	row := data.db.QueryRow(data.sentence.SQLUserByUID, uid)
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
	row := data.db.QueryRow(data.sentence.SQLUserByEmail, email)
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
	row := data.db.QueryRow(data.sentence.SQLUserByUsername, username)
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
	row := data.db.QueryRow(data.sentence.SQLUserAuthByUID, uid)
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
	err := data.db.QueryRow(data.sentence.SQLInsertUser, u.Name,
		u.Username, u.Email, u.Avatar, u.Bio, u.CreatedAt,
		u.IsAdmin, u.Banned, u.Artifact).Scan(&uid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

func (data *Data) InsertUserAuth(auth *user.Auth) error {
	_, err := data.db.Exec(data.sentence.SQLInsertUserAuth, auth.UID,
		auth.Password, auth.Locked, auth.SecurityEmail)
	return err
}

func (data *Data) UpdateUser(u *user.User) error {
	u.Tidy()
	_, err := data.db.Exec(data.sentence.SQLUpdateUser, u.UID, u.Name,
		u.Username, u.Email, u.Avatar, u.Bio, u.CreatedAt,
		u.IsAdmin, u.Banned, u.Artifact)
	return err
}

func (data *Data) UpdateUserAuth(auth *user.Auth) error {
	_, err := data.db.Exec(data.sentence.SQLUpdateUserAuth, auth.UID,
		auth.Password, auth.Locked, auth.SecurityEmail)
	return err
}

func (data *Data) AddArtifact(uid, add int) error {
	_, err := data.db.Exec(data.sentence.SQLAddArtifact, uid, add)
	return err
}

func (data *Data) DeleteUser(uid int) error {
	_, err := data.db.Exec(data.sentence.SQLDeleteUser, uid)
	return err
}

func (data *Data) UsernameExists(username string) (bool, error) {
	username = strings.ToLower(username)
	row := data.db.QueryRow(data.sentence.SQLUsernameExists, username)
	e := true
	err := row.Scan(&e)
	if err != nil {
		return true, err
	}
	return e, nil
}

func (data *Data) EmailExists(email string) (bool, error) {
	email = strings.ToLower(email)
	row := data.db.QueryRow(data.sentence.SQLEmailExists, email)
	e := true
	err := row.Scan(&e)
	if err != nil {
		return true, err
	}
	return e, nil
}
