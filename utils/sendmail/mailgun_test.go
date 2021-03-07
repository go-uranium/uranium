package sendmail_test

//
//import (
//	"html/template"
//	"os"
//	"testing"
//
//	"github.com/mailgun/mailgun-go/v3"
//
//	"github.com/go-ushio/ushio/utils/sendmail"
//	"github.com/go-ushio/ushio/utils/token"
//)
//
//func TestMailGun_Send(t *testing.T) {
//	textEx, err := template.New("email.txt").ParseFiles("views/email.txt")
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	htmlEx, err := template.New("email.html").ParseFiles("views/email.html")
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	mg := &sendmail.MailGun{
//		Client:  mailgun.NewMailgun("ushio.zincic.com", os.Getenv("MAILGUN_API_SEC")),
//		Sender:  "no-reply@ushio.zincic.com",
//		Subject: "email test",
//		Text:    textEx,
//	}
//
//	err = mg.Send("i@iochen.com", token.New())
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	mg.Html = htmlEx
//
//	err = mg.Send("i@iochen.com", token.New())
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//}
