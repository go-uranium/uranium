package thread

import (
	"html/template"
	"time"
)

type Thread struct {
	TID         int           `json:"tid"`
	Date        time.Time     `json:"date"`
	Creator     int           `json:"creator"`
	Title       string        `json:"title"`
	RawContent  string        `json:"-"`
	HTMLContent template.HTML `json:"html_content"`
}
