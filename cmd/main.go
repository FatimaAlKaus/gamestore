package main

import (
	"Lab1/configs"
	"Lab1/db"
	"Lab1/pkg/handler"
	"Lab1/pkg/repository"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	config, err := configs.New()
	if err != nil {
		fmt.Printf("fatal error config file: %s \n", err.Error())
		os.Exit(1)
	}
	db, err := db.NewPostgresDb(config)
	if err != nil {
		fmt.Printf("Faield to init db: %s \n", err.Error())
		os.Exit(1)
	}
	defer db.Close()
	m, err := migrate.New(
		"file://db/migrations",
		config.GetConnectionString())
	if err != nil {
		fmt.Printf("Failed to read migrations: %s", err.Error())
	}
	if err := m.Up(); err != nil {
		fmt.Printf("Failed to apply migration: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	handlers := handler.NewHandler(repository)

	server := handlers.InitRoutes()

	server.Run(config.GetAddress())
	fmt.Printf("Server started")
}
