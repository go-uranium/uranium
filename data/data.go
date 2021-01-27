package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Sentence struct {
	SQLPostByPID      string
	SQLMarkdownByPID  string
	SQLPostInfoByPage string
	SQLInsertPost     string

	SQLSessionByToken  string
	SQLSessionsByUID   string
	SQLInsertSession   string
	SQLDisableSessions string
	SQLDeleteSessions  string

	SQLUserByUID      string
	SQLUserByEmail    string
	SQLUserByUsername string
	SQLInsertUser     string
	SQLUpdateUser     string
	SQLDeleteUser     string
	SQLUsernameExists string
	SQLEmailExists    string
}

type Data struct {
	db       *sql.DB
	sentence Sentence
}

func New(db *sql.DB, sentence Sentence) *Data {
	return &Data{
		db:       db,
		sentence: sentence,
	}
}

func MySQLSentence() Sentence {
	return Sentence{
		SQLPostByPID:      `SELECT pid, title, creator, content, created_at, last_mod, hidden, anonymous FROM ushio.post WHERE pid = ?;`,
		SQLMarkdownByPID:  `SELECT md_raw FROM ushio.post WHERE pid = ?;`,
		SQLPostInfoByPage: `SELECT pid, title, creator, created_at, last_mod, hidden, anonymous FROM ushio.post ORDER BY pid DESC LIMIT ?,?;`,
		SQLInsertPost:     `INSERT INTO ushio.post(title, creator, content, hidden, anonymous, md_raw) VALUES (?,?,?,?,?,?);`,

		SQLSessionByToken:  `SELECT token, uid, UA, IP, time, expire_at FROM ushio.session WHERE token=?;`,
		SQLSessionsByUID:   `SELECT token, uid, UA, IP, time, expire_at FROM ushio.session WHERE uid=?;`,
		SQLInsertSession:   `INSERT INTO ushio.session(token, uid, UA, IP, time, expire_at) VALUES (?,?,?,?,?,?);`,
		SQLDisableSessions: `UPDATE ushio.session SET expire_at = NOW() WHERE uid=?;`,
		SQLDeleteSessions:  `DELETE FROM ushio.session WHERE uid=?;`,

		SQLUserByUID:      `SELECT uid, name, username, email, password, created_at, is_admin, banned, locked, flags FROM ushio.user WHERE uid = ?;`,
		SQLUserByEmail:    `SELECT uid, name, username, email, password, created_at, is_admin, banned, locked, flags FROM ushio.user WHERE email = ?;`,
		SQLUserByUsername: `SELECT uid, name, username, email, password, created_at, is_admin, banned, locked, flags FROM ushio.user WHERE username = ?;`,
		SQLInsertUser:     `INSERT INTO ushio.user(name, username, email, password) VALUES (?,?,?,?);`,
		SQLUpdateUser:     `UPDATE ushio.user SET name=?, username=?, email=?, password=?, is_admin=?, banned=?, locked=?, flags=? WHERE uid=?;`,
		SQLDeleteUser:     `DELETE FROM ushio.user WHERE uid=?;`,
		SQLUsernameExists: `SELECT EXISTS(SELECT * FROM ushio.user WHERE username=?);`,
		SQLEmailExists:    `SELECT EXISTS(SELECT * FROM ushio.user WHERE email=?);`,
	}
}
