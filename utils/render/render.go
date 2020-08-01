package render

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
)

type Engine struct {
	Dir string
	PDir string
	Partials []string
	Ext string
}

func New(dir, pDir,ext string) *Engine {
	return &Engine{
		Dir:dir,
		PDir:pDir,
		Partials: []string{},
		Ext:ext,
	}
}

func (e *Engine)Load() error {
	if err := filepath.Walk(e.PDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		e.Partials = append(e.Partials, path)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (e *Engine)Render(out io.Writer, tpl string, binding interface{}, layout ...string) error {
	var err error
	var tmpl *template.Template

	tpl = tpl+e.Ext
	if tmpl, err = template.New(tpl).ParseFiles(append(e.Partials, filepath.Join(e.Dir,tpl))...); err != nil {
		return err
	}

	if err = tmpl.Execute(out,binding); err != nil {
		return err
	}
	return nil
}