package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gui-laranjeira/go-cleanarch/internal/infrastructure/db"
	"github.com/gui-laranjeira/go-cleanarch/internal/infrastructure/repository"
	usecases "github.com/gui-laranjeira/go-cleanarch/internal/usecases/costumer"
)

func CreateCostumerHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	repo := repository.NewCostumerSQLRepository(db)
	uc := usecases.NewCreateCostumerUseCase(repo)

	var input usecases.CreateCostumerInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while decoding body: %v", err)
		return
	}

	err = uc.CreateCostumer(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error creating costumer: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateCostumerHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	repo := repository.NewCostumerSQLRepository(db)
	uc := usecases.NewUpdateCostumerUseCase(repo)

	id := chi.URLParam(r, "id")

	oldData, err := repo.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Printf("Costumer not found: %v", err)
		return
	}

	var input usecases.UpdateCostumerInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while decoding body: %v", err)
		return
	}

	//#TODO Refactor
	input.ID = id
	if input.Email == "" {
		input.Email = oldData.Email
	}
	if input.Phone == "" {
		input.Phone = oldData.Phone
	}
	if input.Address == "" {
		input.Address = oldData.Address
	}
	if input.Document == "" {
		input.Document = oldData.Document
	}
	if input.FirstName == "" {
		input.FirstName = oldData.FirstName
	}
	if input.LastName == "" {
		input.LastName = oldData.LastName
	}

	output, err := uc.UpdateCostumer(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while updating costumer: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func GetAllCostumersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	repo := repository.NewCostumerSQLRepository(db)
	uc := usecases.NewFindAllCostumersUseCase(repo)

	output, err := uc.FindAllCostumers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while getting costumers: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func GetCostumerByIDHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	repo := repository.NewCostumerSQLRepository(db)
	uc := usecases.NewFindCostumerByIDUseCase(repo)

	id := chi.URLParam(r, "id")

	var input usecases.FindCostumerByIDInput
	input.ID = id

	output, err := uc.FindCostumerByID(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while finding costumer: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
