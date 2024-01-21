package usecases_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	usecases "github.com/gui-laranjeira/go-cleanarch/internal/usecases/costumer"
	"github.com/stretchr/testify/mock"
)

type MockCostumerRepository struct {
	mock.Mock
}

func (m *MockCostumerRepository) Create(c *entity.Costumer) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockCostumerRepository) Update(costumer *entity.Costumer) (int64, error) {
	args := m.Called(costumer)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockCostumerRepository) FindAll() ([]*entity.Costumer, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Costumer), args.Error(1)
}

func (m *MockCostumerRepository) FindByID(id string) (*entity.Costumer, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Costumer), args.Error(1)
}
func TestCreateCostumer(t *testing.T) {
	mockRepo := new(MockCostumerRepository)
	u := usecases.NewCreateCostumerUseCase(mockRepo)

	input := usecases.CreateCostumerInput{
		Email:     "test@test.com",
		Phone:     "1234567890",
		Address:   "Test Address",
		Document:  "1234567890",
		FirstName: "Test",
		LastName:  "User",
	}

	mockRepo.On("Create", mock.Anything).Return(nil)

	err := u.CreateCostumer(input)

	if err != nil {
		t.Errorf("CreateCostumer() error = %v", err)
		return
	}

	mockRepo.AssertExpectations(t)
}

func TestUpdateCostumer(t *testing.T) {
	mockRepo := new(MockCostumerRepository)
	u := usecases.NewUpdateCostumerUseCase(mockRepo)

	id := uuid.New().String()
	input := usecases.UpdateCostumerInput{
		ID:        id,
		Email:     "test@test.com",
		Phone:     "1234567890",
		Address:   "Test Address",
		Document:  "1234567890",
		FirstName: "Test",
		LastName:  "User",
	}

	mockRepo.On("Update", mock.Anything).Return(int64(1), nil)

	output, err := u.UpdateCostumer(input)

	if err != nil {
		t.Errorf("UpdateCostumer() error = %v", err)
		return
	}

	if output.RowsAffected != 1 {
		t.Errorf("UpdateCostumer() RowsAffected = %v, want 1", output.RowsAffected)
	}

	mockRepo.AssertExpectations(t)
}

func TestFindAllCostumers(t *testing.T) {
	mockRepo := new(MockCostumerRepository)
	u := usecases.NewFindAllCostumersUseCase(mockRepo)

	id := uuid.New().String()
	costumer := &entity.Costumer{
		ID:        uuid.MustParse(id),
		Email:     "test@test.com",
		Phone:     "1234567890",
		Address:   "Test Address",
		Document:  "1234567890",
		FirstName: "Test",
		LastName:  "User",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("FindAll").Return([]*entity.Costumer{costumer}, nil)

	output, err := u.FindAllCostumers()

	if err != nil {
		t.Errorf("FindAllCostumers() error = %v", err)
		return
	}

	if len(output.Costumers) != 1 {
		t.Errorf("FindAllCostumers() len(Costumers) = %v, want 1", len(output.Costumers))
	}

	mockRepo.AssertExpectations(t)
}

func TestFindCostumerByID(t *testing.T) {
	mockRepo := new(MockCostumerRepository)
	u := usecases.NewFindCostumerByIDUseCase(mockRepo)

	id := uuid.New().String()
	costumer := &entity.Costumer{
		ID:        uuid.MustParse(id),
		Email:     "test@test.com",
		Phone:     "1234567890",
		Address:   "Test Address",
		Document:  "1234567890",
		FirstName: "Test",
		LastName:  "User",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("FindByID", id).Return(costumer, nil)

	output, err := u.FindCostumerByID(usecases.FindCostumerByIDInput{ID: id})

	if err != nil {
		t.Errorf("FindCostumerByID() error = %v", err)
		return
	}

	if output.ID != id {
		t.Errorf("FindCostumerByID() ID = %v, want %v", output.ID, id)
	}

	mockRepo.AssertExpectations(t)
}
