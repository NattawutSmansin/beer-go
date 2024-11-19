package database

import (
	"gorm.io/gorm"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database interface {
	GetDb() *gorm.DB
}

type DatabaseMongo interface {
	GetDb() *mongo.Client
}