package data

import (
	"github.com/go-ushio/ushio/common/put"
	"github.com/go-ushio/ushio/core/sign_up"
	"github.com/go-ushio/ushio/utils/clean"
)

func (data *Data) SignUpByToken(token string) (*sign_up.SignUp, error) {
	row := data.db.QueryRow(data.sentence.SQLSignUpByToken, token)
	return sign_up.ScanSignUp(row)
}

func (data *Data) SignUpByEmail(email string) (*sign_up.SignUp, error) {
	clean.String(email)
	row := data.db.QueryRow(data.sentence.SQLSignUpByEmail, email)
	return sign_up.ScanSignUp(row)
}

func (data *Data) InsertSignUp(su *sign_up.SignUp) error {
	putter := put.PutterFromDBExec(data.db, data.sentence.SQLInsertSignUp)
	_, err := su.Put(putter)
	return err
}

func (data *Data) DeleteSignUpByEmail(email string) error {
	clean.String(email)
	putter := put.PutterFromDBExec(data.db, data.sentence.SQLDeleteSignUpByEmail)
	_, err := putter.Put(email)
	return err
}

func (data *Data) SignUpExists(email string) (bool, error) {
	clean.String(email)
	row := data.db.QueryRow(data.sentence.SQLSignUpExists, email)
	e := true
	err := row.Scan(&e)
	if err != nil {
		return true, err
	}
	return e, nil
}
