package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	NotifyGo struct {
		Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		Detail    string             `json:"detail,omitempty" bson:"detail,omitempty"`
		Title     string             `json:"title,omitempty" bson:"title,omitempty"`
		CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	}

	CreateNotifyGo struct {
		Detail    string             `json:"detail,omitempty"`
		Title     string             `json:"title,omitempty"`
		CreatedAt time.Time          `json:"created_at"`
	}
)
