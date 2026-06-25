package main

import (
	"log"
	"QuotesService/internal/app"

	"github.com/joho/godotenv"
)

// @title           Quotes Service API
// @version         1.0
// @description     A minimal service that serves random weather, nature, and science facts.
// @BasePath        /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	a := app.New()
	a.Run()
}
