package data

import (
	"strconv"

	"github.com/go-ushio/ushio/user"
)

var (
	SQLUserByUID      = `SELECT * FROM ushio.user WHERE uid = ?;`
	SQLUserByEmail    = `SELECT * FROM ushio.user WHERE email = ?;`
	SQLInsertUser     = `INSERT INTO ushio.user(name, username, email, password) VALUES (?,?,?,?);`
	SQLUserByUsername = `SELECT * FROM ushio.user WHERE username = ?;`
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
	_, err := db.Exec(SQLInsertUser, u.Name, u.Username,
		u.Email, u.Password)
	if err != nil {
		return err
	}
	return nil
}
