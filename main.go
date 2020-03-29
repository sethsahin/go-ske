package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"skeleton/controllers"
)

var server = controllers.Server{}

func init() {
	// load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print(".env file not found")
	}
}

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting the env, %v", err)
	} else {
		fmt.Println(".env values loading")
	}

	server.Connect(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_DRIVER"), os.Getenv("DB_PORT"))

	server.Run(":8080")
}
