package data

import (
	"github.com/go-ushio/ushio/core/category"
	"github.com/go-ushio/ushio/core/post"
	"github.com/go-ushio/ushio/core/session"
	"github.com/go-ushio/ushio/core/sign_up"
	"github.com/go-ushio/ushio/core/user"
)

type Provider interface {
	PostByPID(pid int) (*post.Post, error)
	PostInfoByPID(pid int) (*post.Info, error)
	PostInfoByPage(size, offset int) ([]*post.Info, error)
	PostInfoIndex(size int) ([]*post.Info, error)
	InsertPost(p *post.Post) (int, error)
	InsertPostInfo(info *post.Info) error
	UpdatePost(p *post.Post) error
	UpdatePostTitle(pid int, title string) error
	UpdatePostLimit(pid, limit int) error
	PostNewReply(pid int) error
	PostNewView(pid int) error
	PostNewMod(pid int) error
	PostNewActivity(pid int) error
	PostNewPosVote(pid, uid int) error
	PostNewNegVote(pid, uid int) error

	SessionByToken(token string) (*session.Session, error)
	SessionsByUID(uid int) ([]*session.Session, error)
	SessionBasicByToken(token string) (*session.Basic, error)
	InsertSession(sess *session.Session) error
	DeleteUserSessions(uid int) error

	UserByUID(uid int) (*user.User, error)
	UserByEmail(email string) (*user.User, error)
	UserByUsername(username string) (*user.User, error)
	UserAuthByUID(uid int) (*user.Auth, error)
	InsertUser(u *user.User) (int, error)
	InsertUserAuth(auth *user.Auth) error
	UpdateUser(u *user.User) error
	UpdateUserAuth(auth *user.Auth) error
	AddArtifact(uid, add int) error
	DeleteUser(uid int) error
	UsernameExists(username string) (bool, error)
	EmailExists(email string) (bool, error)
	UserFollow(op, target int) error
	UserUnFollow(op, target int) error

	SignUpByToken(token string) (*sign_up.SignUp, error)
	SignUpByEmail(email string) (*sign_up.SignUp, error)
	InsertSignUp(su *sign_up.SignUp) error
	DeleteSignUpByEmail(email string) error
	SignUpExists(email string) (bool, error)

	GetCategories() ([]*category.Category, error)
}
