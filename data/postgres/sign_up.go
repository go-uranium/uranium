package postgres

import (
	"github.com/go-ushio/ushio/model/sign_up"
	"github.com/go-ushio/ushio/utils/clean"
)

var (
	SQLSignUpByToken       = `SELECT token, email, created_at, expire_at FROM sign_up WHERE token = $1;`
	SQLSignUpByEmail       = `SELECT token, email, created_at, expire_at FROM sign_up WHERE email = $1;`
	SQLInsertSignUp        = `INSERT INTO sign_up(token, email, created_at, expire_at) VALUES ($1, $2, $3, $4);`
	SQLDeleteSignUpByEmail = `DELETE FROM sign_up WHERE email = $1;`
	SQLSignUpExists        = `SELECT EXISTS(SELECT token FROM sign_up WHERE email = $1);`
)

func (pg *Postgres) SignUpByToken(token string) (*sign_up.SignUp, error) {
	row := pg.db.QueryRow(SQLSignUpByToken, token)
	signUp, err := sign_up.ScanSignUp(row)
	if err != nil {
		//if err == sql.ErrNoRows {
		//	return nil, nil
		//}
		return &sign_up.SignUp{}, err
	}
	return signUp, nil
}

func (pg *Postgres) SignUpByEmail(email string) (*sign_up.SignUp, error) {
	clean.String(email)
	row := pg.db.QueryRow(SQLSignUpByEmail, email)
	signUp, err := sign_up.ScanSignUp(row)
	if err != nil {
		//if err == sql.ErrNoRows {
		//	return nil, nil
		//}
		return &sign_up.SignUp{}, err
	}
	return signUp, nil
}

func (pg *Postgres) InsertSignUp(su *sign_up.SignUp) error {
	_, err := pg.db.Exec(SQLInsertSignUp, su.Token, su.Email, su.CreatedAt, su.ExpireAt)
	return err
}

func (pg *Postgres) DeleteSignUpByEmail(email string) error {
	clean.String(email)
	_, err := pg.db.Exec(SQLDeleteSignUpByEmail, email)
	return err
}

func (pg *Postgres) SignUpExists(email string) (bool, error) {
	clean.String(email)
	row := pg.db.QueryRow(SQLSignUpExists, email)
	e := true
	err := row.Scan(&e)
	if err != nil {
		return true, err
	}
	return e, nil
}
