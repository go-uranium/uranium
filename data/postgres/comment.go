package postgres

import (
	"time"

	"github.com/lib/pq"

	"github.com/go-ushio/ushio/core/comment"
)

var (
	SQLCommentsByPost    = `SELECT cid, pid, content, markdown, "user".uid,"user".name,"user".username,"user".avatar, comment.created_at, last_mod, vote_pos, vote_neg FROM ushio.comment INNER JOIN "user" ON uid = creator WHERE pid = $1;`
	SQLCommentByCid      = `SELECT cid, pid, content, markdown, "user".uid,"user".name,"user".username,"user".avatar, comment.created_at, last_mod, vote_pos, vote_neg FROM ushio.comment INNER JOIN "user" ON uid = creator WHERE cid = $1;`
	SQLInsertComment     = `INSERT INTO ushio.comment(pid, content, markdown, creator, created_at, last_mod, vote_pos, vote_neg) VALUES (pid = $1, content = $2, markdown = $3, creator= $4, created_at = $5, last_mod = $6, vote_pos = $7, vote_neg = $8) RETURNING cid;`
	SQLUpdateComment     = `UPDATE ushio.comment SET pid = $2, content = $3, markdown = $4, creator= $5, created_at = $6, last_mod = $7, vote_pos = $8, vote_neg = $9 WHERE cid = $1;`
	SQLCommentNewMod     = `UPDATE ushio.comment SET last_mod = $2 WHERE cid = $1;`
	SQLCommentNewPosVote = `UPDATE ushio.comment SET vote_pos = array_append(vote_pos, $2) WHERE cid = $1;`
	SQLCommentNewNegVote = `UPDATE ushio.comment SET vote_neg = array_append(vote_neg, $2) WHERE cid = $1;`
)

func (pg *Postgres) CommentsByPost(pid int64) ([]*comment.Comment, error) {
	row, err := pg.db.Query(SQLCommentsByPost, pid)
	if err != nil {
		return nil, err
	}
	var cmts []*comment.Comment
	for row.Next() {
		cmt, err := comment.ScanComment(row)
		if err != nil {
			return nil, err
		}
		cmts = append(cmts, cmt)
	}
	return cmts, nil
}

func (pg *Postgres) CommentByCid(cid int64) (*comment.Comment, error) {
	row := pg.db.QueryRow(SQLCommentByCid, cid)
	cmt, err := comment.ScanComment(row)
	return cmt, err
}

func (pg *Postgres) InsertComment(cmt *comment.Comment) (int64, error) {
	var cid int64
	err := pg.db.QueryRow(SQLInsertComment, cmt.PID,
		cmt.Content, cmt.Markdown, cmt.Creator, cmt.CreatedAt,
		cmt.LastMod, pq.Array(cmt.VotePos), pq.Array(cmt.VoteNeg)).
		Scan(&cid)
	return cid, err
}

func (pg *Postgres) UpdateComment(cmt *comment.Comment) error {
	_, err := pg.db.Exec(SQLUpdateComment, cmt.CID, cmt.PID,
		cmt.Content, cmt.Markdown, cmt.Creator, cmt.CreatedAt,
		cmt.LastMod, pq.Array(cmt.VotePos), pq.Array(cmt.VoteNeg))
	return err
}

func (pg *Postgres) CommentNewMod(cid int64) error {
	_, err := pg.db.Exec(SQLCommentNewMod, cid, time.Now())
	return err
}

func (pg *Postgres) CommentNewPosVote(cid, uid int64) error {
	_, err := pg.db.Exec(SQLCommentNewPosVote, cid, uid)
	return err
}

func (pg *Postgres) CommentNewNegVote(cid, uid int64) error {
	_, err := pg.db.Exec(SQLCommentNewNegVote, cid, uid)
	return err
}
