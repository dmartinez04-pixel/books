package handlers

import (
	"books/cmd/models"
	"books/cmd/repositories"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct {
	repo repositories.BookRepository
}

func NewBookHandler(repo repositories.BookRepository) *BookHandler {
	return &BookHandler{repo: repo}
}

func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/books/" {
		if r.Method == "GET" {
			h.getAllBooks(w, r)
		} else if r.Method == "POST" {
			h.createBook(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else {
		// /books/{id}
		parts := strings.Split(strings.Trim(path, "/"), "/")
		if len(parts) != 2 || parts[0] != "books" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case "GET":
			h.getBook(w, r, id)
		case "PUT":
			h.updateBook(w, r, id)
		case "DELETE":
			h.deleteBook(w, r, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (h *BookHandler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) getBook(w http.ResponseWriter, r *http.Request, id int) {
	book, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) createBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	createdBook, err := h.repo.Create(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}

func (h *BookHandler) updateBook(w http.ResponseWriter, r *http.Request, id int) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	updatedBook, err := h.repo.Update(id, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

func (h *BookHandler) deleteBook(w http.ResponseWriter, r *http.Request, id int) {
	err := h.repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
