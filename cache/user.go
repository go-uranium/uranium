package cache

import (
	"github.com/go-ushio/ushio/data"
	"github.com/go-ushio/ushio/user"
)

var userByUID = map[int]*user.SimpleUser{}

var userByUsername = map[string]*user.SimpleUser{}

func UserByUID(uid int) (*user.SimpleUser, error) {
	v, ok := userByUID[uid]
	if ok {
		return v, nil
	}

	u, err := data.UserByUID(uid)
	if err != nil {
		return &user.SimpleUser{}, err
	}

	su := &user.SimpleUser{
		UID:      u.UID,
		Name:     u.Name,
		Username: u.Username,
	}
	userByUID[u.UID] = su
	return su, nil
}

func UserByUsername(username string) (*user.SimpleUser, error) {
	v, ok := userByUsername[username]
	if ok {
		return v, nil
	}

	u, err := data.UserByUsername(username)
	if err != nil {
		return &user.SimpleUser{}, err
	}

	su := &user.SimpleUser{
		UID:      u.UID,
		Name:     u.Name,
		Username: u.Username,
	}
	userByUsername[u.Username] = su
	return su, nil
}
