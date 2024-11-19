package repositories

import (
	"beer/database"
	"beer/module/beers/models"
	"beer/module/beers/validations"
	"fmt"

	"github.com/labstack/gommon/log"
)

type BeerDatabaseRepository struct {
	db database.Database
}

type BeerRepository interface {
	List(name string ,page,limit int) ([]*models.Beer, int64, error)
	Data(beerId uint32) (*models.Beer, error)
	Store(in *models.CreateBeer) error
	Update(in *models.UpdateBeer) error
	Delete(beerId uint32) error
}

func NewBeerRepository(db database.Database) BeerRepository {
	return &BeerDatabaseRepository{db: db}
}

func (r *BeerDatabaseRepository) Store(in *models.CreateBeer) error {
	// Validate the input
	if err := validations.ValidateBeerData(in); err != nil {
		log.Errorf("validation failed: %v", err)
		return err
	}

	data := &models.CreateBeer{
		Name:         in.Name,
		Description:  in.Description,
		CategoryId:   in.CategoryId,
		ImageFileIds: in.ImageFileIds,
		IsActive:     true,
	}

	result := r.db.GetDb().Create(data)
	if result.Error != nil {
		log.Errorf("create upload: %v", result.Error)
		return result.Error
	}

	log.Debugf("create upload : %v", result.RowsAffected)
	return nil
}

// List retrieves all active Beers from the database with pagination.
func (r *BeerDatabaseRepository) List(name string, page, limit int) ([]*models.Beer, int64, error) {
	var beers []*models.Beer
	var total int64

	// Validate page and limit to ensure they are positive numbers
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10 // Default limit
	}

	// Construct the query with optional name filter
	query := r.db.GetDb().Where("is_active = ?", true)
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	// Count the total number of records that match the query
	if err := query.Model(&models.Beer{}).Count(&total).Error; err != nil {
		log.Errorf("count beers: %v", err)
		return nil, 0, err
	}

	// Apply pagination (limit and offset)
	offset := (page - 1) * limit
	result := query.Limit(limit).Offset(offset).Find(&beers)

	// Handle potential errors
	if result.Error != nil {
		log.Errorf("list beers: %v", result.Error)
		return nil, 0, result.Error
	}

	log.Debugf("listed beers: %v items on page %v", len(beers), page)
	return beers, total, nil
}


func (r *BeerDatabaseRepository) Data(beerId uint32) (*models.Beer, error) {
	var beer models.Beer
	result := r.db.GetDb().First(&beer, beerId)

	if result.Error != nil {
		log.Errorf("fetch beer data: %v", result.Error)
		return nil, result.Error
	}

	log.Debugf("retrieved beer data: %v", beer)
	return &beer, nil
}

func (r *BeerDatabaseRepository) Update(in *models.UpdateBeer) error {
	// ตรวจสอบว่าเบียร์ที่ต้องการอัปเดตมีอยู่ในฐานข้อมูลหรือไม่
	var existingBeer models.UpdateBeer
	if err := r.db.GetDb().First(&existingBeer, in.Id).Error; err != nil {
		// ถ้าไม่พบข้อมูลในฐานข้อมูล ให้แสดง error
		if err.Error() == "record not found" {
			return fmt.Errorf("beer with ID %d not found", in.Id)
		}
		return err
	}

	// อัปเดตข้อมูลเบียร์
	// ใช้ Update เพื่ออัปเดตข้อมูลที่ตรงกับ id เท่านั้น
	result := r.db.GetDb().Model(&existingBeer).Updates(in)
	if result.Error != nil {
		log.Errorf("update beer: %v", result.Error)
		return result.Error
	}

	log.Debugf("updated beer: %v", in)
	return nil
}

func (r *BeerDatabaseRepository) Delete(beerId uint32) error {
	result := r.db.GetDb().Model(&models.Beer{}).Where("id = ?", beerId).Update("is_active", false)

	if result.Error != nil {
		log.Errorf("delete (deactivate) beer: %v", result.Error)
		return result.Error
	}

	resultDelete := r.db.GetDb().Model(&models.Beer{}).Where("id = ?", beerId).Delete(&models.Beer{})
	if resultDelete.Error != nil {
		log.Errorf("delete (soft delete) beer: %v", resultDelete.Error)
		return resultDelete.Error
	}

	log.Debugf("deactivated beer with ID: %v", beerId)
	return nil
}
