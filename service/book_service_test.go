package service

import (
	"testing"
	"unit-test/entity"
	"unit-test/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var bookRepository = &repository.BookRepositoryMock{Mock: mock.Mock{}}
var bookService = BookService{Repository: bookRepository}

func TestBookService_Get(t *testing.T) {
	// Program Mock
	bookRepository.Mock.On("FindById", 1).Return(nil)

	book, err := bookService.Get(1)
	assert.Nil(t, book)
	assert.NotNil(t, err)
}

func TestBookService_GetFound(t *testing.T)  {
	book := entity.Book{
		Id : 2,
		Name: "Atomic Habit",
	}
	// Program Mock
	bookRepository.Mock.On("FindById", 2).Return(book)

	result, err := bookService.Get(2)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, book.Id, result.Id)
	assert.Equal(t, book.Name, result.Name)
}