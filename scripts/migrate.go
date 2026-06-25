package main

import (
	"QuotesService/internal/config"
	"QuotesService/internal/models"
	"QuotesService/internal/provider"
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	action := flag.String("action", "up", "Action to perform: up or down")
	flag.Parse()

	cfg := config.New()
	conn := provider.GetConnection(cfg)
	if conn == nil {
		panic("Failed to connect to the db")
	}

	switch *action {
	case "up":
		migrateUp(conn)
	case "down":
		migrateDown(conn)
	default:
		fmt.Printf("Unknown action: %s. Use 'up' or 'down'\n", *action)
	}
}

func migrateUp(db *gorm.DB) {
	fmt.Println("Migrating up...")
	if err := db.AutoMigrate(&models.Quote{}); err != nil {
		fmt.Printf("Error migrating up: %v\n", err)
	} else {
		fmt.Println("Migration up completed successfully.")
	}
}

func migrateDown(db *gorm.DB) {
	fmt.Println("Migrating down...")
	if err := db.Migrator().DropTable(&models.Quote{}); err != nil {
		fmt.Printf("Error migrating down: %v\n", err)
	} else {
		fmt.Println("Migration down completed successfully.")
	}
}
