package render

import (
	"html/template"
	"io"
)

type Engine struct {
	dir string
	ext string
}

func New(dir, ext string) *Engine {
	e := &Engine{
		dir: dir,
		ext: ext,
	}
	return e
}

func (e *Engine) Load() error {
	return nil
}

func (e *Engine) Render(w io.Writer, s string, b interface{}, l ...string) error {
	r := template.New(s + e.ext)
	pl := make([]string, len(l))
	for i := range l {
		pl[i] = e.dir + l[i] + e.ext
	}
	tpl, err := r.ParseFiles(pl...)
	if err != nil {
		return err
	}
	return tpl.Execute(w, b)
}
