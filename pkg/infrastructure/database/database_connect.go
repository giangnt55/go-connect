package database

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func SqlServerConnect() (*gorm.DB, error) {
	dsn := os.Getenv("CONNECTION_STRING")

	// Open a connection to the SQL Server database
	db, err := gorm.Open("mssql", dsn)
	if err != nil {
		return nil, err
	}

	// Disable default transaction behavior
	db.LogMode(true)
	db.SingularTable(true)

	return db, nil
}
