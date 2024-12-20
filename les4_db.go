package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := os.Getenv("POSTGRES_DB")
	fmt.Printf("DB name - %s \n", db)

}
