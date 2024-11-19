package repositories

import (
	"beer/database"
	"beer/module/uploads/models"

	"github.com/labstack/gommon/log"
)

type UploadDatabaseRepository struct {
	db database.Database
}

type UploadRepository interface {
	CreateUploadData(in *models.CreateUploadData) (int64, error) 
	GetUploadData(uploadId int) (*models.UploadData, error) 

}

func NewUploadRepository(db database.Database) UploadRepository {
	return &UploadDatabaseRepository{db: db}
}

// ปรับปรุงการประกาศฟังก์ชัน
func (r *UploadDatabaseRepository) CreateUploadData(in *models.CreateUploadData) (int64, error) {
	data := &models.Upload{
		OriginalName: in.OriginalName,
		FileName:     in.FileName,
		FilePath:     in.FilePath,
		FileSize:     in.FileSize,
		FileType:     in.FileType,
		IsActive:     true,
	}

	result := r.db.GetDb().Create(data)

	if result.Error != nil {
		log.Errorf("create upload: %v", result.Error)
		return 0, result.Error // คืนค่า 0 และ error
	}

	log.Debugf("create upload : %v", result.RowsAffected)
	return data.Id, nil
}

func (r *UploadDatabaseRepository) GetUploadData(uploadId int) (*models.UploadData, error) {
	var upload models.UploadData

	// Query category data by ID with WHERE condition
	result := r.db.GetDb().Where("id = ?", uploadId).Where("is_active = ?", true).First(&upload)

	if result.Error != nil {
		log.Errorf("fetch beer data: %v", result.Error)
		return nil, result.Error
	}

	log.Debugf("retrieved beer data: %v", upload)
	return &upload, nil
}
