package service

import (
	"Bookee/domain/book"
	"Bookee/infra/repository"
	"sync"
)

type BookService interface {
	Get(bookId int64) *book.Book
}

var (
	bookService     BookService
	onceBookService sync.Once
)

func BookSvc() BookService {
	onceBookService.Do(func() {
		bookService = NewBookSvc(nil)
	})
	return bookService
}

func NewBookSvc(bookRepo repository.BookRepository) BookService {
	bookSvc := &defaultBookServiceImp{}
	if bookRepo != nil {
		bookSvc.bookRepo = bookRepo
	} else {
		bookSvc.bookRepo = repository.BookRepo()
	}
	return bookSvc
}

//
// Implement BookService
//
type defaultBookServiceImp struct {
	bookRepo repository.BookRepository
}

func (this *defaultBookServiceImp) Get(bookId int64) *book.Book {
	return &book.Book{Id: bookId}
}
