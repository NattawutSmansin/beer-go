package repositories

import (
	"beer/database"
	"beer/module/catagories/models"

	"github.com/labstack/gommon/log"
)

type CategoryDatabaseRepository struct {
	db database.Database
}

type CategoryRepository interface {
	Data(CategoryId int) (*models.GetCategory, error)
}

func NewCategoryRepository(db database.Database) CategoryRepository {
	return &CategoryDatabaseRepository{db: db}
}

// Data retrieves a single Beer by ID.
func (r *CategoryDatabaseRepository) Data(categoryId int) (*models.GetCategory, error) {
	var category models.GetCategory

	// Query category data by ID with WHERE condition
	result := r.db.GetDb().Where("id = ?", categoryId).Where("is_active = ?", true).First(&category)

	if result.Error != nil {
		log.Errorf("fetch beer data: %v", result.Error)
		return nil, result.Error
	}

	log.Debugf("retrieved beer data: %v", category)
	return &category, nil
}
