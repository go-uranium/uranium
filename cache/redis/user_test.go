package rcache_test

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/pkg/errors"

	"github.com/go-uranium/uranium/model/user"
	"github.com/go-uranium/uranium/utils/sqlnull"
)

func TestRCache_UserBasicByUID(t *testing.T) {
	if err := Init(); err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if err := Close(); err != nil {
			t.Error(err)
			return
		}
	}()

	// ================== TEST 1 ======================
	basic, hit, err := cache.UserBasicByUID(0)
	if err == nil {
		t.Error("expected sql.ErrNoRows, got (nil)")
		return
	}
	if err != sql.ErrNoRows {
		t.Errorf(`expected sql.ErrNoRows, got "%s"`, err)
		return
	}
	if hit {
		t.Error("expected not hit, but hit")
	}

	// test again
	basic, hit, err = cache.UserBasicByUID(0)
	if err == nil {
		t.Error("expected sql.ErrNoRows, got (nil)")
		return
	}
	if err != sql.ErrNoRows {
		t.Errorf(`expected sql.ErrNoRows, got "%s"`, err)
		return
	}
	if hit {
		t.Error("expected not hit, but hit")
	}

	// ================== TEST 2 ======================

	basic1, _ := (&testingStorage{}).UserBasicByUID(1)

	// load from storage
	basic, hit, err = cache.UserBasicByUID(1)
	if err != nil {
		t.Error(err)
		return
	}
	if !basic.Equal(basic1) {
		t.Errorf("user not equal, want: %#v, got: %#v", *basic1.Core(), *basic.Core())
		return
	}
	if hit {
		t.Error("expected not hit, but hit")
	}

	// read from cache
	basic, hit, err = cache.UserBasicByUID(1)
	if err != nil {
		t.Error(err)
		return
	}
	if !basic.Equal(basic1) {
		t.Errorf("user not equal, want: %#v, got: %#v", *basic1, *basic)
		return
	}
	if !hit {
		t.Error("expected hit, but not hit")
	}

	// ================== TEST 3 ======================

	basic1, _ = (&testingStorage{}).UserBasicByUID(2)

	// load from storage
	basic, hit, err = cache.UserBasicByUID(2)
	if err != nil {
		t.Error(err)
		return
	}
	if !basic.Equal(basic1) {
		t.Errorf("user not equal, want: %#v, got: %#v", *basic1.Core(), *basic.Core())
		return
	}
	if hit {
		t.Error("expected not hit, but hit")
	}

	// read from cache
	basic, hit, err = cache.UserBasicByUID(2)
	if err != nil {
		t.Error(err)
		return
	}
	if !basic.Equal(basic1) {
		t.Errorf("user not equal, want: %#v, got: %#v", *basic1, *basic)
		return
	}
	if !hit {
		t.Error("expected hit, but not hit")
	}

	// ================== TEST 4 ======================

	// first time
	basic, hit, err = cache.UserBasicByUID(1024)
	if err == nil {
		t.Error("expected testingErr, got (nil)")
		return
	}
	if err != testingErr {
		t.Errorf(`expected testingErr, got "%s"`, err)
		return
	}
	if hit {
		t.Error("expected not hit, but hit")
	}

	// next time
	basic, hit, err = cache.UserBasicByUID(1024)
	if err == nil {
		t.Error("expected testingErr, got (nil)")
		return
	}
	if err != testingErr {
		t.Errorf(`expected testingErr, got "%s"`, err)
		return
	}
	if hit {
		t.Error("expected not hit, but hit")
	}
}

