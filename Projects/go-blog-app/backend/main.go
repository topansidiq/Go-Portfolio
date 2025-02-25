package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	DB.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}
	port := os.Getenv("PORT")
}
