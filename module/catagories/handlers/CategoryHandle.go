package handlers

import (
	"beer/module/catagories/models"
	"beer/module/catagories/usecases"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type CategoryHandler interface {
	GetDataCategory(categoryId int) (*models.GetCategory ,error)
}

type categoryHttpHandler struct {
	categoryUsecase usecases.CategoryUsecase
}

func NewCategoryHttpHandler(categoryUsecase usecases.CategoryUsecase) CategoryHandler {
	return &categoryHttpHandler{
		categoryUsecase: categoryUsecase,
	}
}

// GetCategory ดึงข้อมูล category ตาม ID
func (h *categoryHttpHandler) GetDataCategory(categoryId int) (*models.GetCategory, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ดึงข้อมูล category จาก Usecase
	dataCategory, err := h.categoryUsecase.CategoryDataProcess(categoryId)
	if err != nil || dataCategory == nil {
		return nil, err // คืนค่าผลลัพธ์เป็น nil และ error
	}

	return dataCategory, nil
}
