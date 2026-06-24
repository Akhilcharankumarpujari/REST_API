package routes

import (
	"net/http"

	handler "restapi/handlers"
	"restapi/middleware"
)

func Setup(mux *http.ServeMux) {
	books := middleware.Logger(middleware.Auth(http.HandlerFunc(handler.Books)))
	bookByID := middleware.Logger(middleware.Auth(http.HandlerFunc(handler.BookByID)))

	mux.Handle("/api/v1/books", books)
	mux.Handle("/api/v1/books/", bookByID)
}
