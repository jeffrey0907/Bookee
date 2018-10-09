package repository

import (
    "Bookee/domain/user"
    "sync"
)

type UserRepository interface {
    GetByUid(uid int64) *user.User
    GetByOpenId(openId string) *user.User
}

var (
    userRepo           UserRepository
    onceUserRepository sync.Once
)

func UserRepo() UserRepository {
    onceUserRepository.Do(func() {
        userRepo = NewUserRepo()
    })
    return userRepo
}

func NewUserRepo() UserRepository {
    repo := &userRepositoryImp{}
    return repo
}

//
// Implement UserRepository
//
type userRepositoryImp struct {
}

func (userRepo *userRepositoryImp) GetByUid(uid int64) *user.User {

}

func (userRepo *userRepositoryImp) GetByOpenId(openId string) *user.User {

}
