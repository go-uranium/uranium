package data

import (
	"database/sql"

	"github.com/go-ushio/ushio/core/sign_up"
	"github.com/go-ushio/ushio/utils/clean"
)

func (data *Data) SignUpByToken(token string) (*sign_up.SignUp, error) {
	row := data.db.QueryRow(data.sentence.SQLSignUpByToken, token)
	signUp, err := sign_up.ScanSignUp(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &sign_up.SignUp{}, err
	}
	return signUp, nil
}

func (data *Data) SignUpByEmail(email string) (*sign_up.SignUp, error) {
	clean.String(email)
	row := data.db.QueryRow(data.sentence.SQLSignUpByEmail, email)
	signUp, err := sign_up.ScanSignUp(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &sign_up.SignUp{}, err
	}
	return signUp, nil
}

func (data *Data) InsertSignUp(su *sign_up.SignUp) error {
	_, err := data.db.Exec(data.sentence.SQLInsertSignUp, su.Token, su.Email, su.CreatedAt, su.ExpireAt)
	return err
}

func (data *Data) DeleteSignUpByEmail(email string) error {
	clean.String(email)
	_, err := data.db.Exec(data.sentence.SQLDeleteSignUpByEmail, email)
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
