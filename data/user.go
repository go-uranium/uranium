package data

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-ushio/ushio/user"
)

var (
	SQLUserByUID      = `SELECT uid, name, username, email, password, created_at, is_admin, banned, locked, flags FROM ushio.user WHERE uid = ?;`
	SQLUserByEmail    = `SELECT uid, name, username, email, password, created_at, is_admin, banned, locked, flags FROM ushio.user WHERE email = ?;`
	SQLUserByUsername = `SELECT uid, name, username, email, password, created_at, is_admin, banned, locked, flags FROM ushio.user WHERE username = ?;`

	SQLInsertUser = `INSERT INTO ushio.user(name, username, email, password) VALUES (?,?,?,?);`
	SQLUpdateUser = `UPDATE ushio.user SET name=?, username=?, email=?, password=?, is_admin=?, banned=?, locked=?, flags=? WHERE uid=?;`
	SQLDeleteUser = `DELETE FROM ushio.user WHERE uid=?;`

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

func UpdateUser(u *user.User) error {
	u.Tidy()
	fmt.Println(u)
	_, err := db.Exec(SQLUpdateUser, u.Name, u.Username,
		u.Email, u.Password, u.IsAdmin, u.Banned, u.Locked, u.Flags, u.UID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(uid int) error {
	_, err := db.Exec(SQLDeleteUser, uid)
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
