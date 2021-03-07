package postgres

import (
	"database/sql"

	"github.com/go-ushio/ushio/core/session"
)

var (
	SQLSessionByToken      = `SELECT token, uid, ua, ip, created_at, expire_at FROM ushio.session WHERE token = $1;`
	SQLSessionsByUID       = `SELECT token, uid, ua, ip, created_at, expire_at FROM ushio.session WHERE uid = $1;`
	SQLSessionBasicByToken = `SELECT token, uid, expire_at FROM ushio.session WHERE token = $1;`
	SQLInsertSession       = `INSERT INTO ushio.session(token, uid, ua, ip, created_at, expire_at) VALUES ($1, $2, $3, $4, $5, $6);`
	SQLDeleteUserSessions  = `DELETE FROM ushio.session WHERE uid = $1;`
)

func (pg *Postgres) SessionByToken(token string) (*session.Session, error) {
	row := pg.db.QueryRow(SQLSessionByToken, token)
	sess, err := session.ScanSession(row)
	if err != nil {
		//if err == sql.ErrNoRows {
		//	return nil, nil
		//}
		return &session.Session{}, err
	}
	return sess, nil
}

func (pg *Postgres) SessionsByUID(uid int64) ([]*session.Session, error) {
	rows, err := pg.db.Query(SQLSessionsByUID, uid)
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

func (pg *Postgres) SessionBasicByToken(token string) (*session.Basic, error) {
	row := pg.db.QueryRow(SQLSessionBasicByToken, token)
	bsc, err := session.ScanBasic(row)
	if err != nil {
		//if err == sql.ErrNoRows {
		//	return nil, nil
		//}
		return &session.Basic{}, err
	}
	return bsc, nil
}

func (pg *Postgres) InsertSession(sess *session.Session) error {
	_, err := pg.db.Exec(SQLInsertSession, sess.Token, sess.UID, sess.UA,
		sess.IP, sess.CreatedAt, sess.ExpireAt)
	return err
}

func (pg *Postgres) DeleteUserSessions(uid int64) error {
	_, err := pg.db.Exec(SQLDeleteUserSessions, uid)
	return err
}
