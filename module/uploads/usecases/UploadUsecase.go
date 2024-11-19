package usecases

import (
	"beer/module/uploads/models"
	"beer/module/uploads/repositories"

	"fmt"
)

type UploadUsecase interface {
	UploadDataProcessing(in *models.CreateUploadData) (int64, error)
	GetUploadData(uploadId int) (*models.UploadData, error)

}

type UploadUsecaseImpl struct {
	uploadRepository repositories.UploadRepository
}

func NewUploadUsecaseImpl(uploadRepository repositories.UploadRepository) UploadUsecase {
	return &UploadUsecaseImpl{
		uploadRepository: uploadRepository,
	}
}

func (u *UploadUsecaseImpl) UploadDataProcessing(in *models.CreateUploadData) (int64, error) {
	if in == nil {
		return 0, fmt.Errorf("input data is nil")
	}

	id, err := u.uploadRepository.CreateUploadData(in) // เปลี่ยนให้รับค่า ID
	if err != nil {
		return 0, fmt.Errorf("failed to upload data: %w", err)
	}

	return id, nil // คืนค่า ID ที่ถูกสร้าง
}


func (u *UploadUsecaseImpl) GetUploadData(uploadId int) (*models.UploadData, error) {
	result, err := u.uploadRepository.GetUploadData(uploadId) // เปลี่ยนให้รับค่า ID
	if err != nil {
		return nil, fmt.Errorf("failed to upload data: %w", err)
	}

	return result, nil // คืนค่า ID ที่ถูกสร้าง
}

