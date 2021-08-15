package comment

import (
	"html/template"
	"time"

	"github.com/go-ushio/ushio/model/markdown"
	"github.com/go-ushio/ushio/model/user"
)

type Comment struct {
	CID       int64              `json:"cid"`
	PID       int64              `json:"pid"`
	Content   template.HTML      `json:"content"`
	Markdown  *markdown.Markdown `json:"markdown"`
	Creator   user.Basic         `json:"creator"`
	CreatedAt time.Time          `json:"created_at"`
	LastMod   time.Time          `json:"last_mod"`
	// uid list
	VotePos []int64 `json:"vote_pos"`
	VoteNeg []int64 `json:"vote_neg"`
}

