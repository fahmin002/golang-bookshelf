package service

import (
	"github.com/fahminbl/bookshelf/internal/model"

	"github.com/fahminbl/bookshelf/internal/repository"
)

type BookService struct {
	Repo *repository.BookRepository
}

func (s *BookService) GetAllBooks() ([]model.Book, error) {
	return s.Repo.GetAll()
}
