package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

type mySQLDatabase struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *mySQLDatabase
)

func ConMySQLDatabase() *mySQLDatabase {
	// โหลดไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// เรียกใช้งานเพียงครั้งเดียว
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_PASS"),
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_NAME"),
		)

		var err error
		dbInstance, err = connectDatabase(dsn)
		if err != nil {
			panic("failed to connect database")
		}
	})
	return dbInstance
}

func connectDatabase(dsn string) (*mySQLDatabase, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &mySQLDatabase{Db: db}, nil
}

// GetDb returns the GORM DB instance
func (m *mySQLDatabase) GetDb() *gorm.DB {
	return m.Db
}
