package user

import (
	"time"
)

const sessionDuration = time.Hour * 24

type Session struct {
	Uid      int64
	AccToken string
	Expire   time.Time
	Status   int
}

func NewSession(uid int64) (session Session) {
	session.Uid = uid
	session.AccToken = ""
	session.Expire = time.Now()
	return
}

func (s *Session) IsExpired() bool {
	return time.Now().Sub(s.Expire) > sessionDuration
}

func (s *Session) CheckToken(token string) bool {
	return s.AccToken == token
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

func (s *Session) RefreshToken() {

}
