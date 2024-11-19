package repositories

import (
	"beer/database"
	"beer/module/notifies/models"
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/labstack/gommon/log"
)

type NotifyDatabaseRepository struct {
	db database.DatabaseMongo
}

type NotifyRepository interface {
	CreateNotifyData(in *models.CreateNotifyGo) error
}

func NewNotifyRepository(db database.DatabaseMongo) NotifyRepository {
	return &NotifyDatabaseRepository{db: db}
}

// Adjust function implementation for correct database access
func (r *NotifyDatabaseRepository) CreateNotifyData(in *models.CreateNotifyGo) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	data := &models.NotifyGo{
		Title:     in.Title,
		Detail:    in.Detail,
		CreatedAt: in.CreatedAt,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	dbName := os.Getenv("MONGO_DB_DATABASE")
	collectionName := os.Getenv("MONGO_DB_COLLECTION")

	collection := r.db.GetDb().Database(dbName).Collection(collectionName)
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Errorf("create notify: %v", err)
		return err
	}

	log.Debugf("create notify succeeded with ID: %v", result.InsertedID)
	return nil
}
