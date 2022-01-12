package main

import (
	"books/infra/database"
	"books/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	database.InitDatabase()

	server := server.NewServer()

	if err != nil {
		log.Fatal(".env file not found")
	}

	server.Run()
}
