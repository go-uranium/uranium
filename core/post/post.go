package post

import (
	"database/sql"
	"html/template"
	"time"

	"github.com/lib/pq"

	"github.com/go-ushio/ushio/common/put"
	"github.com/go-ushio/ushio/common/scan"
)

type Info struct {
	PID       int       `json:"pid"`
	Title     string    `json:"title"`
	Creator   int       `json:"creator"`
	CreatedAt time.Time `json:"created_at"`
	LastMod   time.Time `json:"last_mod"`
	Replies   int       `json:"replies"`
	Views     int       `json:"views"`
	Activity  time.Time `json:"activity"`
	Hidden    bool      `json:"hidden"`
	Anonymous bool      `json:"anonymous"`
	// uid list
	VotePos []int `json:"vote_pos"`
	VoteNeg []int `json:"vote_neg"`
}

type Post struct {
	PID      int           `json:"pid"`
	Info     *Info         `json:"info"`
	Content  template.HTML `json:"content"`
	Markdown string        `json:"markdown"`
}

func ScanPost(scanner scan.Scanner) (*Post, error) {
	post := &Post{}
	err := scanner.Scan(&post.PID, &post.Content)
	if err != nil {
		return &Post{}, err
	}
	return post, nil
}

func ScanInfo(scanner scan.Scanner) (*Info, error) {
	info := &Info{}
	err := scanner.Scan(&info.PID, &info.Title, &info.Creator,
		&info.CreatedAt, &info.LastMod,
		&info.Replies, &info.Views, &info.Activity,
		&info.Hidden, &info.Anonymous,
		pq.Array(&info.VotePos), pq.Array(&info.VoteNeg))
	if err != nil {
		return &Info{}, err
	}
	return info, nil
}

func (post *Post) Put(putter put.Putter) (sql.Result, error) {
	return putter.Put(post.PID, post.Content, post.Markdown)
}

func (info *Info) Put(putter put.Putter) (sql.Result, error) {
	return putter.Put(info.PID, info.Title, info.Creator,
		info.CreatedAt, info.LastMod, info.Replies, info.Views,
		info.Activity, info.Hidden, info.Anonymous,
		pq.Array(info.VotePos), pq.Array(info.VoteNeg))
}

func (post *Post) Copy() *Post {
	post2 := *post
	post2.Info = post.Info.Copy()
	return &post2
}

func (info *Info) Copy() *Info {
	info2 := *info
	return &info2
}
