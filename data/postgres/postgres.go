package postgres

import "database/sql"

type Postgres struct {
	db *sql.DB
}

func New(db *sql.DB) *Postgres {
	return &Postgres{
		db: db,
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
		SQLInsertPost: `INSERT INTO ushio.post(content, markdown) VALUES ($1, $2) RETURNING pid;`,
		SQLInsertPostInfo: `INSERT INTO ushio.post_info(pid, title, creator, created_at, 
last_mod, replies, views, activity, hidden, vote_pos, vote_neg, "limit")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`,
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

		postgres.SQLUserByUID:,
		postgres.SQLUserByEmail: `SELECT uid, name, username, email, avatar, bio, created_at, 
is_admin, banned, artifact FROM ushio."user" WHERE email = $1;`,
		postgres.SQLUserByUsername: `SELECT uid, name, username, email, avatar, bio, created_at, 
is_admin, banned, artifact FROM ushio."user" WHERE username = $1;`,
		postgres.SQLUserAuthByUID:  `SELECT uid, password, locked, security_email FROM ushio.user_auth WHERE uid = $1;`,
		postgres.SQLInsertUser:     `INSERT INTO ushio.user(name, username, email, avatar, bio, created_at, is_admin, banned, artifact) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING uid;`,
		postgres.SQLInsertUserAuth: `INSERT INTO ushio.user_auth(uid, password, locked, security_email) VALUES ($1, $2, $3, $4);`,
		postgres.SQLUpdateUser:     `UPDATE ushio."user" SET name = $2, username = $3, email = $4, avatar = $5, bio = $6, created_at = $7, is_admin = $8, banned = $9, artifact = $10 WHERE uid = $1;`,
		postgres.SQLUpdateUserAuth: `UPDATE ushio.user_auth SET password = $2, locked = $3, security_email = $4 WHERE uid = $1;`,
		postgres.SQLAddArtifact:    `UPDATE ushio."user" SET artifact = artifact + $2 WHERE uid = $1;`,
		postgres.SQLDeleteUser: `DELETE FROM ushio.user_auth WHERE uid = $1;
DELETE FROM ushio.session WHERE uid = $1;
DELETE FROM ushio."user" WHERE uid = $1;
UPDATE ushio.post_info SET creator = 0 WHERE creator = $1;
UPDATE ushio.comment SET creator = 0 WHERE creator = $1;`,
		postgres.SQLUsernameExists: `SELECT EXISTS(SELECT uid FROM ushio."user" WHERE username = $1);`,
		postgres.SQLEmailExists:    `SELECT EXISTS(SELECT uid FROM ushio."user" WHERE email = $1);`,

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

		SQLGetCategories: `SELECT tid, tname, name, color, admin FROM ushio.category;`,
	}
}
