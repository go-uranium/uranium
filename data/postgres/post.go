package postgres

import (
	"database/sql"
	"time"

	"github.com/lib/pq"

	"github.com/go-ushio/ushio/model/post"
)

var (
	SQLPostByPID     = `SELECT pid, content, markdown FROM posts WHERE pid = $1;`
	SQLPostInfoByPID = `SELECT post_info.pid, post_info.title, users.uid, users.name, users.username, users.avatar, post_info.created_at, post_info.last_mod, post_info.replies, post_info.views, post_info.activity, post_info.hidden, post_info.vote_pos, post_info.vote_neg, post_info."limit", categories.tid, categories.tname, categories.name, categories.color, categories.admin FROM post_info INNER JOIN users ON uid = creator INNER JOIN categories ON tid = post_info.category WHERE pid = $1;`

	SQLPostsInfoByActivityWithHidden = `SELECT post_info.pid, post_info.title, users.uid,users.name,users.username,users.avatar, post_info.created_at, post_info.last_mod, post_info.replies, post_info.views, post_info.activity, post_info.hidden, post_info.vote_pos, post_info.vote_neg, post_info."limit", categories.tid, categories.tname, categories.name, categories.color, categories.admin FROM post_info INNER JOIN users ON uid = creator INNER JOIN categories ON tid = post_info.category ORDER BY activity DESC LIMIT $1 OFFSET $2;`
	SQLPostsInfoByActivity           = `SELECT post_info.pid, post_info.title, users.uid,users.name,users.username,users.avatar, post_info.created_at, post_info.last_mod, post_info.replies, post_info.views, post_info.activity, post_info.hidden, post_info.vote_pos, post_info.vote_neg, post_info."limit", categories.tid, categories.tname, categories.name, categories.color, categories.admin FROM post_info INNER JOIN users ON uid = creator INNER JOIN categories ON tid = post_info.category WHERE hidden = false ORDER BY activity DESC LIMIT $1 OFFSET $2;`

	SQLPostsInfoByCategoryWithHidden = `SELECT post_info.pid, post_info.title, users.uid,users.name,users.username,users.avatar, post_info.created_at, post_info.last_mod, post_info.replies, post_info.views, post_info.activity, post_info.hidden, post_info.vote_pos, post_info.vote_neg, post_info."limit", categories.tid, categories.tname, categories.name, categories.color, categories.admin FROM post_info INNER JOIN users ON uid = creator INNER JOIN categories ON tid = post_info.category WHERE post_info.category = $3 ORDER BY activity DESC LIMIT $1 OFFSET $2;`
	SQLPostsInfoByCategory           = `SELECT post_info.pid, post_info.title, users.uid,users.name,users.username,users.avatar, post_info.created_at, post_info.last_mod, post_info.replies, post_info.views, post_info.activity, post_info.hidden, post_info.vote_pos, post_info.vote_neg, post_info."limit", categories.tid, categories.tname, categories.name, categories.color, categories.admin FROM post_info INNER JOIN users ON uid = creator INNER JOIN categories ON tid = post_info.category WHERE post_info.category = $3 AND hidden = false ORDER BY activity DESC LIMIT $1 OFFSET $2;`

	SQLPostsInfoByPIDWithHidden = `SELECT post_info.pid, post_info.title, users.uid,users.name,users.username,users.avatar, post_info.created_at, post_info.last_mod, post_info.replies, post_info.views, post_info.activity, post_info.hidden, post_info.vote_pos, post_info.vote_neg, post_info."limit", categories.tid, categories.tname, categories.name, categories.color, categories.admin FROM post_info INNER JOIN users ON uid = creator INNER JOIN categories ON tid = post_info.category ORDER BY pid DESC LIMIT $1 OFFSET $2;`
	SQLPostsInfoByPID           = `SELECT post_info.pid, post_info.title, users.uid,users.name,users.username,users.avatar, post_info.created_at, post_info.last_mod, post_info.replies, post_info.views, post_info.activity, post_info.hidden, post_info.vote_pos, post_info.vote_neg, post_info."limit", categories.tid, categories.tname, categories.name, categories.color, categories.admin FROM post_info INNER JOIN users ON uid = creator INNER JOIN categories ON tid = post_info.category ORDER BY pid DESC LIMIT $1 OFFSET $2;`

	SQLPostsInfoByUID = `SELECT post_info.pid, post_info.title, users.uid,users.name,users.username,users.avatar, post_info.created_at, post_info.last_mod, post_info.replies, post_info.views, post_info.activity, post_info.hidden, post_info.vote_pos, post_info.vote_neg, post_info."limit", categories.tid, categories.tname, categories.name, categories.color, categories.admin FROM post_info INNER JOIN users ON uid = creator INNER JOIN categories ON tid = post_info.category WHERE post_info.creator = $3 ORDER BY activity DESC LIMIT $1 OFFSET $2;`

	SQLPostsInfoByCommentCreator = `SELECT post_info.pid, post_info.title, users.uid,users.name,users.username,users.avatar, post_info.created_at, post_info.last_mod, post_info.replies, post_info.views, post_info.activity, post_info.hidden, post_info.vote_pos, post_info.vote_neg, post_info."limit", categories.tid, categories.tname, categories.name, categories.color, categories.admin FROM post_info INNER JOIN users ON uid = creator INNER JOIN categories ON tid = post_info.category WHERE post_info.pid = ANY (SELECT comments.pid FROM comments WHERE comments.creator = $3) ORDER BY activity DESC LIMIT $1 OFFSET $2;`

	SQLInsertPost     = `INSERT INTO posts(content, markdown) VALUES ($1, $2) RETURNING pid;`
	SQLInsertPostInfo = `INSERT INTO post_info(pid, title, creator, created_at, last_mod, replies, views, activity, hidden, vote_pos, vote_neg, "limit", category) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);`

	SQLUpdatePost      = `UPDATE posts SET content = $2, markdown = $3 WHERE pid = $1;`
	SQLUpdatePostTitle = `UPDATE post_info SET title = $2 WHERE pid = $1;`
	SQLUpdatePostLimit = `UPDATE post_info SET "limit" = $2 WHERE pid = $1;`

	SQLPostNewReply    = `UPDATE post_info SET replies = replies + 1 WHERE pid = $1;`
	SQLPostNewView     = `UPDATE post_info SET views = views + 1 WHERE pid = $1;`
	SQLPostNewMod      = `UPDATE post_info SET last_mod = $2, activity = $2 WHERE pid = $1;`
	SQLPostNewActivity = `UPDATE post_info SET activity = $2 WHERE pid = $1;`
	SQLPostNewPosVote  = `UPDATE post_info SET vote_pos = array_append(vote_pos, $2) WHERE pid = $1;`
	SQLPostNewNegVote  = `UPDATE post_info SET vote_neg = array_append(vote_neg, $2) WHERE pid = $1;`

	SQLPostRemovePosVote = `UPDATE post_info SET vote_pos = array_remove(vote_pos, $2) WHERE pid = $1;`
	SQLPostRemoveNegVote = `UPDATE post_info SET vote_neg = array_remove(vote_neg, $2) WHERE pid = $1;`

	//SQLPostedBy = `SELECT post_info.pid, post_info.title, users.uid,users.name,users.username,users.avatar, post_info.created_at, post_info.last_mod, post_info.replies, post_info.views, post_info.activity, post_info.hidden, post_info.vote_pos, post_info.vote_neg, post_info."limit", categories.tid, categories.tname, categories.name, categories.color, categories.admin FROM post_info INNER JOIN users ON uid = creator INNER JOIN categories ON tid = post_info.category WHERE creator = $1 ORDER BY activity DESC;`

	//SQLPostedByAfter        = `SELECT pid, title, creator, created_at, last_mod, replies, views, activity, hidden, vote_pos, vote_neg, "limit", category FROM post_info WHERE creator = $1;`
)

