package viasendgrid

import (
	"bytes"
	"context"
	htmltemplate "html/template"
	"path/filepath"
	texttemplate "text/template"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGrid struct {
	client   *sendgrid.Client
	from     *mail.Email
	toName   string
	subject  string
	textTmpl *texttemplate.Template
	htmlTmpl *htmltemplate.Template
	timeout  time.Duration
}

type SendGridConfig struct {
	Key      string
	FromName string
	From     string
	ToName   string
	Subject  string
	Timeout  time.Duration
}

type TmplConfig struct {
	TextPath string
	HtmlPath string
}

func New(sc *SendGridConfig, tc *TmplConfig) (*SendGrid, error) {
	_, textFilename := filepath.Split(tc.TextPath)
	textTmpl, err := texttemplate.New(textFilename).ParseFiles(tc.TextPath)
	if err != nil {
		return &SendGrid{}, err
	}

	_, htmlFilename := filepath.Split(tc.HtmlPath)
	htmlTmpl, err := htmltemplate.New(htmlFilename).ParseFiles(tc.HtmlPath)

	return &SendGrid{
		client:   sendgrid.NewSendClient(sc.Key),
		from:     mail.NewEmail(sc.FromName, sc.From),
		toName:   sc.ToName,
		subject:  sc.Subject,
		textTmpl: textTmpl,
		htmlTmpl: htmlTmpl,
		timeout:  sc.Timeout,
	}, nil
}

func (sg *SendGrid) Send(dst, msg string) error {
	textContentBuf := &bytes.Buffer{}
	err := sg.textTmpl.Execute(textContentBuf, map[string]interface{}{
		"Message": msg,
	})
	if err != nil {
		return err
	}

	htmlContentBuf := &bytes.Buffer{}
	err = sg.textTmpl.Execute(htmlContentBuf, map[string]interface{}{
		"Message": msg,
	})
	if err != nil {
		return err
	}

	sgMail := mail.NewSingleEmail(sg.from, sg.subject,
		mail.NewEmail(sg.toName, dst),
		textContentBuf.String(), htmlContentBuf.String())
	ctx, _ := context.WithTimeout(context.Background(), sg.timeout)
	_, err = sg.client.SendWithContext(ctx, sgMail)
	return err
}
