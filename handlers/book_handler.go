package handlers

import (
	"errors"
	"net/http"
	"strings"

	"restapi/models"
	"restapi/repositories"
	"restapi/services"
	"restapi/utils"
)

var bookService = services.NewBookService(repositories.NewBookRepository())

func Books(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetBooks(w, r)
	case http.MethodPost:
		CreateBook(w, r)
	default:
		utils.Error(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func BookByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetBook(w, r)
	case http.MethodPut:
		UpdateBook(w, r)
	case http.MethodDelete:
		DeleteBook(w, r)
	default:
		utils.Error(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	utils.Success(w, http.StatusOK, bookService.GetAll())
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id := idFromPath(r)
	book, err := bookService.GetByID(id)
	if err != nil {
		utils.Error(w, http.StatusNotFound, "book not found")
		return
	}

	utils.Success(w, http.StatusOK, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := utils.ReadJSON(r, &book); err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	created, err := bookService.Create(book)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "name and roll_number are required")
		return
	}

	utils.Success(w, http.StatusCreated, created)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := idFromPath(r)

	var book models.Book
	if err := utils.ReadJSON(r, &book); err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	updated, err := bookService.Update(id, book)
	if errors.Is(err, repositories.ErrBookNotFound) {
		utils.Error(w, http.StatusNotFound, "book not found")
		return
	}
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "name and roll_number are required")
		return
	}

	utils.Success(w, http.StatusOK, updated)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := idFromPath(r)
	if err := bookService.Delete(id); err != nil {
		utils.Error(w, http.StatusNotFound, "book not found")
		return
	}

	utils.Success(w, http.StatusOK, map[string]string{"message": "book deleted"})
}

func idFromPath(r *http.Request) string {
	return strings.TrimPrefix(r.URL.Path, "/api/v1/books/")
}
