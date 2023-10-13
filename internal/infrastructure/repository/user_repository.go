package repository

import (
	"database/sql"

	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	repository "github.com/gui-laranjeira/go-cleanarch/internal/infrastructure/repository/errors"
)

type UserSQLRepository struct {
	db *sql.DB
}

func NewUserSQLRepository(db *sql.DB) entity.IUserRepository {
	return &UserSQLRepository{
		db: db,
	}
}

func (r *UserSQLRepository) Create(user *entity.User) error {
	sqlStatement := `INSERT INTO users (id_user, email, phone, password, first_name, last_name, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	rows, err := r.db.Exec(sqlStatement, user.ID, user.Email, user.Phone, user.Password, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt)
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

func (r *UserSQLRepository) Update(newUser *entity.User) (int64, error) {
	SqlStatement := `UPDATE users SET email = $1, phone = $2, first_name = $3, last_name = $4, updated_at = $5 WHERE id_user=$6`
	rows, err := r.db.Exec(SqlStatement, newUser.Email, newUser.Phone, newUser.FirstName, newUser.LastName, newUser.UpdatedAt, newUser.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

// #TODO implement Login
func (r *UserSQLRepository) Login(email string, password string) (*entity.User, error) {
	return nil, nil
}

func (r *UserSQLRepository) ChangePassword(id string, newPassword string) (int64, error) {
	sqlStatement := `UPDATE users SET password = $1 WHERE id_user = $2`
	rows, err := r.db.Exec(sqlStatement, newPassword, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (r *UserSQLRepository) GetAllUsers() ([]*entity.User, error) {
	sqlStatement := `SELECT id_user, email, phone, first_name, last_name, created_at, updated_at FROM users`
	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.Email, &user.Phone, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserSQLRepository) GetUserByID(id string) (*entity.User, error) {
	sql := `SELECT id_user, email, phone, first_name, last_name, created_at, updated_at FROM users WHERE id_user = $1`
	rows, err := r.db.Query(sql, id)
	if err != nil {
		return nil, err
	}

	var user entity.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.Phone, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}
