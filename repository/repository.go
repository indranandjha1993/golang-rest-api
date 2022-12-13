package repository

type Repository interface {
	Insert(model interface{}) (interface{}, error)
	// Update(string, *models.User) error
	Delete(string) error
	// GetByID(string) (*models.User, error)
	GetAll() ([]interface{}, error)
}
