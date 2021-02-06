package sign_up

import (
	"database/sql"
	"time"

	"github.com/go-ushio/ushio/common/put"
	"github.com/go-ushio/ushio/common/scan"
	"github.com/go-ushio/ushio/utils/clean"
	"github.com/go-ushio/ushio/utils/token"
)

type SignUp struct {
	Token     string
	Email     string
	CreatedAt time.Time
	ExpireAt  time.Time
}

func New(email string, dur time.Duration) *SignUp {
	now := time.Now()
	return &SignUp{
		Email:     clean.String(email),
		Token:     token.New(),
		CreatedAt: now,
		ExpireAt:  now.Add(dur),
	}
}

func ScanSignUp(scanner scan.Scanner) (*SignUp, error) {
	su := &SignUp{}
	err := scanner.Scan(&su.Token, &su.Email, &su.CreatedAt, &su.ExpireAt)
	if err != nil {
		return &SignUp{}, err
	}
	return su, nil
}

func (su *SignUp) PutWithTokenFirst(putter put.Putter) (sql.Result, error) {
	return putter.Put(su.Token, su.Email, su.CreatedAt, su.ExpireAt)
}

func (su *SignUp) Valid() bool {
	return time.Now().Before(su.ExpireAt)
}
