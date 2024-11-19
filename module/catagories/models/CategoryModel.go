package models

import (
	"time"
)

type (
	InsertCategoryDto struct {
		Id        int32      `gorm:"primaryKey;autoIncrement" json:"id"`
		Name      string     `gorm:"type:varchar(256)" json:"name"`
		CreatedAt time.Time  `gorm:"type:timestamp" json:"created_at"`
		CreatedBy *int32     `json:"created_by"`
		UpdatedAt *time.Time `gorm:"type:datetime" json:"updated_at"`
		UpdatedBy *int32     `json:"updated_by"`
		IsActive  bool       `gorm:"default:true" json:"is_active"`
		DeletedAt *time.Time `gorm:"type:datetime" json:"deleted_at"`
	}

	Category struct {
		Id        int32      `json:"id"`
		Name      string     `json:"name"`
		CreatedAt time.Time  `json:"created_at"`
		CreatedBy *int32     `json:"created_by"`
		UpdatedAt *time.Time `json:"updated_at"`
		UpdatedBy *int32     `json:"updated_by"`
		IsActive  bool       `json:"is_active"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	GetCategory struct {
		Id        int32      `json:"id"`
		Name      string     `json:"name"`
	}
)


func (Category) TableName() string {
	return "categories"
}

func (GetCategory) TableName() string {
	return "categories"
}