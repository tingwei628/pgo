package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tingwei628/pgo/blog/models"
	"github.com/tingwei628/pgo/blog/routes"
	"github.com/tingwei628/pgo/blog/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if os.Getenv("GIN_MODE") == "" {
		if err := godotenv.Load(filepath.Join(utils.Basepath, ".env")); err != nil {
			log.Fatalf("Error loading .env err  %v", err)
		}
	}
	// DB connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	// Migrate
	if err := models.MigrateDB(db); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Setup routes
	r := routes.SetupRouter(db)
	r.Run(":" + os.Getenv("PORT"))
}
