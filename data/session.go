package data

import (
	"database/sql"
	"strings"

	"github.com/go-ushio/ushio/session"
)

var (
	SQLSessionByToken = `SELECT token, uid, UA, IP, time, expire_at FROM ushio.session WHERE token=?;`
	SQLSessionsByUID  = `SELECT token, uid, UA, IP, time, expire_at FROM ushio.session WHERE uid=?;`

	SQLInsertSession   = `INSERT INTO ushio.session(token, uid, UA, IP, time, expire_at) VALUES (?,?,?,?,?,?);`
	SQLDisableSessions = `UPDATE ushio.session SET expire_at = NOW() WHERE uid=?;`
	SQLDeleteSessions  = `DELETE FROM ushio.session WHERE uid=?;`
)

func SessionByToken(token string) (*session.Session, error) {
	token = strings.ToLower(token)
	row := db.QueryRow(SQLSessionByToken, token)
	s := &session.Session{}
	err := row.Scan(&s.Token, &s.UID, &s.UA, &s.IP, &s.Time, &s.Expire)
	if err != nil {
		return &session.Session{}, err
	}
	return s, nil
}

func SessionsByUID(uid int) ([]*session.Session, error) {
	rows, err := db.Query(SQLSessionsByUID, uid)
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

func InsertSession(s *session.Session) error {
	_, err := db.Exec(SQLInsertSession, s.Token, s.UID, s.UA, s.IP, s.Time, s.Expire)
	if err != nil {
		return err
	}
	return nil
}

func DisableSessions(uid int) error {
	_, err := db.Exec(SQLDisableSessions, uid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSessions(uid int) error {
	_, err := db.Exec(SQLDeleteSessions, uid)
	if err != nil {
		return err
	}
	return nil
}
