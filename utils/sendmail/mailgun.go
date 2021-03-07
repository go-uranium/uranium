package sendmail

//
//import (
//	"context"
//	"time"
//
//	"github.com/mailgun/mailgun-go/v3"
//)
//
//type MailGun struct {
//	Client  *mailgun.MailgunImpl
//	Sender  string
//	Subject string
//	Text    Executor
//	Html    Executor
//}
//
//func (mg *MailGun) Send(dst, token string) error {
//	textBuf, err := ExecuteToBuf(mg.Text, Map{
//		"Dst":   dst,
//		"Token": token,
//	})
//	if err != nil {
//		return err
//	}
//
//	msg := mg.Client.NewMessage(
//		mg.Sender,
//		mg.Subject,
//		textBuf.String(),
//		dst)
//
//	if mg.Html != nil {
//		htmlBuf, err := ExecuteToBuf(mg.Html, Map{
//			"Dst":   dst,
//			"Token": token,
//		})
//		if err != nil {
//			return err
//		}
//		msg.SetHtml(htmlBuf.String())
//	}
//
//
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
//	defer cancel()
//
//	_, _, err = mg.Client.Send(ctx, msg)
//
//	return err
//}
