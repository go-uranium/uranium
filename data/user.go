package data

import (
	"strings"

	"github.com/go-ushio/ushio/common/put"
	"github.com/go-ushio/ushio/core/user"
	"github.com/go-ushio/ushio/utils/clean"
)

func (data *Data) UserByUID(uid int) (*user.User, error) {
	row := data.db.QueryRow(data.sentence.SQLUserByUID, uid)
	u, err := user.ScanUser(row)
	if err != nil {
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
		return &user.User{}, err
	}
	u.Tidy()
	return u, nil
}

func (data *Data) UserAuthByUID(uid int) (*user.Auth, error) {
	row := data.db.QueryRow(data.sentence.SQLUserAuthByUID, uid)
	auth, err := user.ScanAuth(row)
	if err != nil {
		return &user.Auth{}, err
	}
	return auth, nil
}

func (data *Data) InsertUser(u *user.User) (int, error) {
	u.Tidy()
	putter := put.PutterFromDBExec(data.db, data.sentence.SQLInsertUser)
	result, err := u.Put(putter)
	if err != nil {
		return 0, err
	}
	uid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(uid), nil
}

func (data *Data) InsertUserAuth(auth *user.Auth) error {
	putter := put.PutterFromDBExec(data.db, data.sentence.SQLInsertUserAuth)
	_, err := auth.PutWithUIDFirst(putter)
	return err
}

func (data *Data) UpdateUser(u *user.User) error {
	u.Tidy()
	putter := put.PutterFromDBExec(data.db, data.sentence.SQLUpdateUser)
	_, err := u.PutWithUID(putter)
	return err
}

func (data *Data) UpdateUserAuth(auth *user.Auth) error {
	putter := put.PutterFromDBExec(data.db, data.sentence.SQLUpdateUserAuth)
	_, err := auth.PutWithUIDLast(putter)
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
