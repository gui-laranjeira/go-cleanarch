package entity

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `json:"id_book"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Pages     int       `json:"pages"`
	Publisher string    `json:"publisher"`
	Year      int       `json:"year"`
	ISBN      string    `json:"isbn"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IBookRepository interface {
	Create(book *Book) error
	Update(book *Book, newBook *Book) error
	FindAll() ([]*Book, error)
	FindByID(id string) (*Book, error)
	FindByTitle(title string) (*Book, error)
	FindByAuthor(author string) ([]*Book, error)
	FindByPublisher(publisher string) ([]*Book, error)
	Delete(id string) (int64, error)
}

type IBookUseCase interface {
	CreateBook(book *Book) error
	UpdateBook(book *Book, newBook *Book) error
	GetBooks() ([]*Book, error)
	GetBookByID(id string) (*Book, error)
	GetBookByTitle(title string) (*Book, error)
	GetBookByAuthor(author string) ([]*Book, error)
	GetBookByPublisher(publisher string) ([]*Book, error)
}

func NewBookFactory(title string, author string, pages int, publisher string, year int, isbn string) (*Book, error) {
	r := Book{
		ID:        uuid.New(),
		Title:     title,
		Author:    author,
		Pages:     pages,
		Publisher: publisher,
		Year:      year,
		ISBN:      isbn,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}
	err := r.ValidateBook()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return &r, nil
}

func (b *Book) ValidateBook() error {
	if b.Title == "" || b.Author == "" || b.Pages <= 0 || b.Publisher == "" || b.Year <= 0 || b.ISBN == "" || b.CreatedAt.String() == "" {
		return ErrInvalidEntity
	}
	return nil
}
