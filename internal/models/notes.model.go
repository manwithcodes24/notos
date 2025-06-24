package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID primitive.ObjectID `bson:"_id"`
	Title *string `json:"title" bson:"title" validate:"required,min=3,max=32"`
	PdfLink *string `json:"pdfLink" bson:"pdfLink" validate:"required,min=3,max=32"`
	SubjectID primitive.ObjectID `json:"subjectId" bson:"subjectId"`
	UserID primitive.ObjectID `json:"userId" bson:"userId"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
