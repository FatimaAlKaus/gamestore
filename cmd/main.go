package main

import (
	"Lab1/configs"
	"Lab1/db"
	"Lab1/pkg/handler"
	"Lab1/pkg/repository"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	config, err := configs.New()
	log := logrus.New()
	if err != nil {
		log.Error("fatal error config file: ", err.Error())
		os.Exit(1)
	}
	db, err := db.NewPostgresDb(config)
	if err != nil {
		log.Error("Failed to init db: ", err.Error())
		os.Exit(1)
	}
	defer db.Close()
	m, err := migrate.New(
		"file://db/migrations",
		config.GetConnectionString())
	if err != nil {
		log.Error("Failed to read migrations: ", err.Error())
	}
	if err := m.Up(); err != nil {
		log.Error("Failed to apply migration: ", err.Error())
	}

	repository := repository.NewRepository(db)
	handlers := handler.NewHandler(repository)

	server := handlers.InitRoutes()

	server.Run(config.GetAddress())
	log.Printf("Server started")
}
