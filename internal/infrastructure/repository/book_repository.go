package repository

import (
	"database/sql"

	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
)

type BookSQLRepository struct {
	db *sql.DB
}

func (r *BookSQLRepository) Create(book *entity.Book) error {
	return nil
}

func (r *BookSQLRepository) Update(book *entity.Book, newBook *entity.Book) error {
	return nil
}

func (r *BookSQLRepository) FindAll() ([]*entity.Book, error) {
	return nil, nil
}

func (r *BookSQLRepository) FindByID(id string) (*entity.Book, error) {
	return nil, nil
}

func (r *BookSQLRepository) FindByTitle(title string) (*entity.Book, error) {
	return nil, nil
}

func FindByAuthor(author string) ([]*entity.Book, error) {
	return nil, nil
}

func FindByPublisher(publisher string) ([]*entity.Book, error) {
	return nil, nil
}
