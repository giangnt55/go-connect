package database

import (
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func SqlServerConnect() (*gorm.DB, error) {
	dsn := os.Getenv("CONNECTION_STRING")

	// Open a connection to the SQL Server database
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	//Migration
	MigrateDB(db.Debug())

	return db, nil
}
