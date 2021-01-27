package post

import (
	"html/template"
	"time"

	"github.com/go-ushio/ushio/common/scan"
)

type Post struct {
	Info     *Info
	Content  template.HTML `json:"content"`
	Markdown string        `json:"markdown"`
}

type Info struct {
	PID       int       `json:"pid"`
	Title     string    `json:"title"`
	Creator   int       `json:"creator"`
	Posters   []int     `json:"posters"`
	CreatedAt time.Time `json:"created_at"`
	LastMod   time.Time `json:"last_mod"`
	Replies   int       `json:"replies"`
	Views     int       `json:"views"`
	Activity  time.Time `json:"activity"`
	Hidden    bool      `json:"hidden"`
	Anonymous bool      `json:"anonymous"`
}

func ScanInfo(scanner scan.Scanner) (*Info, error) {
	info := &Info{}
	err := scanner.Scan(&info.PID, &info.Title, &info.Creator,
		&info.Posters, &info.CreatedAt, &info.LastMod,
		&info.Replies, &info.Views, &info.Activity,
		&info.Hidden, &info.Anonymous)
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
