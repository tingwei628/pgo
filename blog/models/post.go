package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	gorm.Model
}

// MigrateDB migrates the schema
func MigrateDB(db *gorm.DB) error {
	if !db.Migrator().HasTable(&Post{}) {
		if err := db.AutoMigrate(&Post{}); err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
	}
	return nil
}
