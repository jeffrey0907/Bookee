package repository

import (
    "Bookee/domain/user"
    "Bookee/infra/myerr"
    "sync"
)

type SessionRepository interface {
    Get(uid int64) *user.Session
    Save(session *user.Session) error
    Delete(uid int64) (*user.Session, error)
}

var (
    sessionRepo           SessionRepository
    onceSeesionRepository sync.Once
)

func SessionRepo() SessionRepository {
    onceSeesionRepository.Do(func() {
        sessionRepo = NewSessionRepo()
    })
    return sessionRepo
}

func NewSessionRepo() SessionRepository {
    //repo := newSessionRepoMySql()
    repo := newSessionRepoMem()
    return repo
}

//
// Implement SessionRepository With MySql
//
func newSessionRepoMySql() SessionRepository {
    return &sessionRepoMysql{}
}

type sessionRepoMysql struct {
}

func (this *sessionRepoMysql) Get(uid int64) *user.Session {
    return nil
}

func (this *sessionRepoMysql) Save(session *user.Session) error {
    return nil
}

func (this *sessionRepoMysql) Delete(uid int64) (*user.Session, error) {
    return nil, nil
}

//
// Implement SessionRepository With Memory
//
func newSessionRepoMem() SessionRepository {
    repo := &sessionRepoMem{}
    repo.sessions = make(map[int64]*user.Session)
    return repo
}

type sessionRepoMem struct {
    sessions map[int64]*user.Session
}

func (this *sessionRepoMem) Get(uid int64) *user.Session {
    if se, ok := this.sessions[uid]; ok {
        return se
    } else {
        return nil
    }
}

func (this *sessionRepoMem) Save(session *user.Session) error {
    if session != nil {
        this.sessions[session.Uid] = session
        return nil
    } else {
        return myerr.ErrIllegalArgument
    }
}

func (this *sessionRepoMem) Delete(uid int64) (*user.Session, error) {
    if session, ok := this.sessions[uid]; ok {
        delete(this.sessions, uid)
        return session, nil
    } else {
        return nil, myerr.ErrIllegalArgument
    }
}
