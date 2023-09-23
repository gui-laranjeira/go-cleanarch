package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            uuid.UUID  `json:"id_user"`
	Email         string     `json:"email"`
	Phone         string     `json:"phone"`
	Password      string     `json:"password"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Books         []*Book    `json:"books"`
	CurrentBookID *uuid.UUID `json:"current_book_id"`
}

type IUserRepository interface {
	Create(user *User) error
	Update(newUser *User) (int64, error)
	Login(email string, password string) (*User, error)
	ChangePassword(id string, newPassword string) (int64, error)
	BorrowBook(user *User, book *Book) error
	ReturnBook(user *User, book *Book) error
}

func NewUserFactory(email string, phone string, password string, firstName string, lastName string) (*User, error) {
	err := validadePassword(password)
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrPasswordHashing
	}

	r := User{
		ID:            uuid.New(),
		Email:         email,
		Phone:         phone,
		Password:      string(hash),
		FirstName:     firstName,
		LastName:      lastName,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Time{},
		Books:         make([]*Book, 5),
		CurrentBookID: nil,
	}

	err = r.validateUser()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return &r, nil
}

func (u *User) validateUser() error {
	if u.Email == "" || u.Phone == "" || u.Password == "" || u.FirstName == "" || u.LastName == "" || u.CreatedAt.String() == "" {
		return ErrInvalidEntity
	}
	return nil
}

func validadePassword(password string) error {
	if password == "" {
		return ErrInvalidEntity
	}
	return nil
}
