package data_test

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-ushio/ushio/data"
	"github.com/go-ushio/ushio/user"
	"github.com/go-ushio/ushio/utils/hash"
)

func TestUserByUID(t *testing.T) {
	// Connect
	err := data.Init("mysql", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		t.Error(err)
	}
	defer data.Quit()

	// TEST 1
	user, err := data.UserByUID(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(user)

	// TEST 2
	_, err = data.UserByUID(0)
	if err != sql.ErrNoRows {
		t.Error("want: error(no rows), get nil")
	}
}

func TestUserByEmail(t *testing.T) {
	// Connect
	err := data.Init("mysql", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		t.Error(err)
	}
	defer data.Quit()

	// Test 1
	user, err := data.UserByEmail("i@iochen.com")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(user)

	// Test 2
	_, err = data.UserByEmail("not-exist@iochen.com")
	if err != sql.ErrNoRows {
		t.Error("want: error(no rows), get nil")
	}
}

func TestUserByUsername(t *testing.T) {
	// Connect
	err := data.Init("mysql", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		t.Error(err)
	}
	defer data.Quit()

	// Test 1
	user, err := data.UserByUsername("iochen")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(user)

	// Test 2
	_, err = data.UserByUsername("not-exist")
	if err != sql.ErrNoRows {
		t.Error("want: error(no rows), get nil")
	}
}

func TestInsertUser(t *testing.T) {
	// Connect
	err := data.Init("mysql", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		t.Error(err)
	}
	defer data.Quit()

	rand.Seed(time.Now().UnixNano())

	r := strconv.Itoa(rand.Intn(999999999))
	u := &user.User{
		Name:     "Test User" + r,
		Username: "t" + r,
		Password: hash.Hash(r),
		Email:    r + "@iochen.com",
	}

	err = data.InsertUser(u)
	if err != nil {
		t.Error(err)
	}

	uGot, err := data.UserByEmail(u.Email)
	if err != nil {
		t.Error(err)
	}

	if !uGot.Valid(r) {
		t.Error("cannot validate password")
	}

	if uGot.Valid("NOT VALID") {
		t.Error("cannot validate password")
	}

}
