package routes

import (
	"books/cmd/handlers"
	"net/http"
)

func SetupBookRoutes(bookHandler *handlers.BookHandler) {
	http.HandleFunc("/books/", bookHandler.HandleBooks)
}
