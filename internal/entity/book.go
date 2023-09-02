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

func NewBook(title string, author string, pages int, publisher string, year int, isbn string) (*Book, error) {
	r := Book{
		ID:        uuid.New(),
		Title:     title,
		Author:    author,
		Pages:     pages,
		Publisher: publisher,
		Year:      year,
		ISBN:      isbn,
		CreatedAt: time.Now(),
	}
	err := r.validateBook()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return &r, nil
}

func (b *Book) validateBook() error {
	if b.Title == "" || b.Author == "" || b.Pages <= 0 || b.Publisher == "" || b.Year <= 0 || b.ISBN == "" || b.CreatedAt.String() == "" {
		return ErrInvalidEntity
	}
	return nil
}
