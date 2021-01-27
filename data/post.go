package data

import (
	"strconv"

	"github.com/go-ushio/ushio/core/post"
)

func (data *Data) PostInfoByPage(offset, size int) ([]*post.Info, error) {
	row, err := data.db.Query(data.sentence.SQLPostInfoByPage, offset, size)
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

// PostByPID returns a *post.Post by querying the database
// WARNING: the value returned does NOT contain raw markdown content,
//          if necessary, use MarkdownByPID to query.
func (data *Data) PostByPID(pid int) (*post.Post, error) {
	row := data.db.QueryRow(data.sentence.SQLPostByPID, strconv.Itoa(pid))
	p := &post.Post{}
	err := row.Scan(&p.PID, &p.Title, &p.Creator,
		&p.Content, &p.CreatedAt, &p.LastMod, &p.Hidden, &p.Anonymous)
	if err != nil {
		return &post.Post{}, err
	}
	return p, nil
}

func (data *Data) MarkdownByPID(pid int) (string, error) {
	row := data.db.QueryRow(data.sentence.SQLMarkdownByPID, strconv.Itoa(pid))
	md := ""
	err := row.Scan(&md)
	if err != nil {
		return "", err
	}
	return md, nil
}

func (data *Data) InsertPost(p *post.Post) error {
	_, err := data.db.Exec(data.sentence.SQLInsertPost, p.Title, p.Creator, p.Content, p.Hidden, p.Anonymous, p.Markdown)
	if err != nil {
		return err
	}
	return nil
}
