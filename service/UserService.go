package service

import (
	"Bookee/domain/user"
	"Bookee/infra/client/wxclient"
	"sync"
)

type UserService interface {
	Get(uid int64) *user.User
	GetByWxOpenId(openid string) (*user.User, error)
	Login(uid int64, password string) (string, error)
	LoginWX(code string) (string, error)
	GetAccessToken(uid int64) string
}

var (
	userService     UserService
	onceUserService sync.Once
)

func UserSvc() UserService {
	onceUserService.Do(func() {
		userService = NewUserSvc(nil)
	})
	return userService
}

func NewUserSvc(session SessionService) (service UserService) {
	if session != nil {
		service = &userServiceImp{session: session}
	} else {
		service = &userServiceImp{session: SessionSvc()}
	}
	return
}

//
// Implement UserService
//

type userServiceImp struct {
	session SessionService
}

func (userService *userServiceImp) Get(uid int64) *user.User {
	return &user.User{Uid: uid}
}

func (userService *userServiceImp) GetByWxOpenId(openid string) (*user.User, error) {
	return nil, nil
}

func (userService *userServiceImp) LoginWX(code string) (token string, err error) {
	openid, sessionKey, unionid, err := wxclient.Code2Session(code)
	if err == nil {
		user, err := userService.GetByWxOpenId(openid)
		if err == nil {
			token, err = userService.session.UpdateWX(user, openid, sessionKey, unionid)
		}
	}
	return
}

func (userService *userServiceImp) Login(uit int64, passwrod string) (token string, err error) {
	//todo check password

	//todo get or create token

	//todo update session
	return "", nil
}

func (userService *userServiceImp) GetAccessToken(uid int64) string {
	return "tokenValue"
}
