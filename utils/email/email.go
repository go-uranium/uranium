package email

import (
	"encoding/json"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

type Verify struct {
	EmailAddress string `json:"address"`
	Code         string `json:"code"`
}


type Email struct {
	Email string `json:"email"`
}

type Personalizations struct {
	To []Email `json:"to"`
	DynamicTemplateData *Verify `json:"dynamic_template_data"`
}

type SendGrid struct {
	From Email `json:"from"`
	Personalizations []Personalizations `json:"personalizations"`
	TemplateID string `json:"template_id"`
}

func Send(from string,v *Verify) error {
	sg := &SendGrid{
		From:Email{Email:from},
		Personalizations: []Personalizations{{To: []Email{{Email:v.EmailAddress}},DynamicTemplateData:v}},
		TemplateID: os.Getenv("SENDGRID_TPL_ID"),
	}
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var err error
	request.Body, err = json.Marshal(sg)
	if err != nil {
		return err
	}
	_, err = sendgrid.API(request)
	return err
}
