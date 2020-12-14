package data_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/go-ushio/ushio/data"
	"github.com/go-ushio/ushio/session"
)

func TestInsertSession(t *testing.T) {
	// Connect
	err := data.Init("mysql", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		t.Error(err)
	}
	defer data.Quit()

	s := &session.Session{
		UID:    0,
		Token:  uuid.New().String(),
		UA:     "TEST UA",
		IP:     "0.0.0.0",
		Time:   time.Now(),
		Expire: time.Now().Add(720 * time.Hour),
	}

	err = data.InsertSession(s)
	if err != nil {
		t.Error(err)
	}
}

func TestSessionByUID(t *testing.T) {
	// Connect
	err := data.Init("mysql", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		t.Error(err)
	}
	defer data.Quit()

	// Test 1
	user, err := data.SessionByUID(0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(user)

	// Test 2
	_, err = data.SessionByUID(-1)
	if err != sql.ErrNoRows {
		t.Error("want: error(no rows), get nil")
	}
}

func TestSessionByToken(t *testing.T) {
	// Connect
	err := data.Init("mysql", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		t.Error(err)
	}
	defer data.Quit()

	// Test 1
	user, err := data.SessionByToken("9bd22785-4ba6-468e-9087-c98d79e67133")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(user)

	// Test 2
	_, err = data.SessionByToken(uuid.New().String())
	if err != sql.ErrNoRows {
		t.Error("want: error(no rows), get nil")
	}
}
