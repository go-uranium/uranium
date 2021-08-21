package post

import (
	"html/template"
	"time"

	"github.com/go-uranium/uranium/model/category"
	"github.com/go-uranium/uranium/model/user"
)

// Info is stored in storage
type Info struct {
	PID       int32             `json:"pid"`
	Title     string            `json:"title"`
	Category  category.Category `json:"category"`
	Creator   user.Basic        `json:"creator"`
	Statistic *Statistic        `json:"statistic"`
	Limit     int32             `json:"limit"`
	Hidden    bool              `json:"hidden"`
	Modified  time.Time         `json:"modified"`
	Created   time.Time         `json:"created"`
}

// Statistic is stored in Cache
type Statistic struct {
	PID      int32     `json:"pid"`
	Views    int32     `json:"views"`
	Replies  int32     `json:"replies"`
	Activity time.Time `json:"activity"`
	VotePos  []int32   `json:"vote_pos"`
	VoteNeg  []int32   `json:"vote_neg"`
}

type Post struct {
	PID     int64          `json:"pid"`
	Info    *Info          `json:"info"`
	Mention user.BasicList `json:"mention"`
	Content template.HTML  `json:"content"`
}

type RawContent struct {
}

func (info *Info) VotePosContains(uid int64) bool {
	for _, u := range info.VotePos {
		if u == uid {
			return true
		}
	}
	return false
}

func (info *Info) VoteNegContains(uid int64) bool {
	for _, u := range info.VoteNeg {
		if u == uid {
			return true
		}
	}
	return false
}

func (info *Info) VotePosCount() int {
	return len(info.VotePos)
}

func (info *Info) VoteNegCount() int {
	return len(info.VoteNeg)
}

func (info *Info) VoteCount() int {
	return len(info.VotePos) - len(info.VoteNeg)
}

//func (post *Post) Copy() *Post {
//	post2 := *post
//	post2.Info = post.Info.Copy()
//	post2.Markdown = post.Markdown.Copy()
//	return &post2
//}
//
//func (info *Info) Copy() *Info {
//	info2 := *info
//	return &info2
//}
