package postgres

import (
	"time"

	"github.com/lib/pq"

	"github.com/go-ushio/ushio/core/post"
)

var (
	SQLPostByPID       = `SELECT pid, content, markdown FROM ushio.post WHERE pid = $1;`
	SQLPostInfoByPID   = `SELECT post_info.pid, post_info.title, "user".uid,"user".name,"user".username,"user".avatar, post_info.created_at, post_info.last_mod, post_info.replies, post_info.views, post_info.activity, post_info.hidden, post_info.vote_pos, post_info.vote_neg, post_info."limit", ushio.category.tid, ushio.category.tname, ushio.category.name, ushio.category.color, ushio.category.admin FROM ushio.post_info INNER JOIN "user" ON uid = creator INNER JOIN "category" ON tid = post_info.category WHERE pid = $1;`
	SQLPostInfoByPage  = `SELECT post_info.pid, post_info.title, "user".uid,"user".name,"user".username,"user".avatar, post_info.created_at, post_info.last_mod, post_info.replies, post_info.views, post_info.activity, post_info.hidden, post_info.vote_pos, post_info.vote_neg, post_info."limit", ushio.category.tid, ushio.category.tname, ushio.category.name, ushio.category.color, ushio.category.admin FROM ushio.post_info INNER JOIN "user" ON uid = creator INNER JOIN "category" ON tid = post_info.category ORDER BY pid DESC LIMIT $1 OFFSET $2;`
	SQLPostInfoIndex   = `SELECT post_info.pid, post_info.title, "user".uid,"user".name,"user".username,"user".avatar, post_info.created_at, post_info.last_mod, post_info.replies, post_info.views, post_info.activity, post_info.hidden, post_info.vote_pos, post_info.vote_neg, post_info."limit", ushio.category.tid, ushio.category.tname, ushio.category.name, ushio.category.color, ushio.category.admin FROM ushio.post_info INNER JOIN "user" ON uid = creator INNER JOIN "category" ON tid = post_info.category WHERE hidden = false ORDER BY last_mod DESC LIMIT $1 OFFSET 0;`
	SQLInsertPost      = `INSERT INTO ushio.post(content, markdown) VALUES ($1, $2) RETURNING pid;`
	SQLInsertPostInfo  = `INSERT INTO ushio.post_info(pid, title, creator, created_at, last_mod, replies, views, activity, hidden, vote_pos, vote_neg, "limit", category) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);`
	SQLUpdatePost      = `UPDATE ushio.post SET content = $2, markdown = $3 WHERE pid = $1;`
	SQLUpdatePostTitle = `UPDATE ushio.post_info SET title = $2 WHERE pid = $1;`
	SQLUpdatePostLimit = `UPDATE ushio.post_info SET "limit" = $2 WHERE pid = $1;`
	SQLPostNewReply    = `UPDATE ushio.post_info SET replies = replies + 1 WHERE pid = $1;`
	SQLPostNewView     = `UPDATE ushio.post_info SET views = views + 1 WHERE pid = $1;`
	SQLPostNewMod      = `UPDATE ushio.post_info SET last_mod = $2, activity = $2 WHERE pid = $1;`
	SQLPostNewActivity = `UPDATE ushio.post_info SET activity = $2 WHERE pid = $1;`
	SQLPostNewPosVote  = `UPDATE ushio.post_info SET vote_pos = array_append(vote_pos, $2) WHERE pid = $1;`
	SQLPostNewNegVote  = `UPDATE ushio.post_info SET vote_neg = array_append(vote_neg, $2) WHERE pid = $1;`
	SQLPostedBy        = `SELECT post_info.pid, post_info.title, "user".uid,"user".name,"user".username,"user".avatar, post_info.created_at, post_info.last_mod, post_info.replies, post_info.views, post_info.activity, post_info.hidden, post_info.vote_pos, post_info.vote_neg, post_info."limit", ushio.category.tid, ushio.category.tname, ushio.category.name, ushio.category.color, ushio.category.admin FROM ushio.post_info INNER JOIN "user" ON uid = creator INNER JOIN "category" ON tid = post_info.category WHERE creator = $1;`
	//SQLPostedByAfter        = `SELECT pid, title, creator, created_at, last_mod, replies, views, activity, hidden, vote_pos, vote_neg, "limit", category FROM ushio.post_info WHERE creator = $1;`
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

func (pg *Postgres) PostInfoByPage(size, offset int64) ([]*post.Info, error) {
	row, err := pg.db.Query(SQLPostInfoByPage, size, offset)
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

func (pg *Postgres) PostInfoIndex(size int64) ([]*post.Info, error) {
	row, err := pg.db.Query(SQLPostInfoIndex, size)
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
	_, err := pg.db.Exec(SQLInsertPostInfo, info.PID, info.Title, info.Creator,
		info.CreatedAt, info.LastMod, info.Replies, info.Views,
		info.Activity, info.Hidden, pq.Array(info.VotePos),
		pq.Array(info.VoteNeg), info.Limit, info.Category)
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

func (pg *Postgres) PostedBy(uid int64) ([]*post.Info, error) {
	rows, err := pg.db.Query(SQLPostedBy, uid)
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
