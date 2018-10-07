package service

import (
    "Bookee/domain/user"
    "Bookee/infra/repository"
    "github.com/gin-gonic/gin"

    "sync"
)

var (
    sessionService     SessionService
    onceSessionService sync.Once
)

type SessionService interface {
    CheckJWT(c *gin.Context)
    UpdateWX(user *user.User, openid string, sessionKey string, unionid string) (string, error)
}

func SessionSvc() SessionService {
    onceSessionService.Do(func() {
        sessionService = NewSessionSvc(nil)
    })
    return sessionService
}

func NewSessionSvc(sessionRepo repository.SessionRepository) SessionService {
    sessionSvc := &sessionServiceImp{}
    if sessionRepo != nil {
        sessionSvc.sessionRepo = sessionRepo
    } else {
        sessionSvc.sessionRepo = repository.SessionRepo()
    }
    return sessionSvc
}

//
// Implement SessionService
//

type sessionServiceImp struct {
    sessionRepo repository.SessionRepository
}

func (this *sessionServiceImp) UpdateWX(
    user *user.User, openid string, sessionKey string, unionid string) (token string, err error) {
    this.GetWXSession(user)
    return
}

func (this *sessionServiceImp) CheckJWT(c *gin.Context) {

}
