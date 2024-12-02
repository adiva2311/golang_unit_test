package service

import (
	"errors"
	"unit-test/entity"
	"unit-test/repository"
)

type BookService struct {
	Repository repository.BookRepository
}

func (receiver BookService) Get(id int) (*entity.Book, error) {
	book := receiver.Repository.FindById(id)

	if book == nil{
		return nil, errors.New("book not found")
	} else {
		return book, nil
	}
}