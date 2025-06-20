package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" unique:"true"`
	Name     string `json:"name" bson:"name" validate:"required,min=3,max=32"`
	Username string `json:"username" bson:"username" unique:"true" validate:"required,min=3"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Role     int `json:"role" bson:"role" default:"0"`
	Password string `json:"password" bson:"password" validate:"required,min=6"`
	TotalNotes int `json:"totalNotes" bson:"totalNotes" default:"0"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	
}
