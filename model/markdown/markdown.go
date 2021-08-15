package markdown

const (
	TYPE_POST uint8 = iota
	TYPE_COMMENT
)

type Markdown struct {
	// 0 stands for post, 1 stands for comment
	Type uint8 `json:"type"`
	Content string `json:"content"`
}

//func (md *Markdown)Copy() *Markdown {
//	md2 := *md
//	return &md2
//}
