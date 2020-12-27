package post

import (
	"encoding/json"
	"html/template"
	"time"
)

type Post struct {
	PID       int           `json:"pid"`
	Title     string        `json:"title"`
	Creator   int           `json:"creator"`
	Content   template.HTML `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
	LastMod   time.Time     `json:"last_mod"`
	Hidden    bool          `json:"hidden"`
	Anonymous bool          `json:"-"`
	Markdown  string        `json:"markdown"`
}

func (p *Post) Json() ([]byte, error) {
	post := Post{
		PID:       p.PID,
		Title:     p.Title,
		Creator:   p.Creator,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		LastMod:   p.LastMod,
		Hidden:    p.Hidden,
		Markdown:  p.Markdown,
	}
	if p.Anonymous {
		post.Creator = 2
	}
	return json.Marshal(post)
}
