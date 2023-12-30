package repositories

import (
	"go-connect/pkg/domain/entities"
	base "go-connect/pkg/domain/repositories/base"

	"gorm.io/gorm"
)

type UserRepository struct {
	*base.BaseRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{base.NewBaseRepository(db)}
}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(user *entities.User) error {
	return r.Create(user)
}
