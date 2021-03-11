package data

import (
	"github.com/go-ushio/ushio/core/category"
	"github.com/go-ushio/ushio/core/comment"
	"github.com/go-ushio/ushio/core/post"
	"github.com/go-ushio/ushio/core/session"
	"github.com/go-ushio/ushio/core/sign_up"
	"github.com/go-ushio/ushio/core/user"
)

type Provider interface {
	PostByPID(pid int64) (*post.Post, error)
	PostInfoByPID(pid int64) (*post.Info, error)
	PostInfoByPage(size, offset int64) ([]*post.Info, error)
	PostInfoIndex(size int64) ([]*post.Info, error)
	InsertPost(p *post.Post) (int64, error)
	InsertPostInfo(info *post.Info) error
	UpdatePost(p *post.Post) error
	UpdatePostTitle(pid int64, title string) error
	UpdatePostLimit(pid, limit int64) error
	PostNewReply(pid int64) error
	PostNewView(pid int64) error
	PostNewMod(pid int64) error
	PostNewActivity(pid int64) error
	PostNewPosVote(pid, uid int64) error
	PostNewNegVote(pid, uid int64) error
	PostedBy(uid int64) ([]*post.Info, error)

	CommentsByPost(pid int64) ([]*comment.Comment, error)
	CommentByCid(cid int64) (*comment.Comment, error)
	InsertComment(cmt *comment.Comment) (int64, error)
	UpdateComment(cmt *comment.Comment) error
	CommentNewMod(cid int64) error
	CommentNewPosVote(cid, uid int64) error
	CommentNewNegVote(cid, uid int64) error

	SessionByToken(token string) (*session.Session, error)
	SessionsByUID(uid int64) ([]*session.Session, error)
	SessionBasicByToken(token string) (*session.Basic, error)
	InsertSession(sess *session.Session) error
	DeleteUserSessions(uid int64) error

	UserByUID(uid int64) (*user.User, error)
	UserByEmail(email string) (*user.User, error)
	UserByUsername(username string) (*user.User, error)
	UserAuthByUID(uid int64) (*user.Auth, error)
	InsertUser(u *user.User) (int64, error)
	InsertUserAuth(auth *user.Auth) error
	UpdateUser(u *user.User) error
	UpdateUserAuth(auth *user.Auth) error
	AddArtifact(uid, add int64) error
	DeleteUser(uid int64) error
	UsernameExists(username string) (bool, error)
	EmailExists(email string) (bool, error)
	//UserFollow(op, target int64) error
	//UserUnFollow(op, target int64) error
	//AlreadyFollow(op, target int64) (bool, error)
	//Followings(uid int64)  ([]*user.User, error)
	//Followers(uid int64)  ([]*user.User, error)

	SignUpByToken(token string) (*sign_up.SignUp, error)
	SignUpByEmail(email string) (*sign_up.SignUp, error)
	InsertSignUp(su *sign_up.SignUp) error
	DeleteSignUpByEmail(email string) error
	SignUpExists(email string) (bool, error)

	GetCategories() ([]*category.Category, error)

	Close() error
}
