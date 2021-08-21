package comment

import (
	"html/template"
	"time"

	"github.com/go-uranium/uranium/model/user"
)

type Comment struct {
	CID      int32         `json:"cid"`
	PID      int32         `json:"pid"`
	Creator  user.Basic    `json:"creator"`
	Content  template.HTML `json:"content"`
	Modified time.Time     `json:"modified"`
	Created  time.Time     `json:"created"`
}

type Statistic struct {
	VotePos []int32 `json:"vote_pos"`
	VoteNeg []int32 `json:"vote_neg"`
}
