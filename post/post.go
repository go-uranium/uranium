package post

import (
	"html/template"
	"time"
)

type Post struct {
	PID         int           `json:"pid"`
	Thread      int           `json:"thread"`
	Date        time.Time     `json:"date"`
	Creator     int           `json:"creator"`
	RawContent  string        `json:"-"`
	HTMLContent template.HTML `json:"html_content"`
}
