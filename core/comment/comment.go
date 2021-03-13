package comment

import (
	"html/template"
	"time"

	"github.com/lib/pq"

	"github.com/go-ushio/ushio/common/scan"
	"github.com/go-ushio/ushio/core/user"
)

type Comment struct {
	// CID is an UUID
	CID       string        `json:"cid"`
	PID       int64         `json:"pid"`
	Content   template.HTML `json:"content"`
	Markdown  string        `json:"markdown"`
	Creator   user.Simple   `json:"creator"`
	CreatedAt time.Time     `json:"created_at"`
	LastMod   time.Time     `json:"last_mod"`
	// uid list
	VotePos []int64 `json:"vote_pos"`
	VoteNeg []int64 `json:"vote_neg"`
}

func ScanComment(scanner scan.Scanner) (*Comment, error) {
	cmt := &Comment{}
	err := scanner.Scan(&cmt.CID, &cmt.PID, &cmt.Content, &cmt.Markdown,
		// Creator
		&cmt.Creator.UID, &cmt.Creator.Name, &cmt.Creator.Username,
		&cmt.Creator.Avatar,
		&cmt.CreatedAt, &cmt.LastMod, pq.Array(&cmt.VotePos),
		pq.Array(&cmt.VoteNeg))
	if err != nil {
		return &Comment{}, err
	}
	return cmt, nil
}
