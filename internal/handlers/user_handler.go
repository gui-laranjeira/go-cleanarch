package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gui-laranjeira/go-cleanarch/internal/infrastructure/db"
	"github.com/gui-laranjeira/go-cleanarch/internal/infrastructure/repository"
	usecases "github.com/gui-laranjeira/go-cleanarch/internal/usecases/user"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	defer db.Close()

	repo := repository.NewUserSQLRepository(db)
	uc := usecases.NewCreateUserUseCase(repo)

	var input usecases.CreateUserInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while decoding user data: %v", err)
		return
	}

	err = uc.CreateUser(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while creating user: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	defer db.Close()

	repo := repository.NewUserSQLRepository(db)
	uc := usecases.NewUpdateUserUseCase(repo)

	var input usecases.UpdateUserInput

	id := chi.URLParam(r, "id")

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while deconding request body: %v", err)
		return
	}
	input.ID = id

	output, err := uc.UpdateUser(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while updateing user: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	defer db.Close()

	repo := repository.NewUserSQLRepository(db)
	uc := usecases.NewUpdatePasswordUseCase(repo)

	var input usecases.UpdatePasswordInput

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while decoding body: %v", err)
		return
	}

	id := chi.URLParam(r, "id")
	input.ID = id

	output, err := uc.UpdatePassword(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error updating password: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}

	repo := repository.NewUserSQLRepository(db)
	uc := usecases.NewFindAllUsersUseCase(repo)

	output, err := uc.FindAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error getting all users: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}

	repo := repository.NewUserSQLRepository(db)
	uc := usecases.NewFindUserByIDUseCase(repo)

	var input usecases.FindUserByIDInput
	id := chi.URLParam(r, "id")
	input.ID = id

	output, err := uc.FindUserByID(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error get user data: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
