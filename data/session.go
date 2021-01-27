package data

import (
	"database/sql"
	"strings"

	"github.com/go-ushio/ushio/core/session"
)

func (data *Data) SessionByToken(token string) (*session.Session, error) {
	token = strings.ToLower(token)
	row := data.db.QueryRow(data.sentence.SQLSessionByToken, token)
	s := &session.Session{}
	err := row.Scan(&s.Token, &s.UID, &s.UA, &s.IP, &s.Time, &s.Expire)
	if err != nil {
		return &session.Session{}, err
	}
	return s, nil
}

func (data *Data) SessionsByUID(uid int) ([]*session.Session, error) {
	rows, err := data.db.Query(data.sentence.SQLSessionsByUID, uid)
	if err != nil {
		return nil, err
	}
	var ss []*session.Session
	for rows.Next() {
		s := &session.Session{}
		err := rows.Scan(&s.Token, &s.UID, &s.UA, &s.IP, &s.Time, &s.Expire)
		if err != nil {
			return nil, err
		}
		ss = append(ss, s)
	}
	if len(ss) == 0 {
		return nil, sql.ErrNoRows
	}
	return ss, nil
}

func (data *Data) InsertSession(s *session.Session) error {
	_, err := data.db.Exec(data.sentence.SQLInsertSession, s.Token, s.UID, s.UA, s.IP, s.Time, s.Expire)
	if err != nil {
		return err
	}
	return nil
}

func (data *Data) DisableSessions(uid int) error {
	_, err := data.db.Exec(data.sentence.SQLDisableSessions, uid)
	if err != nil {
		return err
	}
	return nil
}

func (data *Data) DeleteSessions(uid int) error {
	_, err := data.db.Exec(data.sentence.SQLDeleteSessions, uid)
	if err != nil {
		return err
	}
	return nil
}
