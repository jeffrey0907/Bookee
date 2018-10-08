package repository

import (
    "Bookee/domain/user"
    "sync"
)

type SessionRepository interface {
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
type sessionRepoMysql struct {
}

func newSessionRepoMySql() SessionRepository {

}

//
// Implement SessionRepository With Memory
//
type sessionRepoMem struct {
    sessions map[int64]*user.Session
}

func newSessionRepoMem() SessionRepository {
    repo := &sessionRepoMem{}

}
