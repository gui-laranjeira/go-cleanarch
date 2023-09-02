package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id_user"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Books     []Book    `json:"books"`
}

func NewUser(email string, password string, firstName string, lastName string) (*User, error) {
	err := validadePassword(password)
	if err != nil {
		return nil, ErrInvalidEntity
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrPasswordHashing
	}

	r := User{
		ID:        uuid.New(),
		Email:     email,
		Password:  string(hash),
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
	}

	err = r.validateUser()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return &r, nil
}

func (u *User) validateUser() error {
	if u.Email == "" || u.Password == "" || u.FirstName == "" || u.LastName == "" || u.CreatedAt.String() == "" {
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
