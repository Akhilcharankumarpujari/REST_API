package services

import (
	"errors"
	"strings"

	"restapi/models"
	"restapi/repositories"
)

type BookService struct {
	repo *repositories.BookRepository
}

func NewBookService(repo *repositories.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAll() []models.Book {
	return s.repo.FindAll()
}

func (s *BookService) GetByID(id string) (*models.Book, error) {
	return s.repo.FindByID(id)
}

func (s *BookService) Create(book models.Book) (models.Book, error) {
	book.Name = strings.TrimSpace(book.Name)

	if book.Name == "" || book.RollNumber <= 0 {
		return models.Book{}, ErrInvalidBook
	}

	return s.repo.Create(book), nil
}

func (s *BookService) Update(id string, book models.Book) (*models.Book, error) {
	book.Name = strings.TrimSpace(book.Name)

	if book.Name == "" || book.RollNumber <= 0 {
		return nil, ErrInvalidBook
	}

	return s.repo.Update(id, book)
}

func (s *BookService) Delete(id string) error {
	return s.repo.Delete(id)
}

var ErrInvalidBook = errors.New("invalid book")
