package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RefreshToken struct {
	ID primitive.ObjectID `bson:"_id"`
	UserID primitive.ObjectID `bson:"user_id" json:"user_id" validate:"required"`
	Token string `bson:"token" json:"token" validate:"required"`
	ExpiresAt time.Time `bson:"expires_at" json:"expires_at" validate:"required"`
	Validated bool `bson:"validated" json:"validated" validate:"required"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}