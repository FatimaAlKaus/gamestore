package main

import (
	"Lab1/configs"
	"Lab1/db"
	"Lab1/pkg/handler"
	"Lab1/pkg/repository"
	"fmt"
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
		// os.Exit(1)
	}
	defer db.Close()

	repository := repository.NewRepository(db)
	handlers := handler.NewHandler(repository)

	server := handlers.InitRoutes()

	server.Run(config.GetAddress())
}
