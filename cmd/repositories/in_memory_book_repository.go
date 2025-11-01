package repositories

import (
	"books/cmd/models"
	"errors"
)

type InMemoryBookRepository struct {
	books  map[int]models.Book
	nextID int
}

func NewInMemoryBookRepository() *InMemoryBookRepository {
	return &InMemoryBookRepository{
		books:  make(map[int]models.Book),
		nextID: 1,
	}
}

func (r *InMemoryBookRepository) GetAll() ([]models.Book, error) {
	books := make([]models.Book, 0, len(r.books))
	for _, book := range r.books {
		books = append(books, book)
	}
	return books, nil
}

func (r *InMemoryBookRepository) GetByID(id int) (models.Book, error) {
	book, exists := r.books[id]
	if !exists {
		return models.Book{}, errors.New("book not found")
	}
	return book, nil
}

func (r *InMemoryBookRepository) Create(book models.Book) (models.Book, error) {
	book.ID = r.nextID
	r.nextID++
	r.books[book.ID] = book
	return book, nil
}

func (r *InMemoryBookRepository) Update(id int, book models.Book) (models.Book, error) {
	if _, exists := r.books[id]; !exists {
		return models.Book{}, errors.New("book not found")
	}
	book.ID = id
	r.books[id] = book
	return book, nil
}

func (r *InMemoryBookRepository) Delete(id int) error {
	if _, exists := r.books[id]; !exists {
		return errors.New("book not found")
	}
	delete(r.books, id)
	return nil
}
