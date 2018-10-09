package repository

import "sync"

type BookRepository interface {
}

var (
	bookRepo           BookRepository
	onceBookRepository sync.Once
)

func BookRepo() BookRepository {
	onceBookRepository.Do(func() {
		bookRepo = NewBookRepo()
	})
	return sessionRepo
}

func NewBookRepo() BookRepository {
	repo := &bookRepositoryImp{}
	return repo
}

//
// Implement BookRepository
//
type bookRepositoryImp struct {
}
