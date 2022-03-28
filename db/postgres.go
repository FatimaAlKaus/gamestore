package db

import (
	"Lab1/configs"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

func NewPostgresDb(connection *configs.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection.GetConnectionString())
	time.Sleep(3 * time.Second)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
