package models

import (
	"beer/module/catagories/models"
	"time"
)

type (
	InsertBeerDto struct {
		Id           int32           `gorm:"primaryKey;autoIncrement" json:"id"`
		Name         string          `gorm:"type:varchar(256)" json:"name" validate:"required"`
		Description  string          `json:"description" validate:"required"`
		CategoryId   int32           `json:"category_id" validate:"required"`
		Category     models.Category `gorm:"foreignKey:CategoryId" json:"category"`
		ImageFileIds string          `gorm:"type:varchar(256)" json:"image_files" validate:"required"`
		CreatedAt    time.Time       `gorm:"type:timestamp" json:"created_at"`
		CreatedBy    *int32          `json:"created_by"`
		UpdatedAt    *time.Time      `gorm:"type:datetime" json:"updated_at"`
		UpdatedBy    *int32          `json:"updated_by"`
		IsActive     bool            `gorm:"default:true" json:"is_active"`
		DeletedAt    *time.Time      `gorm:"type:datetime" json:"deleted_at"`
	}

	Beer struct {
		Id           uint32          `gorm:"primaryKey;autoIncrement" json:"id"`
		Name         string          `gorm:"type:varchar(256)" json:"name" validate:"required"`
		Description  string          `json:"description" validate:"required"`
		CategoryId   uint32          `json:"category_id" validate:"required"`
		Category     models.Category `gorm:"foreignKey:CategoryId" json:"category"`
		ImageFileIds string          `gorm:"type:varchar(256)" json:"image_files" validate:"required"`
		CreatedAt    time.Time       `gorm:"type:timestamp" json:"created_at"`
		CreatedBy    *int32          `json:"created_by"`
		UpdatedAt    *time.Time      `gorm:"type:datetime" json:"updated_at"`
		UpdatedBy    *int32          `json:"updated_by"`
		IsActive     bool            `gorm:"default:true" json:"is_active"`
		DeletedAt    *time.Time      `gorm:"type:datetime" json:"deleted_at"`
	}

	CreateBeer struct {
		Name         string `gorm:"type:varchar(256)" json:"name" validate:"required"`
		Description  string `json:"description" validate:"required"`
		CategoryId   uint32 `json:"category_id" validate:"required"`
		ImageFileIds string `gorm:"type:varchar(256)" json:"image_files" validate:"required"`
		IsActive     bool   `gorm:"default:true" json:"is_active"`
	}

	UpdateBeer struct {
		Id           uint32 `gorm:"primaryKey;autoIncrement" json:"id"`
		Name         string `gorm:"type:varchar(256)" json:"name" validate:"required"`
		Description  string `json:"description" validate:"required"`
		CategoryId   uint32 `json:"category_id" validate:"required"`
		ImageFileIds string `gorm:"type:varchar(256)" json:"image_files" validate:"required"`
	}
)

func (CreateBeer) TableName() string {
	return "beers"
}

func (UpdateBeer) TableName() string {
	return "beers"
}

func (Beer) TableName() string {
	return "beers"
}
