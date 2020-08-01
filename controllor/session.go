package controllor

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	sid uuid.UUID
	LoggedIn bool
	User *User
	ExpireDate time.Time
}

