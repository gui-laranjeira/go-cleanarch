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
		log.Printf("Error while openning connection: %v", err)
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
		log.Printf("Error while openning connection: %v", err)
		return
	}
	repo := repository.NewCostumerSQLRepository(db)
	uc := usecases.NewUpdateCostumerUseCase(repo)

	var input usecases.UpdateCostumerInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while decoding body: %v", err)
		return
	}

	input.ID = chi.URLParam(r, "id")

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

// TODO Implement BorrowBook and ReturnBook Handlers
func BorrowBookHandler(w http.ResponseWriter, r *http.Request) {}

func ReturnBookHandler(w http.ResponseWriter, r *http.Request) {}