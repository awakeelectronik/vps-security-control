package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"vps-security-control/backend/internal/config"
	"vps-security-control/backend/internal/database"
)

func main() {
	godotenv.Load()

	cfg := config.LoadConfig()

	db, err := database.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	fmt.Println("\u2705 Migrations completed successfully")
}
