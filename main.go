package main

import (
	"books/infra/database"
	"books/server"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	fmt.Println(os.Getenv("DB_HOST"))
	database.InitDatabase()

	server := server.NewServer()

	if err != nil {
		log.Fatal(".env file not found")
	}

	server.Run()
}
