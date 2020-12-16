package recaptcha

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var DefaultSecret = os.Getenv("G_SECRET")

type Request struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
	RemoteIP string `json:"remoteip"`
}

type Response struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func Verify(grrc string) (bool, error) {
	if len(grrc) == 0 {
		return false, nil
	}
	response, err := VerifyFull(&Request{
		Secret:   DefaultSecret,
		Response: grrc,
	})
	if err != nil {
		return false, err
	}
	if response.ErrorCodes != nil {
		if len(response.ErrorCodes) != 0 {
			errs := strings.Join(response.ErrorCodes, ", ")
			return response.Success, errors.New(errs)
		}
	}
	return response.Success, nil
}

func VerifyFull(req *Request) (*Response, error) {
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", req.Values())
	if err != nil {
		return &Response{}, err
	}
	decoder := json.NewDecoder(resp.Body)
	r := &Response{}
	err = decoder.Decode(r)
	return r, err
}

func (req *Request) Values() url.Values {
	v := url.Values{}
	v["secret"] = []string{req.Secret}
	v["response"] = []string{req.Response}
	v["remoteip"] = []string{req.RemoteIP}
	return v
}
