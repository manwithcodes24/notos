package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Upvote struct {
	ID primitive.ObjectID `bson:"_id"`
	NoteID primitive.ObjectID `json:"noteId" bson:"noteId" validate:"required"`
	UserID primitive.ObjectID `json:"userId" bson:"userId" validate:"required"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
