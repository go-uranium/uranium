package sendmail

import (
	"fmt"
	"net/smtp"
)

type SMTPClient struct {
	From     string
	Password string
	Host     string
	Port     string
	Subject  string
	Text     Executor
	//Html    Executor
}

func (sc *SMTPClient) Send(dst, token string) error {
	textBuf, err := ExecuteToBuf(sc.Text, Map{
		"Dst":     dst,
		"Token":   token,
		"Subject": sc.Subject,
		"From":    sc.From,
	})
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", sc.From, sc.Password, sc.Host)

	fmt.Println(textBuf.Bytes())

	err = smtp.SendMail(sc.Host+":"+sc.Port, auth,
		sc.From, []string{dst}, textBuf.Bytes())
	fmt.Println(err)
	return err
}
