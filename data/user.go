package data

import (
	"strconv"
	"strings"

	"github.com/go-ushio/ushio/user"
)

var (
	SQLUserByUID      = `SELECT uid, name, username, email, password, created_at, is_admin, banned, locked, flags FROM ushio.user WHERE uid = ?;`
	SQLUserByEmail    = `SELECT uid, name, username, email, password, created_at, is_admin, banned, locked, flags FROM ushio.user WHERE email = ?;`
	SQLUserByUsername = `SELECT uid, name, username, email, password, created_at, is_admin, banned, locked, flags FROM ushio.user WHERE username = ?;`

	SQLInsertUser = `INSERT INTO ushio.user(name, username, email, password) VALUES (?,?,?,?);`

	SQLUsernameExists = `SELECT EXISTS(SELECT * FROM ushio.user WHERE username=?);`
	SQLEmailExists    = `SELECT EXISTS(SELECT * FROM ushio.user WHERE email=?);`
)

func UserByUID(uid int) (*user.User, error) {
	row := db.QueryRow(SQLUserByUID, strconv.Itoa(uid))
	u := &user.User{}
	err := row.Scan(&u.UID, &u.Name, &u.Username, &u.Email,
		&u.Password, &u.CreatedAt, &u.IsAdmin,
		&u.Banned, &u.Locked, &u.Flags)
	if err != nil {
		return &user.User{}, err
	}
	return u, nil
}

func UserByEmail(email string) (*user.User, error) {
	email = strings.ToLower(email)
	row := db.QueryRow(SQLUserByEmail, email)
	u := &user.User{}
	err := row.Scan(&u.UID, &u.Name, &u.Username, &u.Email,
		&u.Password, &u.CreatedAt, &u.IsAdmin,
		&u.Banned, &u.Locked, &u.Flags)
	if err != nil {
		return &user.User{}, err
	}
	return u, nil
}

func UserByUsername(username string) (*user.User, error) {
	username = strings.ToLower(username)
	row := db.QueryRow(SQLUserByUsername, username)
	u := &user.User{}
	err := row.Scan(&u.UID, &u.Name, &u.Username, &u.Email,
		&u.Password, &u.CreatedAt, &u.IsAdmin,
		&u.Banned, &u.Locked, &u.Flags)
	if err != nil {
		return &user.User{}, err
	}
	return u, nil
}

func InsertUser(u *user.User) error {
	u.Tidy()
	_, err := db.Exec(SQLInsertUser, u.Name, u.Username,
		u.Email, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func UsernameExists(username string) (bool, error) {
	username = strings.ToLower(username)
	row := db.QueryRow(SQLUsernameExists, username)
	e := true
	err := row.Scan(&e)
	if err != nil {
		return true, err
	}
	return e, nil
}

func EmailExists(email string) (bool, error) {
	email = strings.ToLower(email)
	row := db.QueryRow(SQLEmailExists, email)
	e := true
	err := row.Scan(&e)
	if err != nil {
		return true, err
	}
	return e, nil
}
