package data

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-ushio/ushio/core/user"
)

func (data *Data) UserByUID(uid int) (*user.User, error) {
	row := data.db.QueryRow(data.sentence.SQLUserByUID, strconv.Itoa(uid))
	u := &user.User{}
	err := row.Scan(&u.UID, &u.Name, &u.Username, &u.Email,
		&u.Password, &u.CreatedAt, &u.IsAdmin,
		&u.Banned, &u.Locked, &u.Flags)
	if err != nil {
		return &user.User{}, err
	}
	u.HashEmail()
	return u, nil
}

func (data *Data) UserByEmail(email string) (*user.User, error) {
	email = strings.ToLower(email)
	row := data.db.QueryRow(data.sentence.SQLUserByEmail, email)
	u := &user.User{}
	err := row.Scan(&u.UID, &u.Name, &u.Username, &u.Email,
		&u.Password, &u.CreatedAt, &u.IsAdmin,
		&u.Banned, &u.Locked, &u.Flags)
	if err != nil {
		return &user.User{}, err
	}
	u.HashEmail()
	return u, nil
}

func (data *Data) UserByUsername(username string) (*user.User, error) {
	username = strings.ToLower(username)
	row := data.db.QueryRow(data.sentence.SQLUserByUsername, username)
	u := &user.User{}
	err := row.Scan(&u.UID, &u.Name, &u.Username, &u.Email,
		&u.Password, &u.CreatedAt, &u.IsAdmin,
		&u.Banned, &u.Locked, &u.Flags)
	if err != nil {
		return &user.User{}, err
	}
	u.HashEmail()
	return u, nil
}

func (data *Data) InsertUser(u *user.User) error {
	u.Tidy()
	_, err := data.db.Exec(data.sentence.SQLInsertUser, u.Name, u.Username,
		u.Email, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (data *Data) UpdateUser(u *user.User) error {
	u.Tidy()
	fmt.Println(u)
	_, err := data.db.Exec(data.sentence.SQLUpdateUser, u.Name, u.Username,
		u.Email, u.Password, u.IsAdmin, u.Banned, u.Locked, u.Flags, u.UID)
	if err != nil {
		return err
	}
	return nil
}

func (data *Data) DeleteUser(uid int) error {
	_, err := data.db.Exec(data.sentence.SQLDeleteUser, uid)
	if err != nil {
		return err
	}
	return nil
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
