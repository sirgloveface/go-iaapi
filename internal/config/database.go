package config

import (
	"fmt"
	"log"
	"os"

	"github.com/sirgloveface/go-iaapi/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() *gorm.DB {
	dsn := os.Getenv("POSTGRES_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=goapi port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	// Auto-migrate models
	if err := db.AutoMigrate(&model.Task{}); err != nil {
		log.Fatalf("failed to auto-migrate: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL")
	return db
}
