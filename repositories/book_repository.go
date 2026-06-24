package repositories

import (
	"errors"
	"strconv"
	"sync"

	"restapi/models"
)

var ErrBookNotFound = errors.New("book not found")

type BookRepository struct {
	mu     sync.RWMutex
	books  map[string]models.Book
	nextID int
}

func NewBookRepository() *BookRepository {
	return &BookRepository{
		books: map[string]models.Book{
			"1": {ID: "1", Name: "Manasa", RollNumber: 101},
			"2": {ID: "2", Name: "Rahul", RollNumber: 102},
		},
		nextID: 3,
	}
}

func (r *BookRepository) FindAll() []models.Book {
	r.mu.RLock()
	defer r.mu.RUnlock()

	books := make([]models.Book, 0, len(r.books))
	for _, book := range r.books {
		books = append(books, book)
	}
	return books
}

func (r *BookRepository) FindByID(id string) (*models.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	book, ok := r.books[id]
	if !ok {
		return nil, ErrBookNotFound
	}
	return &book, nil
}

func (r *BookRepository) Create(book models.Book) models.Book {
	r.mu.Lock()
	defer r.mu.Unlock()

	book.ID = strconv.Itoa(r.nextID)
	r.nextID++
	r.books[book.ID] = book
	return book
}

func (r *BookRepository) Update(id string, book models.Book) (*models.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.books[id]; !ok {
		return nil, ErrBookNotFound
	}

	book.ID = id
	r.books[id] = book
	return &book, nil
}

func (r *BookRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.books[id]; !ok {
		return ErrBookNotFound
	}

	delete(r.books, id)
	return nil
}
