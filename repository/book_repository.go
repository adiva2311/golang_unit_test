package repository

import "unit-test/entity"

type BookRepository interface {
	FindById(id int) *entity.Book
}