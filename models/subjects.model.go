package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subject struct {
	ID primitive.ObjectID `bson:"_id"`
	Name string `json:"name" bson:"name" validate:"required,min=3,max=32"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
