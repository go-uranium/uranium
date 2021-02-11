package data

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/go-ushio/ushio/core/post"
	"github.com/go-ushio/ushio/core/session"
	"github.com/go-ushio/ushio/core/sign_up"
	"github.com/go-ushio/ushio/core/user"
)

type Sentence struct {
	SQLPostByPID       string
	SQLPostInfoByPID   string
	SQLPostInfoByPage  string
	SQLPostInfoIndex   string
	SQLInsertPost      string
	SQLInsertPostInfo  string
	SQLUpdatePost      string
	SQLUpdatePostTitle string
	SQLUpdatePostLimit string
	SQLPostNewReply    string
	SQLPostNewView     string
	SQLPostNewMod      string
	SQLPostNewActivity string
	SQLPostNewPosVote  string
	SQLPostNewNegVote  string

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
	SQLAddArtifact    string
	SQLDeleteUser     string
	SQLUsernameExists string
	SQLEmailExists    string

	//SQLInsertMessage     string
	//SQLMessageByMID      string
	//SQLMessageByReceiver string
	//SQLMessageBySender   string
	//SQLMessageHasRead    string
	//SQLDeleteMessage     string

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

type Provider interface {
	PostByPID(pid int) (*post.Post, error)
	PostInfoByPID(pid int) (*post.Info, error)
	PostInfoByPage(size, offset int) ([]*post.Info, error)
	PostInfoIndex(size int) ([]*post.Info, error)
	InsertPost(p *post.Post) (int, error)
	InsertPostInfo(info *post.Info) error
	UpdatePost(p *post.Post) error
	UpdatePostTitle(pid int, title string) error
	UpdatePostLimit(pid, limit int) error
	PostNewReply(pid int) error
	PostNewView(pid int) error
	PostNewMod(pid int) error
	PostNewActivity(pid int) error
	PostNewPosVote(pid, uid int) error
	PostNewNegVote(pid, uid int) error

	SessionByToken(token string) (*session.Session, error)
	SessionsByUID(uid int) ([]*session.Session, error)
	SessionBasicByToken(token string) (*session.Basic, error)
	InsertSession(sess *session.Session) error
	DeleteUserSessions(uid int) error

	UserByUID(uid int) (*user.User, error)
	UserByEmail(email string) (*user.User, error)
	UserByUsername(username string) (*user.User, error)
	UserAuthByUID(uid int) (*user.Auth, error)
	InsertUser(u *user.User) (int, error)
	InsertUserAuth(auth *user.Auth) error
	UpdateUser(u *user.User) error
	UpdateUserAuth(auth *user.Auth) error
	AddArtifact(uid, add int) error
	DeleteUser(uid int) error
	UsernameExists(username string) (bool, error)
	EmailExists(email string) (bool, error)

	SignUpByToken(token string) (*sign_up.SignUp, error)
	SignUpByEmail(email string) (*sign_up.SignUp, error)
	InsertSignUp(su *sign_up.SignUp) error
	DeleteSignUpByEmail(email string) error
	SignUpExists(email string) (bool, error)
}

func New(db *sql.DB, sentence Sentence) *Data {
	return &Data{
		db:       db,
		sentence: sentence,
	}
}

func SQLSentence() Sentence {
	return Sentence{
		SQLPostByPID: `SELECT pid, content, markdown FROM ushio.post WHERE pid = $1;`,
		SQLPostInfoByPID: `SELECT pid, title, creator, created_at, last_mod, replies, 
views, activity, hidden, vote_pos, vote_neg, "limit" FROM ushio.post_info WHERE pid = $1;`,
		SQLPostInfoByPage: `SELECT pid, title, creator, created_at, last_mod, replies, 
views, activity, hidden, vote_pos, vote_neg, "limit" FROM ushio.post_info 
ORDER BY pid DESC LIMIT $1 OFFSET $2;`,
		SQLPostInfoIndex: `SELECT pid, title, creator, created_at, last_mod, replies, 
views, activity, hidden, vote_pos, vote_neg, "limit" FROM ushio.post_info 
WHERE hidden = false ORDER BY last_mod DESC LIMIT $1 OFFSET 0;`,
		SQLInsertPost: `INSERT INTO ushio.post(content, markdown) VALUES ($2, $3);`,
		SQLInsertPostInfo: `INSERT INTO ushio.post_info(title, creator, created_at, 
last_mod, replies, views, activity, hidden, vote_pos, vote_neg, "limit")
VALUES ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`,
		SQLUpdatePost:      `UPDATE ushio.post SET content = $2, markdown = $3 WHERE pid = $1;`,
		SQLUpdatePostTitle: `UPDATE ushio.post_info SET title = $2 WHERE pid = $1;`,
		SQLUpdatePostLimit: `UPDATE ushio.post_info SET "limit" = $2 WHERE pid = $1;`,
		SQLPostNewReply:    `UPDATE ushio.post_info SET replies = replies + 1 WHERE pid = $1;`,
		SQLPostNewView:     `UPDATE ushio.post_info SET views = views + 1 WHERE pid = $1;`,
		SQLPostNewMod:      `UPDATE ushio.post_info SET last_mod = $2, activity = $2 WHERE pid = $1;`,
		SQLPostNewActivity: `UPDATE ushio.post_info SET activity = $2 WHERE pid = $1;`,
		SQLPostNewPosVote:  `UPDATE ushio.post_info SET vote_pos = array_append(vote_pos, $2) WHERE pid = $1;`,
		SQLPostNewNegVote:  `UPDATE ushio.post_info SET vote_neg = array_append(vote_neg, $2) WHERE pid = $1;`,

		SQLSessionByToken: `SELECT token, uid, ua, ip, created_at, expire_at 
FROM ushio.session WHERE token = $1;`,
		SQLSessionsByUID: `SELECT token, uid, ua, ip, created_at, expire_at 
FROM ushio.session WHERE uid = $1;`,
		SQLSessionBasicByToken: `SELECT token, uid, expire_at FROM ushio.session WHERE token = $1;`,
		SQLInsertSession: `INSERT INTO ushio.session(token, uid, ua, ip, created_at, expire_at) 
VALUES ($1, $2, $3, $4, $5, $6);`,
		SQLDeleteUserSessions: `DELETE FROM ushio.session WHERE uid = $1;`,

		SQLUserByUID: `SELECT uid, name, username, email, avatar, bio, created_at, 
is_admin, banned, artifact FROM ushio."user" WHERE uid = $1;`,
		SQLUserByEmail: `SELECT uid, name, username, email, avatar, bio, created_at, 
is_admin, banned, artifact FROM ushio."user" WHERE email = $1;`,
		SQLUserByUsername: `SELECT uid, name, username, email, avatar, bio, created_at, 
is_admin, banned, artifact FROM ushio."user" WHERE username = $1;`,
		SQLUserAuthByUID:  `SELECT uid, password, locked, security_email FROM ushio.user_auth WHERE uid = $1;`,
		SQLInsertUser:     `INSERT INTO ushio.user(uid, name, username, email, avatar, bio, created_at, is_admin, banned, artifact) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`,
		SQLInsertUserAuth: `INSERT INTO ushio.user_auth(uid, password, locked, security_email) VALUES ($1, $2, $3, $4);`,
		SQLUpdateUser:     `UPDATE ushio."user" SET name = $2, username = $3, email = $4, avatar = $5, bio = $6, created_at = $7, is_admin = $8, banned = $9, artifact = $10 WHERE uid = $1;`,
		SQLUpdateUserAuth: `UPDATE ushio.user_auth SET password = $2, locked = $3, security_email = $4 WHERE uid = $1;`,
		SQLAddArtifact:    `UPDATE ushio."user" SET artifact = artifact + $2 WHERE uid = $1;`,
		SQLDeleteUser: `DELETE FROM ushio.user_auth WHERE uid = $1;
DELETE FROM ushio.session WHERE uid = $1;
DELETE FROM ushio."user" WHERE uid = $1;
UPDATE ushio.post_info SET creator = 0 WHERE creator = $1;
UPDATE ushio.comment SET creator = 0 WHERE creator = $1;`,
		SQLUsernameExists: `SELECT EXISTS(SELECT uid FROM ushio."user" WHERE username = $1);`,
		SQLEmailExists:    `SELECT EXISTS(SELECT uid FROM ushio."user" WHERE email = $1);`,

		//SQLInsertMessage:     ``,
		//SQLMessageByMID:      ``,
		//SQLMessageByReceiver: ``,
		//SQLMessageBySender:   ``,
		//SQLMessageHasRead:    ``,
		//SQLDeleteMessage:     ``,

		SQLSignUpByToken:       `SELECT token, email, created_at, expire_at FROM ushio.sign_up WHERE token = $1;`,
		SQLSignUpByEmail:       `SELECT token, email, created_at, expire_at FROM ushio.sign_up WHERE email = $1;`,
		SQLInsertSignUp:        `INSERT INTO ushio.sign_up(token, email, created_at, expire_at) VALUES ($1, $2, $3, $4);`,
		SQLDeleteSignUpByEmail: `DELETE FROM ushio.sign_up WHERE email = $1;`,
		SQLSignUpExists:        `SELECT EXISTS(SELECT token FROM ushio.sign_up WHERE email = $1);`,
	}
}