func (pg *Postgres) PostByPID(pid int64) (*post.Post, error) {
	row := pg.db.QueryRow(SQLPostByPID, pid)
	p, err := post.ScanPost(row)
	if err != nil {
		//if err == sql.ErrNoRows {
		//	return nil, nil
		//}
		return &post.Post{}, err
	}
	info, err := pg.PostInfoByPID(pid)
	if err != nil {
		return &post.Post{}, err
	}
	p.Info = info
	return p, nil
}

func (pg *Postgres) PostInfoByPID(pid int64) (*post.Info, error) {
	row := pg.db.QueryRow(SQLPostInfoByPID, pid)
	info, err := post.ScanInfo(row)
	if err != nil {
		//if err == sql.ErrNoRows {
		//	return nil, nil
		//}
		return &post.Info{}, err
	}
	return info, nil
}

func (pg *Postgres) PostsInfoByActivity(hidden bool, size, offset int64) ([]*post.Info, error) {
	var row *sql.Rows
	var err error
	if hidden {
		row, err = pg.db.Query(SQLPostsInfoByActivityWithHidden, size, offset)
	} else {
		row, err = pg.db.Query(SQLPostsInfoByActivity, size, offset)
	}
	if err != nil {
		return nil, err
	}

	var postInfoList []*post.Info
	for row.Next() {
		postInfo, err := post.ScanInfo(row)
		if err != nil {
			return nil, err
		}
		postInfoList = append(postInfoList, postInfo)
	}
	return postInfoList, nil
}

