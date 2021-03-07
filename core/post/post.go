package post

import (
	"html/template"
	"time"

	"github.com/lib/pq"

	"github.com/go-ushio/ushio/common/scan"
)

type Info struct {
	PID       int64     `json:"pid"`
	Title     string    `json:"title"`
	Creator   int64     `json:"creator"`
	CreatedAt time.Time `json:"created_at"`
	LastMod   time.Time `json:"last_mod"`
	Replies   int64     `json:"replies"`
	Views     int64     `json:"views"`
	Activity  time.Time `json:"activity"`
	Hidden    bool      `json:"hidden"`
	// uid list
	VotePos  []int64 `json:"vote_pos"`
	VoteNeg  []int64 `json:"vote_neg"`
	Limit    int64   `json:"limit"`
	Category int64   `json:"category"`
}

type Post struct {
	PID      int64         `json:"pid"`
	Info     *Info         `json:"info"`
	Content  template.HTML `json:"content"`
	Markdown string        `json:"markdown"`
}

func ScanPost(scanner scan.Scanner) (*Post, error) {
	post := &Post{}
	err := scanner.Scan(&post.PID, &post.Content, &post.Markdown)
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
		&info.Hidden, pq.Array(&info.VotePos),
		pq.Array(&info.VoteNeg), &info.Limit, &info.Category)
	if err != nil {
		return &Info{}, err
	}
	return info, nil
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
