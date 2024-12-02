package repository

import (
	"unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type BookRepositoryMock struct {
	Mock mock.Mock
}

// FindById implements BookRepositoryIMock.
func (b *BookRepositoryMock) FindById(id int) *entity.Book {
	argument := b.Mock.Called(id)

	if argument.Get(0) == nil {
		return nil
	} else {
		book := argument.Get(0).(entity.Book)
		return &book
	}
}