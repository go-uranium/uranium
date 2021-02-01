package post

import (
	"database/sql"
	"html/template"
	"time"

	"github.com/go-ushio/ushio/common/put"
	"github.com/go-ushio/ushio/common/scan"
)

type Post struct {
	PID      int `json:"pid"`
	Info     *Info
	Content  template.HTML `json:"content"`
	Markdown string        `json:"markdown"`
}

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
		&info.Hidden, &info.Anonymous)
	if err != nil {
		return &Info{}, err
	}
	return info, nil
}

func (post *Post) Put(putter put.Putter) (sql.Result, error) {
	return putter.Put(post.Content, post.Markdown)
}

func (post *Post) PutWithPIDFirst(putter put.Putter) (sql.Result, error) {
	return putter.Put(post.PID, post.Content, post.Markdown)
}

func (post *Post) PutWithPIDLast(putter put.Putter) (sql.Result, error) {
	return putter.Put(post.Content, post.Markdown, post.PID)
}

func (info *Info) Put(putter put.Putter) (sql.Result, error) {
	return putter.Put(info.Title, info.Creator,
		info.CreatedAt, info.LastMod, info.Replies, info.Views,
		info.Activity, info.Hidden, info.Anonymous)
}

func (info *Info) PutWithPIDFirst(putter put.Putter) (sql.Result, error) {
	return putter.Put(info.PID, info.Title, info.Creator,
		info.CreatedAt, info.LastMod, info.Replies, info.Views,
		info.Activity, info.Hidden, info.Anonymous)
}

func (info *Info) PutWithPIDLast(putter put.Putter) (sql.Result, error) {
	return putter.Put(info.Title, info.Creator, info.CreatedAt,
		info.LastMod, info.Replies, info.Views, info.Activity,
		info.Hidden, info.Anonymous, info.PID)
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
