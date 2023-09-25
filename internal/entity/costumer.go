package entity

import (
	"time"

	"github.com/google/uuid"
)

type Costumer struct {
	ID            uuid.UUID  `json:"id_costumer"`
	Email         string     `json:"email"`
	Phone         string     `json:"phone"`
	Adress        string     `json:"adress"`
	Document      string     `json:"document"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Books         []*Book    `json:"books"`
	CurrentBookID *uuid.UUID `json:"current_book_id"`
}

type ICostumerRepository interface {
	Create(costumer *Costumer) error
	Update(costumer *Costumer) (int64, error)
	BorrowBook(costumer_id string, book_id string) error
	ReturnBook(costumer_id string, book_id string) error
}

func NewCostumerFactory(email string, phone string, adress string, document string, firstName string, lastName string) (*Costumer, error) {
	r := Costumer{
		ID:            uuid.New(),
		Email:         email,
		Phone:         phone,
		Adress:        adress,
		Document:      document,
		FirstName:     firstName,
		LastName:      lastName,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Time{},
		Books:         make([]*Book, 5),
		CurrentBookID: nil,
	}

	err := r.validateUser()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return &r, nil
}

func (u *Costumer) validateUser() error {
	if u.Email == "" || u.Phone == "" || u.Adress == "" || u.Document == "" || u.FirstName == "" || u.LastName == "" || u.CreatedAt.String() == "" {
		return ErrInvalidEntity
	}
	return nil
}
