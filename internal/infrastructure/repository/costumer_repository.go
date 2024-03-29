package repository

import (
	"database/sql"

	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	repository "github.com/gui-laranjeira/go-cleanarch/internal/infrastructure/repository/errors"
)

type CostumerSQLRepository struct {
	db *sql.DB
}

func NewCostumerSQLRepository(db *sql.DB) entity.ICostumerRepository {
	return &CostumerSQLRepository{
		db: db,
	}
}

func (r *CostumerSQLRepository) Create(Costumer *entity.Costumer) error {
	sqlStatement := `INSERT INTO costumers (id_costumer, email, phone, address, document, first_name, last_name, created_at, updated_at) 
					      VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	rows, err := r.db.Exec(sqlStatement, Costumer.ID.String(), Costumer.Email, Costumer.Phone, Costumer.Address, Costumer.Document, Costumer.FirstName, Costumer.LastName, Costumer.CreatedAt, Costumer.UpdatedAt)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return repository.NoRowsAffectedError()
	}

	if rowsAffected > 1 {
		return repository.MultipleRowsAffectedError()
	}

	return nil
}

func (r *CostumerSQLRepository) Update(newCostumer *entity.Costumer) (int64, error) {

	SqlStatement := `UPDATE Costumers 
						SET email = $1, phone = $2, first_name = $3, last_name = $4, updated_at = $5, 
					 		address = $6, document = $7 
					  WHERE id_Costumer=$8`

	rows, err := r.db.Exec(SqlStatement, newCostumer.Email, newCostumer.Phone,
		newCostumer.FirstName, newCostumer.LastName, newCostumer.UpdatedAt,
		newCostumer.Address, newCostumer.Document, newCostumer.ID)

	if err != nil {
		return 0, err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (r *CostumerSQLRepository) FindAll() ([]*entity.Costumer, error) {
	query := `SELECT * FROM costumers`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var costumers []*entity.Costumer
	for rows.Next() {
		var costumer entity.Costumer
		err = rows.Scan(&costumer.ID, &costumer.Email, &costumer.Phone, &costumer.Address, &costumer.Document,
			&costumer.FirstName, &costumer.LastName, &costumer.CreatedAt, &costumer.UpdatedAt, &costumer.CurrentBookID)

		if err != nil {
			return nil, err
		}
		costumers = append(costumers, &costumer)
	}

	return costumers, nil
}

func (r *CostumerSQLRepository) FindByID(id string) (*entity.Costumer, error) {
	query := `SELECT * FROM costumers WHERE id_costumer = $1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var costumer entity.Costumer
	for rows.Next() {
		err = rows.Scan(&costumer.ID, &costumer.Email, &costumer.Phone, &costumer.Address, &costumer.Document,
			&costumer.FirstName, &costumer.LastName, &costumer.CreatedAt, &costumer.UpdatedAt, &costumer.CurrentBookID)

		if err != nil {
			return nil, err
		}
	}
	return &costumer, nil
}
