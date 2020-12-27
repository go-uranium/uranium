package data

import (
	"strconv"

	"github.com/go-ushio/ushio/post"
)

var (
	SQLPostByPID     = `SELECT pid, title, creator, content, created_at, last_mod, hidden, anonymous FROM ushio.post WHERE pid = ?`
	SQLMarkdownByPID = `SELECT md_raw FROM ushio.post WHERE pid = ?`

	SQLInsertPost = `INSERT INTO ushio.post(title, creator, content, hidden, anonymous, md_raw) VALUES (?,?,?,?,?,?);`
)

// PostByPID returns a *post.Post by querying the database
// WARNING: the value returned does NOT contain raw markdown content,
//          if necessary, use MarkdownByPID to query.
func PostByPID(pid int) (*post.Post, error) {
	row := db.QueryRow(SQLPostByPID, strconv.Itoa(pid))
	p := &post.Post{}
	err := row.Scan(&p.PID, &p.Title, &p.Creator,
		&p.Content, &p.CreatedAt, &p.LastMod, &p.Hidden, &p.Anonymous)
	if err != nil {
		return &post.Post{}, err
	}
	return p, nil
}

func MarkdownByPID(pid int) (string, error) {
	row := db.QueryRow(SQLMarkdownByPID, strconv.Itoa(pid))
	md := ""
	err := row.Scan(&md)
	if err != nil {
		return "", err
	}
	return md, nil
}

func InsertPost(p *post.Post) error {
	_, err := db.Exec(SQLInsertPost, p.Title, p.Creator, p.Content, p.Hidden, p.Anonymous, p.Markdown)
	if err != nil {
		return err
	}
	return nil
}
