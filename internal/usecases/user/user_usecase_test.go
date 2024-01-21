package usecases_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	usecases "github.com/gui-laranjeira/go-cleanarch/internal/usecases/user"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Update(newUser *entity.User) (int64, error) {
	args := m.Called(newUser)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockUserRepository) Login(email string, password string) (*entity.User, error) {
	args := m.Called(email, password)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) ChangePassword(id string, newPassword string) (int64, error) {
	args := m.Called(id, newPassword)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockUserRepository) GetAllUsers() ([]*entity.User, error) {
	args := m.Called()
	return args.Get(0).([]*entity.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(id string) (*entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	u := usecases.NewCreateUserUseCase(mockRepo)

	input := usecases.CreateUserInput{
		Email:     "test@test.com",
		Phone:     "1234567890",
		Password:  "password",
		FirstName: "Test",
		LastName:  "User",
	}

	mockRepo.On("Create", mock.Anything).Return(nil)

	err := u.CreateUser(input)

	if err != nil {
		t.Errorf("CreateUser() error = %v", err)
		return
	}

	mockRepo.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	u := usecases.NewUpdateUserUseCase(mockRepo)

	id := uuid.New().String()
	input := usecases.UpdateUserInput{
		ID:        id,
		Email:     "test@test.com",
		Phone:     "1234567890",
		FirstName: "Test",
		LastName:  "User",
	}

	mockRepo.On("Update", mock.Anything).Return(int64(1), nil)

	output, err := u.UpdateUser(input)
	if err != nil {
		t.Errorf("UpdateUser() error = %v", err)
		return
	}

	if output.RowsAffected != 1 {
		t.Errorf("UpdateUser() RowsAffected = %v, want 1", output.RowsAffected)
	}

	mockRepo.AssertExpectations(t)
}

func TestFindAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	u := usecases.NewFindAllUsersUseCase(mockRepo)

	id := uuid.New().String()
	user := &entity.User{
		ID:        uuid.MustParse(id),
		Email:     "test@test.com",
		Phone:     "1234567890",
		Password:  "password",
		FirstName: "Test",
		LastName:  "User",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("GetAllUsers").Return([]*entity.User{user}, nil)

	output, err := u.FindAllUsers()

	if err != nil {
		t.Errorf("FindAllUsers() error = %v", err)
		return
	}

	if len(output.Users) != 1 {
		t.Errorf("FindAllUsers() len(Users) = %v, want 1", len(output.Users))
	}

	mockRepo.AssertExpectations(t)
}

func TestFindUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	u := usecases.NewFindUserByIDUseCase(mockRepo)

	id := uuid.New().String()
	user := &entity.User{
		ID:        uuid.MustParse(id),
		Email:     "test@test.com",
		Phone:     "1234567890",
		Password:  "password",
		FirstName: "Test",
		LastName:  "User",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("GetUserByID", id).Return(user, nil)

	output, err := u.FindUserByID(usecases.FindUserByIDInput{ID: id})

	if err != nil {
		t.Errorf("FindUserByID() error = %v", err)
		return
	}

	if output.ID != id {
		t.Errorf("FindUserByID() ID = %v, want %v", output.ID, id)
	}

	mockRepo.AssertExpectations(t)
}

func TestUpdatePassword(t *testing.T) {
	mockRepo := new(MockUserRepository)
	u := usecases.NewUpdatePasswordUseCase(mockRepo)

	id := uuid.New().String()
	input := usecases.UpdatePasswordInput{
		ID:       id,
		Password: "newpassword",
	}

	mockRepo.On("ChangePassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(1), nil).Run(func(args mock.Arguments) {
		hash := args.Get(1).(string)
		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(input.Password))
		if err != nil {
			t.Errorf("UpdatePassword() error = %v", err)
		}
	})

	output, err := u.UpdatePassword(input)

	if err != nil {
		t.Errorf("UpdatePassword() error = %v", err)
		return
	}

	if output.RowsAffected != 1 {
		t.Errorf("UpdatePassword() RowsAffected = %v, want 1", output.RowsAffected)
	}

	mockRepo.AssertExpectations(t)
}
