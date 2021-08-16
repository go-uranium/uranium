package sendmail

type Sender interface {
	Send(dst, msg string) error
}
