package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gui-laranjeira/go-cleanarch/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	cfg := configs.GetDB()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = db.Ping()
	return db, err
}
