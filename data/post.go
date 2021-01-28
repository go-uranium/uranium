package data

import (
	"strconv"

	"github.com/go-ushio/ushio/common/put"
	"github.com/go-ushio/ushio/core/post"
)

func (data *Data) PostInfoByPage(offset, size int) ([]*post.Info, error) {
	row, err := data.db.Query(data.sentence.SQLPostInfoByPage, offset, size)
	if err != nil {
		return nil, err
	}
	var postInfoList []*post.Info
	for row.Next() {
		postInfo, err := post.ScanInfo(row)
		if err != nil {
			return nil, err
		}
		postInfoList = append(postInfoList, postInfo)
	}
	return postInfoList, nil
}

func (data *Data) PostInfoByPID(pid int) (*post.Info, error) {
	row := data.db.QueryRow(data.sentence.SQLPostByPID, strconv.Itoa(pid))
	info, err := post.ScanInfo(row)
	if err != nil {
		return &post.Info{}, err
	}
	return info, nil
}

func (data *Data) PostByPID(pid int) (*post.Post, error) {
	row := data.db.QueryRow(data.sentence.SQLPostByPID, strconv.Itoa(pid))
	p, err := post.ScanPost(row)
	if err != nil {
		return &post.Post{}, err
	}
	info, err := data.PostInfoByPID(pid)
	if err != nil {
		return &post.Post{}, err
	}
	p.Info = info
	return p, nil
}

func (data *Data) InsertPostInfo(info *post.Info) error {
	putter := put.PutterFromDBExec(data.db, data.sentence.SQLInsertPostInfo)
	_, err := info.PutWithPIDFirst(putter)
	return err
}

func (data *Data) InsertPost(p *post.Post) error {
	putter := put.PutterFromDBExec(data.db, data.sentence.SQLInsertPost)
	result, err := p.Put(putter)
	if err != nil {
		return err
	}
	pid, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.Info.PID = int(pid)
	err = data.InsertPostInfo(p.Info)
	return err
}

func (data *Data) UpdatePost(p *post.Post) error {
	putter := put.PutterFromDBExec(data.db, data.sentence.SQLUpdatePost)
	_, err := p.PutWithPID(putter)
	if err != nil {
		return err
	}
	err = data.UpdatePostInfo(p.Info)
	return err
}

func (data *Data) UpdatePostInfo(info *post.Info) error {
	putter := put.PutterFromDBExec(data.db, data.sentence.SQLUpdatePostInfo)
	_, err := info.PutWithPIDLast(putter)
	return err
}
