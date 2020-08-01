package email

import "testing"

func TestSend(t *testing.T) {
	err := Send("no-reply@ushio.iochen.com",&Verify{
		EmailAddress:"i@iochen.com",
		Code:"TEST CODE",
	})
	if err != nil {
		t.Error(err)
	}
}
