package sendmail

import (
	"bytes"
	"io"
)

type Executor interface {
	Execute(wr io.Writer, data interface{}) error
}

type Map map[string]interface{}

func ExecuteToBuf(ex Executor, data interface{}) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	err := ex.Execute(buf, data)
	if err != nil {
		return &bytes.Buffer{}, err
	}
	return buf, nil
}
