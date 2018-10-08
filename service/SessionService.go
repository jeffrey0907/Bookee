package service

import (
    "Bookee/domain/user"
    "Bookee/infra/repository"
    "github.com/dgrijalva/jwt-go"
    "time"

    "sync"
)

var (
    sessionService     SessionService
    onceSessionService sync.Once
)

const (
    tokenSecret = `123456`
)

type SessionService interface {
    CheckJWT(jwt string) (int64, error)
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
    usr *user.User, openid string, sessionKey string, unionid string) (token string, err error) {
    session := this.sessionRepo.Get(usr.Uid)
    if session == nil {
        session = user.NewSession(usr.Uid, openid, sessionKey, unionid)
        this.sessionRepo.Save(session)
    }
    token, err = this.makeJWT(session)
    return
}

func (this *sessionServiceImp) makeJWT(session *user.Session) (tokenString string, err error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := make(jwt.MapClaims)
    claims["uid"] = session.Uid
    claims["exp"] = session.Expire.Add(time.Hour * 24).Unix()
    claims["iat"] = time.Now().Unix()
    token.Claims = claims

    tokenString, err = token.SignedString([]byte(tokenSecret))
    return
}

func (this *sessionServiceImp) CheckJWT(tokenString string) (uid int64, err error) {
    token, err :=jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(tokenSecret), nil
    })
    if err != nil && token.Valid {
        uid = -1
    }
    return
}
