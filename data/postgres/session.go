package postgres

import (
	"database/sql"

	"github.com/go-ushio/ushio/core/session"
)

func (data *Data) SessionByToken(token string) (*session.Session, error) {
	row := data.db.QueryRow(data.sentence.SQLSessionByToken, token)
	sess, err := session.ScanSession(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &session.Session{}, err
	}
	return sess, nil
}

func (data *Data) SessionsByUID(uid int) ([]*session.Session, error) {
	rows, err := data.db.Query(data.sentence.SQLSessionsByUID, uid)
	if err != nil {
		return nil, err
	}
	var ss []*session.Session
	for rows.Next() {
		sess, err := session.ScanSession(rows)
		if err != nil {
			return nil, err
		}
		ss = append(ss, sess)
	}
	if len(ss) == 0 {
		return nil, sql.ErrNoRows
	}
	return ss, nil
}

func (data *Data) SessionBasicByToken(token string) (*session.Basic, error) {
	row := data.db.QueryRow(data.sentence.SQLSessionBasicByToken, token)
	bsc, err := session.ScanBasic(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &session.Basic{}, err
	}
	return bsc, nil
}

func (data *Data) InsertSession(sess *session.Session) error {
	_, err := data.db.Exec(data.sentence.SQLInsertSession, sess.Token, sess.UID, sess.UA,
		sess.IP, sess.CreatedAt, sess.ExpireAt)
	return err
}

func (data *Data) DeleteUserSessions(uid int) error {
	_, err := data.db.Exec(data.sentence.SQLDeleteUserSessions, uid)
	return err
}
