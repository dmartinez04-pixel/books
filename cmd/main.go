package main

import (
	"books/cmd/handlers"
	"books/cmd/repositories"
	"books/cmd/routes"
	"log"
	"net/http"
)

func main() {
	repo, err := repositories.NewSqliteBookRepository("books.db")
	if err != nil {
		log.Fatal(err)
	}
	bookHandler := handlers.NewBookHandler(repo)
	routes.SetupBookRoutes(bookHandler)

	log.Println("Server starting on :3030")
	log.Fatal(http.ListenAndServe(":3030", nil))
}
