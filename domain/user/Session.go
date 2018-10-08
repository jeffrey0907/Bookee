package user

import (
    "time"
)

const sessionDuration = time.Hour * 24

type Session struct {
    Uid          int64
    OpenId       string
    WxSessionKey string
    UnionId      string
    Expire       time.Time
}

func NewSession(uid int64, openid string, sessionkey string, unionid string) (session *Session) {
    session = &Session{}
    session.Uid = uid
    session.OpenId = openid
    session.WxSessionKey = sessionkey
    session.UnionId = unionid
    session.Expire = time.Now()
    return
}

func (s *Session) IsExpired() bool {
    return time.Now().Sub(s.Expire) > sessionDuration
}

func (s *Session) CheckToken(token string) bool {
    return s.OpenId == token
}

func (s *Session) Check(token string) bool {
    if s.IsExpired() {
        return false
    }
    if s.CheckToken(token) {
        s.Expire = time.Now()
        return true
    } else {
        return false
    }
}