func TestRCache_UserUIDByUsername(t *testing.T) {
	if err := Init(); err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if err := Close(); err != nil {
			t.Error(err)
			return
		}
	}()

	// ================== TEST 1 ==================
	uid, hit, err := cache.UserUIDByUsername("invalid")
	if err == nil {
		t.Error("expected sql.ErrNoRows, got (nil)")
		return
	}
	if err != sql.ErrNoRows {
		t.Errorf(`expected sql.ErrNoRows, got "%s"`, err)
		return
	}
	if hit {
		t.Error("expected not hit, but hit")
	}

	// again
	uid, hit, err = cache.UserUIDByUsername("invalid")
	if err == nil {
		t.Error("expected sql.ErrNoRows, got (nil)")
		return
	}
	if err != sql.ErrNoRows {
		t.Errorf(`expected sql.ErrNoRows, got "%s"`, err)
		return
	}
	if hit {
		t.Error("expected not hit, but hit")
	}

	// ================== TEST 2 =====================
	// load from storage
	uid, hit, err = cache.UserUIDByUsername("richard")
	if err != nil {
		t.Error(err)
		return
	}
	if uid != 255 {
		t.Errorf("uid not equal, want: 255, got: %d", uid)
		return
	}
	if hit {
		t.Error("expected not hit, but hit")
	}
	// from cache
	uid, hit, err = cache.UserUIDByUsername("richard")
	if err != nil {
		t.Error(err)
		return
	}
	if uid != 255 {
		t.Errorf("uid not equal, want: 255, got: %d", uid)
		return
	}
	if !hit {
		t.Error("expected hit, but not hit")
	}
}

var testingErr = errors.New("this is an error for testing")

func (*testingStorage) UserBasicByUID(uid int32) (*user.Basic, error) {
	if uid == 0 {
		return &user.Basic{}, sql.ErrNoRows
	}
	if uid < 0 || uid == 1024 {
		return &user.Basic{}, testingErr
	}
	return user.NewBasicFromCore(&user.BasicCore{
		UID:      uid,
		Username: strconv.Itoa(int(uid)),
		Admin:    int16(uid % 3),
	}), nil
}

func (*testingStorage) UserUIDByUsername(username string) (int32, error) {
	if username == "invalid" {
		return 0, sql.ErrNoRows
	}
	return 255, nil
}

func (*testingStorage) UserInsertUser(u *user.User) (int32, error)                       { return 0, nil }
func (*testingStorage) UserInsertUserAuth(auth *user.Auth) error                         { return nil }
func (*testingStorage) UserInsertUserProfile(profile *user.Profile) error                { return nil }
func (*testingStorage) UserByUID(uid int32) (*user.User, error)                          { return nil, nil }
func (*testingStorage) UserProfileByUID(uid int32) (*user.Profile, error)                { return nil, nil }
func (*testingStorage) UserAuthByUID(uid int32) (*user.Auth, error)                      { return nil, nil }
func (*testingStorage) UserByUsername(username string) (*user.User, error)               { return nil, nil }
func (*testingStorage) UserByEmail(email string) (*user.User, error)                     { return nil, nil }
func (*testingStorage) UserBasicByUsername(username string) (*user.Basic, error)         { return nil, nil }
func (*testingStorage) UserUsernameExists(username string) (bool, error)                 { return true, nil }
func (*testingStorage) UserEmailExists(email string) (bool, error)                       { return true, nil }
func (*testingStorage) UserUpdateUsername(uid int32, username string) error              { return nil }
func (*testingStorage) UserUpdateEmail(uid int32, email string) error                    { return nil }
func (*testingStorage) UserUpdatePassword(uid int32, hashed []byte) error                { return nil }
func (*testingStorage) UserUpdateProfile(uid int32, profile *user.Profile) error         { return nil }
func (*testingStorage) UserUpdateSecurityEmail(uid int32, se sqlnull.String) error       { return nil }
func (*testingStorage) UserUpdateLocked(uid int32, locked bool, till sqlnull.Time) error { return nil }
func (*testingStorage) UserUpdateDisabled(uid int32, disabled bool) error                { return nil }
func (*testingStorage) UserUpdateElectrons(uid int32, electrons int32) error             { return nil }
func (*testingStorage) UserUpdateDeltaElectrons(uid int32, delta int32) error            { return nil }
func (*testingStorage) UserUpdateAdmin(uid int32, admin int16) error                     { return nil }
