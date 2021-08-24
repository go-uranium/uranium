package storage

import (
	"github.com/pkg/errors"

	"github.com/go-uranium/uranium/model/category"
	"github.com/go-uranium/uranium/model/session"
	"github.com/go-uranium/uranium/model/user"
	"github.com/go-uranium/uranium/utils/sqlnull"
)

var ErrAlreadyExists = errors.New("database already exists")

type Provider interface {
	//PostByPID(pid int64) (*post.Post, error)
	//PostInfoByPID(pid int64) (*post.Info, error)
	//PostsInfoByActivity(hidden bool, size, offset int64) ([]*post.Info, error)
	//PostsInfoByCategory(hidden bool, size, offset, category int64) ([]*post.Info, error)
	//PostsInfoByPID(hidden bool, size, offset int64) ([]*post.Info, error)
	//PostsInfoByCommentCreator(size, offset, uid int64) ([]*post.Info, error)
	//PostsInfoByUID(size, offset, uid int64) ([]*post.Info, error)
	//InsertPost(p *post.Post) (int64, error)
	//InsertPostInfo(info *post.Info) error
	//UpdatePost(p *post.Post) error
	//UpdatePostTitle(pid int64, title string) error
	//UpdatePostLimit(pid, limit int64) error
	//PostNewReply(pid int64) error
	//PostNewView(pid int64) error
	//PostNewMod(pid int64) error
	//PostNewActivity(pid int64) error
	//PostNewPosVote(pid, uid int64) error
	//PostNewNegVote(pid, uid int64) error
	//PostRemovePosVote(pid, uid int64) error
	//PostRemoveNegVote(pid, uid int64) error
	//
	//CommentsByPost(pid int64) ([]*comment.Comment, error)
	//CommentByCID(cid int64) (*comment.Comment, error)
	//CommentByUID(uid int64) ([]*comment.Comment, error)
	//InsertComment(cmt *comment.Comment) (int64, error)
	//UpdateComment(cmt *comment.Comment) error
	//CommentNewMod(cid int64) error
	//CommentNewPosVote(cid, uid int64) error
	//CommentNewNegVote(cid, uid int64) error
	//
	SessionBasicByToken(token string) (*session.Basic, error)
	//SessionsByUID(uid int64) ([]*session.Session, error)
	//SessionBasicByToken(token string) (*session.Basic, error)
	SessionInsertSession(sess *session.Session) error
	//DeleteUserSessions(uid int64) error

	//// user insert
	UserInsertUser(u *user.User) (int32, error)
	UserInsertUserAuth(auth *user.Auth) error
	UserInsertUserProfile(profile *user.Profile) error
	// user query
	UserByUID(uid int32) (*user.User, error)
	UserBasicByUID(uid int32) (*user.Basic, error)
	UserProfileByUID(uid int32) (*user.Profile, error)
	UserAuthByUID(uid int32) (*user.Auth, error)
	UserByUsername(username string) (*user.User, error)
	UserByEmail(email string) (*user.User, error)
	UserBasicByUsername(username string) (*user.Basic, error)
	UserUIDByUsername(username string) (int32, error)
	UserUsernameExists(username string) (bool, error)
	UserEmailExists(email string) (bool, error)

	UserUpdateUsername(uid int32, username string) error
	UserUpdateEmail(uid int32, email string) error
	UserUpdatePassword(uid int32, hashed []byte) error
	UserUpdateProfile(uid int32, profile *user.Profile) error
	UserUpdateSecurityEmail(uid int32, se sqlnull.String) error
	UserUpdateLocked(uid int32, locked bool, till sqlnull.Time) error
	UserUpdateDisabled(uid int32, disabled bool) error
	UserUpdateElectrons(uid int32, electrons int32) error
	UserUpdateDeltaElectrons(uid int32, delta int32) error
	UserUpdateAdmin(uid int32, admin int16) error

	//// user update
	//UpdateUser(u *user.User) error
	//UpdateUserAuth(auth *user.Auth) error
	//UpdateUserProfile(profile *user.Profile) error
	//// user update shortcuts
	//AddElectrons(uid, add int32) error
	//// delete
	//DeleteUser(uid int32) error
	//// check if exist

	//UserFollow(op, target int64) error
	//UserUnFollow(op, target int64) error
	//AlreadyFollow(op, target int64) (bool, error)
	//Followings(uid int64)  ([]*user.User, error)
	//Followers(uid int64)  ([]*user.User, error)

	//SignUpByToken(token string) (*sign_up.SignUp, error)
	//SignUpByEmail(email string) (*sign_up.SignUp, error)
	//InsertSignUp(su *sign_up.SignUp) error
	//DeleteSignUpByEmail(email string) error
	//SignUpExists(email string) (bool, error)
	//
	Categories() ([]*category.Category, error)

	Init() error
	Close() error
}
