package repository

import (
	"database/sql"

	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	repository "github.com/gui-laranjeira/go-cleanarch/internal/infrastructure/repository/errors"
)

type BookSQLRepository struct {
	db *sql.DB
}

func NewBookSQLRepository(db *sql.DB) entity.IBookRepository {
	return &BookSQLRepository{
		db: db,
	}
}

func (r *BookSQLRepository) Create(book *entity.Book) error {
	sqlStatement := `INSERT INTO books (id_book, title, author, publisher, year, isbn, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	rows, err := r.db.Exec(sqlStatement, book.ID, book.Title, book.Author, book.Publisher, book.Year, book.ISBN, book.CreatedAt)
	if err != nil {
		return err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return repository.NoRowsAffectedError()
	}

	if rowsAffected > 1 {
		return repository.MultipleRowsAffectedError()
	}

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

func (r *BookSQLRepository) FindByAuthor(name string) ([]*entity.Book, error) {
	return nil, nil
}

func (r *BookSQLRepository) FindByTitle(title string) (*entity.Book, error) {
	return nil, nil
}

func (r *BookSQLRepository) FindByPublisher(publisher string) ([]*entity.Book, error) {
	return nil, nil
}

func FindByAuthor(author string) ([]*entity.Book, error) {
	return nil, nil
}

func FindByPublisher(publisher string) ([]*entity.Book, error) {
	return nil, nil
}
