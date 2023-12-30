package entities

import "go-connect/pkg/domain/base"

// User represents the structure of the 'users' table
type User struct {
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	base.SoftDeleteModel
}

func (User) TableName() string {
	return "cardpie.dbo.users"
}
