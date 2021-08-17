package sign_up

import (
	"time"

	"github.com/go-uranium/uranium/utils/clean"
	"github.com/go-uranium/uranium/utils/token"
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
		Email:     clean.Email(email),
		Token:     token.New(),
		CreatedAt: now,
		ExpireAt:  now.Add(dur),
	}
}

func (su *SignUp) Valid() bool {
	return time.Now().Before(su.ExpireAt)
}
