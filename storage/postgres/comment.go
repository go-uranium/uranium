package postgres

import (
	"time"

	"github.com/lib/pq"

	"github.com/go-uranium/uranium/model/comment"
)

var (
	SQLCommentsByPost = `SELECT comments.cid, comments.pid, comments.content, comments.markdown, users.uid,users.name,users.username,users.avatar, comments.created_at, last_mod, vote_pos, vote_neg FROM comments INNER JOIN users ON uid = creator WHERE pid = $1;`
	SQLCommentByCID   = `SELECT comments.cid, comments.pid, comments.content, comments.markdown, users.uid,users.name,users.username,users.avatar, comments.created_at, last_mod, vote_pos, vote_neg FROM comments INNER JOIN users ON uid = creator WHERE cid = $1;`
	SQLCommentsByUID  = `SELECT comments.cid, comments.pid, comments.content, comments.markdown, users.uid,users.name,users.username,users.avatar, comments.created_at, last_mod, vote_pos, vote_neg FROM comments INNER JOIN users ON uid = creator WHERE creator = $1;`

	SQLInsertComment     = `INSERT INTO comments(pid, content, markdown, creator, created_at, last_mod, vote_pos, vote_neg) VALUES (pid = $1, content = $2, markdown = $3, creator= $4, created_at = $5, last_mod = $6, vote_pos = $7, vote_neg = $8) RETURNING cid;`
	SQLUpdateComment     = `UPDATE comments SET pid = $2, content = $3, markdown = $4, creator= $5, created_at = $6, last_mod = $7, vote_pos = $8, vote_neg = $9 WHERE cid = $1;`
	SQLCommentNewMod     = `UPDATE comments SET last_mod = $2 WHERE cid = $1;`
	SQLCommentNewPosVote = `UPDATE comments SET vote_pos = array_append(vote_pos, $2) WHERE cid = $1;`
	SQLCommentNewNegVote = `UPDATE comments SET vote_neg = array_append(vote_neg, $2) WHERE cid = $1;`
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

func (pg *Postgres) CommentByCID(cid int64) (*comment.Comment, error) {
	row := pg.db.QueryRow(SQLCommentByCID, cid)
	cmt, err := comment.ScanComment(row)
	return cmt, err
}

func (pg *Postgres) CommentByUID(uid int64) ([]*comment.Comment, error) {
	row, err := pg.db.Query(SQLCommentsByUID, uid)
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
