package db

import (
	"Lab1/configs"
	"database/sql"
	_ "github.com/lib/pq"
)

func NewPostgresDb(connection *configs.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection.GetConnectionString())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
