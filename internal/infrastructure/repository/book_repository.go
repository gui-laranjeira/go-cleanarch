package repository

import (
	"database/sql"
	"strings"

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
	sqlStatement := `INSERT INTO books (id_book, title, author, pages, publisher, year, isbn, created_at, updated_at) 
						  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	rows, err := r.db.Exec(sqlStatement, book.ID, book.Title, book.Author, book.Pages, book.Publisher, book.Year,
		book.ISBN, book.CreatedAt, book.UpdatedAt)
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

func (r *BookSQLRepository) Update(newBook *entity.Book) (int64, error) {
	sqlStatement := `UPDATE books 
						SET title = $1, author = $2, pages = $3, publisher = $4, year = $5, isbn = $6, updated_at = $7 
					  WHERE id_book=$8`

	rows, err := r.db.Exec(sqlStatement, newBook.Title, newBook.Author, newBook.Pages, newBook.Publisher, newBook.Year, newBook.ISBN, newBook.UpdatedAt, newBook.ID)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *BookSQLRepository) FindAll() ([]*entity.Book, error) {
	sqlStatement := `SELECT * FROM books`
	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*entity.Book
	for rows.Next() {
		var book entity.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Pages, &book.Publisher, &book.Year, &book.ISBN, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}

func (r *BookSQLRepository) FindByID(id string) (*entity.Book, error) {
	sqlStatement := `SELECT * FROM books WHERE id_book = $1`
	rows, err := r.db.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}

	var book entity.Book
	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Pages, &book.Publisher, &book.Year, &book.ISBN, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &book, nil
}

func (r *BookSQLRepository) FindByAuthor(author string) ([]*entity.Book, error) {
	sqlStatement := `SELECT * FROM books WHERE UPPER(author) LIKE UPPER($1)`
	author = "%" + strings.ToUpper(author) + "%"
	rows, err := r.db.Query(sqlStatement, author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*entity.Book
	for rows.Next() {
		var book entity.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Pages, &book.Publisher, &book.Year, &book.ISBN, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}

func (r *BookSQLRepository) FindByTitle(title string) ([]*entity.Book, error) {
	sqlStatement := `SELECT * FROM books WHERE UPPER(title) LIKE UPPER($1)`
	title = "%" + strings.ToUpper(title) + "%"
	rows, err := r.db.Query(sqlStatement, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*entity.Book
	for rows.Next() {
		var book entity.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Pages, &book.Publisher, &book.Year, &book.ISBN, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}

func (r *BookSQLRepository) FindByPublisher(publisher string) ([]*entity.Book, error) {
	sqlStatement := `SELECT * FROM books WHERE UPPER(publisher) LIKE UPPER($1)`
	publisher = "%" + strings.ToUpper(publisher) + "%"
	rows, err := r.db.Query(sqlStatement, publisher)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*entity.Book
	for rows.Next() {
		var book entity.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Pages, &book.Publisher, &book.Year, &book.ISBN, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}

func (r *BookSQLRepository) Delete(id string) (int64, error) {
	sqlStatement := `DELETE FROM books WHERE id_book = $1 RETURNING *`
	result, err := r.db.Exec(sqlStatement, id)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}
