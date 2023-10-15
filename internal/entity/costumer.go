package entity

import (
	"time"

	"github.com/google/uuid"
)

type Costumer struct {
	ID            uuid.UUID  `json:"id_costumer"`
	Email         string     `json:"email"`
	Phone         string     `json:"phone"`
	Address       string     `json:"address"`
	Document      string     `json:"document"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	CurrentBookID *uuid.UUID `json:"current_book_id"`
}

type ICostumerRepository interface {
	Create(costumer *Costumer) error
	Update(costumer *Costumer) (int64, error)
	FindAll() ([]*Costumer, error)
	FindByID(id string) (*Costumer, error)
}

func NewCostumerFactory(email string, phone string, address string, document string, firstName string, lastName string) (*Costumer, error) {
	r := Costumer{
		ID:            uuid.New(),
		Email:         email,
		Phone:         phone,
		Address:       address,
		Document:      document,
		FirstName:     firstName,
		LastName:      lastName,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Time{},
		Books:         make([]uuid.UUID, 5),
		CurrentBookID: nil,
	}

	err := r.validateUser()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return &r, nil
}

func (u *Costumer) validateUser() error {
	if u.Email == "" || u.Phone == "" || u.Address == "" || u.Document == "" || u.FirstName == "" || u.LastName == "" || u.CreatedAt.String() == "" {
		return ErrInvalidEntity
	}
	return nil
}
