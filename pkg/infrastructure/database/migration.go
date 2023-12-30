package database

import (
	"go-connect/pkg/domain/entities"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	// Auto-migrate your database schema
	if err := db.AutoMigrate(
		&entities.User{},
	); err != nil {
		return err
	}

	return nil
}
