package controllor

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

)

type User struct {
	UID       uint      `json:"uid"`
	Username  string    `json:"username"`
	Email     string    `json:"-"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	Weight    int       `json:"-"`
	Mod       *Mod      `json:"mod"`
}

func QueryUser(queryStr string) (*User, error) {
	user := &User{}
	var row *sql.Row
	i, err := strconv.Atoi(queryStr)
	if err == nil {
		row = DB.QueryRow(`select uid,username,email,
       password,name,avatar,created_at,weight,mod_of
       from ushio.user where uid=?`, i)
	} else {
		row = DB.QueryRow(`select uid,username,email,
       password,name,avatar,created_at,weight,mod_of
       from ushio.user where username=?`, queryStr)
	}
	var mod string
	if err := row.Scan(&user.UID,&user.Username, &user.Email, &user.Password, &user.Name,
		&user.Avatar, &user.CreatedAt, &user.Weight, &mod); err != nil {
		return &User{}, err
	}
	user.Mod, err = NewMod(mod)
	return user, err
}

func InsertUser(user *User) (string,error) {
	if len(user.Username) >15 {
		return "username too long", nil
	}
	if len(user.Email) >50 {
		return "email too long",nil
	}
	if len(user.Password) != 60 {
		return "",errors.New("not valid hashed password")
	}
	if len(user.Name) >20 {
		return "name too long",nil
	}
	if user.Name == "" {
		user.Name = strings.Title(user.Username)
	}
	if len(user.Avatar) > 128 {
		return "",errors.New("avatar too long")
	}

}