package viasmtp

import (
	"bytes"
	"net/smtp"
	"path/filepath"
	"text/template"
)

type SMTP struct {
	auth smtp.Auth
	tmpl *template.Template
	sc   *SMTPConfig
}

type SMTPConfig struct {
	Identity string
	Username string
	Password string
	Host     string
	// example: "587"
	Port string
	From string
}

type TmplConfig struct {
	TmplPath string
}

func New(sc *SMTPConfig, tc *TmplConfig) (*SMTP, error) {
	_, filename := filepath.Split(tc.TmplPath)
	tmpl, err := template.New(filename).ParseFiles(tc.TmplPath)
	if err != nil {
		return &SMTP{}, err
	}
	return &SMTP{
		auth: smtp.PlainAuth(sc.Identity, sc.Username, sc.Password, sc.Host),
		tmpl: tmpl,
		sc:   sc,
	}, nil
}

func (s *SMTP) Send(dst string, msg string) error {
	mailContentBuf := &bytes.Buffer{}
	err := s.tmpl.Execute(mailContentBuf, map[string]interface{}{
		"Destination": dst,
		"Message":     msg,
		"From":        s.sc.From,
	})
	if err != nil {
		return err
	}
	err = smtp.SendMail(s.sc.Host+":"+s.sc.Port,
		s.auth, s.sc.From, []string{dst}, mailContentBuf.Bytes())
	return err
}
