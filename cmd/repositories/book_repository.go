package repositories

import (
	"books/cmd/models"
)

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetByID(id int) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(id int, book models.Book) (models.Book, error)
	Delete(id int) error
}
