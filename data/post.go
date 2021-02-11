package data

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/lib/pq"

	"github.com/go-ushio/ushio/core/post"
)

func (data *Data) PostByPID(pid int) (*post.Post, error) {
	row := data.db.QueryRow(data.sentence.SQLPostByPID, pid)
	p, err := post.ScanPost(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &post.Post{}, err
	}
	info, err := data.PostInfoByPID(pid)
	if err != nil {
		return &post.Post{}, err
	}
	p.Info = info
	return p, nil
}

func (data *Data) PostInfoByPID(pid int) (*post.Info, error) {
	row := data.db.QueryRow(data.sentence.SQLPostInfoByPID, strconv.Itoa(pid))
	info, err := post.ScanInfo(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &post.Info{}, err
	}
	return info, nil
}

func (data *Data) PostInfoByPage(size, offset int) ([]*post.Info, error) {
	row, err := data.db.Query(data.sentence.SQLPostInfoByPage, size, offset)
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

func (data *Data) PostInfoIndex(size int) ([]*post.Info, error) {
	row, err := data.db.Query(data.sentence.SQLPostInfoIndex, size)
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

func (data *Data) InsertPost(p *post.Post) (int, error) {
	pid := 0
	err := data.db.QueryRow(data.sentence.SQLInsertPost, p.Content, p.Markdown).
		Scan(&pid)
	if err != nil {
		return 0, err
	}
	p.Info.PID = pid
	return pid, data.InsertPostInfo(p.Info)
}

func (data *Data) InsertPostInfo(info *post.Info) error {
	_, err := data.db.Exec(data.sentence.SQLInsertPostInfo, info.PID, info.Title, info.Creator,
		info.CreatedAt, info.LastMod, info.Replies, info.Views,
		info.Activity, info.Hidden, pq.Array(info.VotePos),
		pq.Array(info.VoteNeg), info.Limit)
	return err
}

func (data *Data) UpdatePost(p *post.Post) error {
	_, err := data.db.Exec(data.sentence.SQLUpdatePost, p.PID, p.Content, p.Markdown)
	if err != nil {
		return err
	}
	err = data.InsertPostInfo(p.Info)
	return err
}

func (data *Data) UpdatePostTitle(pid int, title string) error {
	_, err := data.db.Exec(data.sentence.SQLUpdatePostTitle, pid, title)
	return err
}

func (data *Data) UpdatePostLimit(pid, limit int) error {
	_, err := data.db.Exec(data.sentence.SQLUpdatePostLimit, pid, limit)
	return err
}

func (data *Data) PostNewReply(pid int) error {
	_, err := data.db.Exec(data.sentence.SQLPostNewReply, pid)
	return err
}

func (data *Data) PostNewView(pid int) error {
	_, err := data.db.Exec(data.sentence.SQLPostNewView, pid)
	return err
}

func (data *Data) PostNewMod(pid int) error {
	_, err := data.db.Exec(data.sentence.SQLPostNewMod, pid, time.Now())
	return err
}

func (data *Data) PostNewActivity(pid int) error {
	_, err := data.db.Exec(data.sentence.SQLPostNewActivity, pid, time.Now())
	return err
}

func (data *Data) PostNewPosVote(pid, uid int) error {
	_, err := data.db.Exec(data.sentence.SQLPostNewPosVote, pid, uid)
	return err
}

func (data *Data) PostNewNegVote(pid, uid int) error {
	_, err := data.db.Exec(data.sentence.SQLPostNewNegVote, pid, uid)
	return err
}
