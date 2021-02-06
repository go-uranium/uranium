package data

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Sentence struct {
	SQLPostByPID     string
	SQLPostInfoByPID string
	// SQLPostInfoByPage ignores post with `hidden` true
	SQLPostInfoByPage  string
	SQLInsertPost      string
	SQLInsertPostInfo  string
	SQLUpdatePost      string
	SQLUpdatePostInfo  string
	SQLPostNewReply    string
	SQLPostNewView     string
	SQLPostNewActivity string

	SQLSessionByToken      string
	SQLSessionsByUID       string
	SQLSessionBasicByToken string
	SQLInsertSession       string
	SQLDeleteUserSessions  string

	SQLUserByUID      string
	SQLUserByEmail    string
	SQLUserByUsername string
	SQLUserAuthByUID  string
	SQLInsertUser     string
	SQLInsertUserAuth string
	SQLUpdateUser     string
	SQLUpdateUserAuth string
	// SQLDeleteUser also deletes session and user_auth at the same time
	SQLDeleteUser     string
	SQLUsernameExists string
	SQLEmailExists    string

	SQLInsertMessage     string
	SQLMessageByMID      string
	SQLMessageByReceiver string
	SQLMessageBySender   string
	SQLMessageHasRead    string
	SQLDeleteMessage     string

	SQLSignUpByToken       string
	SQLSignUpByEmail       string
	SQLInsertSignUp        string
	SQLDeleteSignUpByEmail string
	SQLSignUpExists        string
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

func SQLSentence() Sentence {
	return Sentence{
		SQLPostByPID: `SELECT pid, content, markdown FROM ushio.post WHERE pid = ?;`,
		SQLPostInfoByPID: `SELECT pid, title, creator, created_at, last_mod, replies, 
views, activity, hidden, anonymous FROM ushio.post_info WHERE pid=? AND hidden=0;`,
		SQLPostInfoByPage: `SELECT pid, title, creator, created_at, last_mod, replies, 
views, activity, hidden, anonymous FROM ushio.post_info ORDER BY pid DESC LIMIT $2 OFFSET $1;`,
		SQLInsertPost: `INSERT INTO ushio.post(pid, content, markdown) VALUES (?,?,?);`,
		SQLInsertPostInfo: `INSERT INTO ushio.post_info(title, creator, created_at, 
last_mod, replies, views, activity, hidden, anonymous) VALUES (?,?,?,?,?,?,?,?,?);`,
		SQLUpdatePost:      ``,
		SQLUpdatePostInfo:  ``,
		SQLPostNewReply:    `UPDATE ushio.post_info SET replies=replies+1 WHERE pid=?;`,
		SQLPostNewView:     `UPDATE ushio.post_info SET views=views+1 WHERE pid=?;`,
		SQLPostNewActivity: `UPDATE ushio.post_info SET activity=current_timestamp() WHERE pid=?;`,

		SQLSessionByToken: `SELECT token, uid, ua, ip, created_at, expire_at 
FROM ushio.session WHERE token=$1;`,
		SQLSessionsByUID: `SELECT token, uid, ua, ip, created_at, expire_at 
FROM ushio.session WHERE uid=$1;`,
		SQLSessionBasicByToken: `SELECT token, uid, expire_at FROM ushio.session WHERE token=$1;`,
		SQLInsertSession: `INSERT INTO ushio.session(token, uid, ua, ip, created_at, expire_at) 
VALUES (?,?,?,?,?,?);`,
		SQLDeleteUserSessions: `DELETE FROM ushio.session WHERE uid=?;`,

		SQLUserByUID: `SELECT uid, name, username, email, avatar, bio, created_at, 
is_admin, banned, flags FROM ushio.user WHERE uid=?;`,
		SQLUserByEmail: `SELECT uid, name, username, email, avatar, bio, created_at, 
is_admin, banned, flags FROM ushio.user WHERE email=?;`,
		SQLUserByUsername: `SELECT uid, name, username, email, avatar, bio, created_at, 
is_admin, banned, flags FROM ushio.user WHERE username=?;`,
		SQLUserAuthByUID:  `SELECT uid, password, locked, security_email FROM ushio.user_auth WHERE uid=?;`,
		SQLInsertUser:     `INSERT INTO ushio.user(name, username, email, avatar, bio, created_at, is_admin, banned, flags) VALUES (?,?,?,?,?,?,?,?,?);`,
		SQLInsertUserAuth: `INSERT INTO ushio.user_auth(uid, password, locked, security_email) VALUES (?,?,?,?);`,
		SQLUpdateUser:     `UPDATE ushio.user SET name=?, username=?, email=?, avatar=?, bio=?, created_at=?, is_admin=?, banned=?, flags=? WHERE uid=?;`,
		SQLUpdateUserAuth: `UPDATE ushio.user_auth SET password=?, locked=?, security_email=? WHERE uid=?;`,
		SQLDeleteUser: `DELETE ushio.user, ushio.user_auth, ushio.session FROM ushio.user 
INNER JOIN ushio.user_auth, ushio.session 
WHERE ushio.user_auth.uid=ushio.session.uid  and ushio.user.uid=?;`,
		SQLUsernameExists: `SELECT EXISTS(SELECT uid FROM ushio.user WHERE username=?);`,
		SQLEmailExists:    `SELECT EXISTS(SELECT uid FROM ushio.user WHERE email=?);`,

		SQLInsertMessage:     ``,
		SQLMessageByMID:      ``,
		SQLMessageByReceiver: ``,
		SQLMessageBySender:   ``,
		SQLMessageHasRead:    ``,
		SQLDeleteMessage:     ``,

		SQLSignUpByToken:       `SELECT token, email, created_at, expire_at FROM ushio.sign_up WHERE token=?;`,
		SQLSignUpByEmail:       `SELECT token, email, created_at, expire_at FROM ushio.sign_up WHERE email=?;`,
		SQLInsertSignUp:        `INSERT INTO ushio.sign_up(token, email, created_at, expire_at) VALUES (?,?,?,?);`,
		SQLDeleteSignUpByEmail: `DELETE FROM ushio.sign_up WHERE email=?;`,
		SQLSignUpExists:        `SELECT EXISTS(SELECT token FROM ushio.sign_up WHERE email=?);`,
	}
}
