package repositories

type Repository interface {
	Create(entity interface{}) error
	FindByID(id interface{}, entity interface{}) error
	Update(entity interface{}) error
	Delete(entity interface{}) error
}
