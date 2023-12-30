package base

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SoftDeleteModel includes soft delete fields in the model
type SoftDeleteModel struct {
	ID        uuid.UUID `gorm:"type:nvarchar(50);primaryKey;default:NEWID()" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