func (pg *Postgres) PostsInfoByCategory(hidden bool, size, offset, category int64) ([]*post.Info, error) {
	var row *sql.Rows
	var err error

	if hidden {
		row, err = pg.db.Query(SQLPostsInfoByCategoryWithHidden, size, offset, category)
	} else {
		row, err = pg.db.Query(SQLPostsInfoByCategory, size, offset, category)
	}
	if err != nil {
		return nil, err
	}

	var postInfoList []*post.Info
	for row.Next() {
		postInfo, err := post.ScanInfo(row)
		if err != nil {
			return nil, err
		}
		postInfoList = append(postInfoList, postInfo)
	}
	return postInfoList, nil
}

func (pg *Postgres) PostsInfoByPID(hidden bool, size, offset int64) ([]*post.Info, error) {
	var row *sql.Rows
	var err error

	if hidden {
		row, err = pg.db.Query(SQLPostsInfoByPIDWithHidden, size, offset)
	} else {
		row, err = pg.db.Query(SQLPostsInfoByPID, size, offset)
	}
	if err != nil {
		return nil, err
	}

	var postInfoList []*post.Info
	for row.Next() {
		postInfo, err := post.ScanInfo(row)
		if err != nil {
			return nil, err
		}
		postInfoList = append(postInfoList, postInfo)
	}
	return postInfoList, nil
}

func (pg *Postgres) PostsInfoByCommentCreator(size, offset, uid int64) ([]*post.Info, error) {
	row, err := pg.db.Query(SQLPostsInfoByCommentCreator, size, offset, uid)
	if err != nil {
		return nil, err
	}

	var postInfoList []*post.Info
	for row.Next() {
		postInfo, err := post.ScanInfo(row)
		if err != nil {
			return nil, err
		}
		postInfoList = append(postInfoList, postInfo)
	}
	return postInfoList, nil
}

func (pg *Postgres) PostsInfoByUID(size, offset, uid int64) ([]*post.Info, error) {
	rows, err := pg.db.Query(SQLPostsInfoByUID,size, offset, uid)
	if err != nil {
		return nil, err
	}
	var posts []*post.Info
	for rows.Next() {
		p, err := post.ScanInfo(rows)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}


func (pg *Postgres) InsertPost(p *post.Post) (int64, error) {
	var pid int64
	err := pg.db.QueryRow(SQLInsertPost, p.Content, p.Markdown).
		Scan(&pid)
	if err != nil {
		return 0, err
	}
	p.Info.PID = pid
	return pid, pg.InsertPostInfo(p.Info)
}

func (pg *Postgres) InsertPostInfo(info *post.Info) error {
	_, err := pg.db.Exec(SQLInsertPostInfo, info.PID, info.Title, info.Creator.UID,
		info.CreatedAt, info.LastMod, info.Replies, info.Views,
		info.Activity, info.Hidden, pq.Array(info.VotePos),
		pq.Array(info.VoteNeg), info.Limit, info.Category.TID)
	return err
}

func (pg *Postgres) UpdatePost(p *post.Post) error {
	_, err := pg.db.Exec(SQLUpdatePost, p.PID, p.Content, p.Markdown)
	if err != nil {
		return err
	}
	err = pg.InsertPostInfo(p.Info)
	return err
}

func (pg *Postgres) UpdatePostTitle(pid int64, title string) error {
	_, err := pg.db.Exec(SQLUpdatePostTitle, pid, title)
	return err
}

func (pg *Postgres) UpdatePostLimit(pid, limit int64) error {
	_, err := pg.db.Exec(SQLUpdatePostLimit, pid, limit)
	return err
}

func (pg *Postgres) PostNewReply(pid int64) error {
	_, err := pg.db.Exec(SQLPostNewReply, pid)
	return err
}

func (pg *Postgres) PostNewView(pid int64) error {
	_, err := pg.db.Exec(SQLPostNewView, pid)
	return err
}

func (pg *Postgres) PostNewMod(pid int64) error {
	_, err := pg.db.Exec(SQLPostNewMod, pid, time.Now())
	return err
}

func (pg *Postgres) PostNewActivity(pid int64) error {
	_, err := pg.db.Exec(SQLPostNewActivity, pid, time.Now())
	return err
}

func (pg *Postgres) PostNewPosVote(pid, uid int64) error {
	_, err := pg.db.Exec(SQLPostNewPosVote, pid, uid)
	return err
}

func (pg *Postgres) PostNewNegVote(pid, uid int64) error {
	_, err := pg.db.Exec(SQLPostNewNegVote, pid, uid)
	return err
}

func (pg *Postgres) PostRemovePosVote(pid, uid int64) error {
	_, err := pg.db.Exec(SQLPostRemovePosVote, pid, uid)
	return err
}

func (pg *Postgres) PostRemoveNegVote(pid, uid int64) error {
	_, err := pg.db.Exec(SQLPostRemoveNegVote, pid, uid)
	return err
}

