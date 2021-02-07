package session

import (
	"database/sql"
	"time"

	"github.com/go-ushio/ushio/common/put"
	"github.com/go-ushio/ushio/common/scan"
)

type Session struct {
	Token     string
	UID       int
	UA        string
	IP        string
	CreatedAt time.Time
	ExpireAt  time.Time
}

type Basic struct {
	Token    string
	UID      int
	ExpireAt time.Time
}

func ScanSession(scanner scan.Scanner) (*Session, error) {
	sess := &Session{}
	err := scanner.Scan(&sess.Token, &sess.UID, &sess.UA,
		&sess.IP, &sess.CreatedAt, &sess.ExpireAt)
	if err != nil {
		return &Session{}, err
	}
	return sess, nil
}

func ScanBasic(scanner scan.Scanner) (*Basic, error) {
	bsc := &Basic{}
	err := scanner.Scan(&bsc.Token, &bsc.UID, &bsc.ExpireAt)
	if err != nil {
		return &Basic{}, err
	}
	return bsc, nil
}

func (sess *Session) Put(putter put.Putter) (sql.Result, error) {
	return putter.Put(sess.Token, sess.UID, sess.UA,
		sess.IP, sess.CreatedAt, sess.ExpireAt)
}

func (bsc *Basic) Put(putter put.Putter) (sql.Result, error) {
	return putter.Put(bsc.Token, bsc.UID, bsc.ExpireAt)
}

func (sess *Session) Valid() bool {
	return time.Now().Before(sess.ExpireAt)
}

func (bsc *Basic) Valid() bool {
	return time.Now().Before(bsc.ExpireAt)
}
