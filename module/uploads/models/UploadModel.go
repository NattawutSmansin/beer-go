package models

import (
	"time"
)

type (
	InsertUploadDto struct {
		Id           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
		OriginalName string    `gorm:"type:varchar(512)" json:"original_name"`
		FileName     string    `gorm:"type:varchar(512)" json:"file_name"`
		FilePath     string    `gorm:"type:varchar(512)" json:"file_path"`
		FileSize     float32   `gorm:"type:float(10,2)" json:"file_size"`
		FileType     string    `gorm:"type:varchar(512)" json:"file_type"`
		CreatedAt    time.Time `gorm:"type:timestamp" json:"created_at"`
		CreatedBy    int32     `json:"created_by"`
		UpdatedAt    time.Time `gorm:"type:datetime" json:"updated_at"`
		UpdatedBy    uint32    `json:"updated_by"`
		IsActive     bool      `gorm:"default:true" json:"is_active"`
	}

	Upload struct {
		Id           int64     `json:"id"`
		OriginalName string    `json:"original_name"`
		FileName     string    `json:"file_name"`
		FilePath     string    `json:"file_path"`
		FileSize     float32   `json:"file_size"`
		FileType     string    `json:"file_type"`
		CreatedAt    time.Time `json:"created_at"`
		CreatedBy    int32     `json:"created_by"`
		UpdatedAt    time.Time `json:"updated_at"`
		UpdatedBy    int32     `json:"updated_by"`
		IsActive     bool      `json:"is_active"`
	}

	CreateUploadData struct {
		Id           int64   `json:"id"`
		OriginalName string  `json:"original_name"`
		FileName     string  `json:"file_name"`
		FilePath     string  `json:"file_path"`
		FileSize     float32 `json:"file_size"`
		FileType     string  `json:"file_type"`
		IsActive     bool    `json:"is_active"`
	}

	UploadData struct {
		Id           int64   `json:"id"`
		OriginalName string  `json:"original_name"`
		FileName     string  `json:"file_name"`
		FilePath     string  `json:"file_path"`
		FileSize     float32 `json:"file_size"`
		FileType     string  `json:"file_type"`
	}
)

func (UploadData) TableName() string {
	return "uploads"
}

func (CreateUploadData) TableName() string {
	return "uploads"
}

func (Upload) TableName() string {
	return "uploads"
}

func (InsertUploadDto) TableName() string {
	return "uploads"
}
