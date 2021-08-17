package post

import (
	"html/template"
	"time"

	"github.com/go-uranium/uranium/model/category"
	"github.com/go-uranium/uranium/model/markdown"
	"github.com/go-uranium/uranium/model/user"
)

type Info struct {
	PID      int64             `json:"pid"`
	Title    string            `json:"title"`
	Category category.Category `json:"category"`
	Creator  user.Basic        `json:"creator"`
	Views    int64             `json:"views"`
	Replies  int64             `json:"replies"`
	Activity time.Time         `json:"activity"`
	// uid list
	VotePos []int64 `json:"vote_pos"`
	// uid list
	VoteNeg []int64 `json:"vote_neg"`
	// if user.permission < postInfo.permission &&
	//    user.permission >= 0  {
	//      // no access to thread/post
	// }
	//
	Permission int32     `json:"limit"`
	Hidden     bool      `json:"hidden"`
	CreatedAt  time.Time `json:"created_at"`
	LastMod    time.Time `json:"last_mod"`
}

type Post struct {
	PID      int64              `json:"pid"`
	Info     *Info              `json:"info"`
	Content  template.HTML      `json:"content"`
	Markdown *markdown.Markdown `json:"markdown"`
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
