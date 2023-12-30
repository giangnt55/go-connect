package repositories

import "gorm.io/gorm"

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func (r *BaseRepository) Create(entity interface{}) error {
	return r.db.Create(entity).Error
}

func (r *BaseRepository) FindByID(id interface{}, entity interface{}) error {
	return r.db.First(entity, id).Error
}

func (r *BaseRepository) Update(entity interface{}) error {
	return r.db.Save(entity).Error
}

func (r *BaseRepository) Delete(entity interface{}) error {
	return r.db.Delete(entity).Error
}
