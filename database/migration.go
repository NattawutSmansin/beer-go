package database

import (
	beerModel "beer/module/beers/models"
	categoryModel "beer/module/catagories/models"
	uploadModel "beer/module/uploads/models"
	"fmt"
)

func Migrate() {
	db := ConMySQLDatabase()

	if db != nil {
		fmt.Println("Database connection established successfully!")
	}

	dropTables(db)

	migrateCategory(db)
	migrateBeer(db)
	migrateUpload(db)

	fmt.Println("Migrate Success!!")
}

func dropTables(db Database) {
	// ตรวจสอบและลบตารางถ้ามีอยู่
	if db.GetDb().Migrator().HasTable(&categoryModel.Category{}) {
		db.GetDb().Migrator().DropTable(&categoryModel.Category{})
		fmt.Println("Dropped table category.")
	}
	if db.GetDb().Migrator().HasTable(&beerModel.Beer{}) {
		db.GetDb().Migrator().DropTable(&beerModel.Beer{})
		fmt.Println("Dropped table beer.")
	}
	if db.GetDb().Migrator().HasTable(&uploadModel.Upload{}) {
		db.GetDb().Migrator().DropTable(&uploadModel.Upload{})
		fmt.Println("Dropped table upload.")
	}
}

func migrateCategory(db Database) {
	db.GetDb().Migrator().CreateTable(&categoryModel.Category{})
	db.GetDb().CreateInBatches([]categoryModel.Category{
		{Name: "เอล", IsActive: true},
		{Name: "ลาเกอร์", IsActive: true},
		{Name: "เบียร์ดำ", IsActive: true},
		{Name: "เบียร์สด", IsActive: true},
		{Name: "ไลท์เบียร์", IsActive: true},
	}, 10)
}

func migrateBeer(db Database) {
	db.GetDb().Migrator().CreateTable(&beerModel.Beer{})
}

func migrateUpload(db Database) {
	db.GetDb().Migrator().CreateTable(&uploadModel.Upload{})
}
