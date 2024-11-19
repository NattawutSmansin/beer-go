package usecases

import (
	"beer/module/catagories/models"
	"beer/module/catagories/repositories"
	"fmt"
)

type CategoryUsecase interface {
	CategoryDataProcess(beerId int) (*models.GetCategory, error)
}

type CategoryUsecaseImpl struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryUsecaseImpl(categoryRepository repositories.CategoryRepository) CategoryUsecase {
	return &CategoryUsecaseImpl{
		categoryRepository: categoryRepository,
	}
}

func (categoryUsecaseImpl *CategoryUsecaseImpl) CategoryDataProcess(categoryId int) (*models.GetCategory, error) {
	category, err := categoryUsecaseImpl.categoryRepository.Data(categoryId)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data: %w", err)
	}

	return category, nil
}
