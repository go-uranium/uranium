package rcache_test

import "github.com/go-uranium/uranium/model/session"

func (*testingStorage) SessionBasicByToken(token string) (*session.Basic, error) {
	return nil, nil
}

func (*testingStorage) SessionInsertSession(sess *session.Session) error {
	return nil
}
