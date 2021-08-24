package postgres

import (
	"github.com/go-uranium/uranium/model/session"
)

var SQLSessionBasicByToken = `SELECT token, uid, mode, expire FROM session WHERE token = $1;`

func (pg *Postgres) SessionBasicByToken(token string) (*session.Basic, error) {
	sb := &session.Basic{}
	err := pg.db.QueryRow(SQLSessionBasicByToken, token).
		Scan(&sb.Token, &sb.UID, &sb.Mode, &sb.Expire)
	return sb, err
}

var SQLSessionInsertSession = `INSERT INTO 
    session (token, uid, mode, ua, ip, created, expire)
 	VALUES ($1, $2, $3, $4, $5, $6, $7);`

func (pg *Postgres) SessionInsertSession(sess *session.Session) error {
	_, err := pg.db.Exec(SQLSessionInsertSession, sess.Token, sess.UID, sess.Mode, sess.UA,
		sess.IP, sess.Created, sess.Expire)
	return err
}
