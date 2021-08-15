package session

import (
	"time"

	"github.com/go-uranium/uranium/utils/scan"
)

type Session struct {
	Token string
	UID   int64
	UA    string
	IP    string
	// if is sudo mode
	Sudo      bool
	CreatedAt time.Time
	ExpireAt  time.Time
}

type Basic struct {
	Token    string
	UID      int64
	Sudo     bool
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

func (sess *Session) Valid() bool {
	return time.Now().Before(sess.ExpireAt)
}

func (bsc *Basic) Valid() bool {
	return time.Now().Before(bsc.ExpireAt)
}
