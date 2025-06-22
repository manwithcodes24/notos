package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"

	"log"
	"notos/internal/database"
	"notos/internal/models"

	"gopkg.in/bluesuncorp/validator.v5"
)

var (
	UsersCollection    *mongo.Collection = database.OpenCollection(database.Client, "users")
	NotesCollection    *mongo.Collection = database.OpenCollection(database.Client, "notes")
	RefreshTokenCollection    *mongo.Collection = database.OpenCollection(database.Client, "refresh_tokens")
	SubjectsCollection *mongo.Collection = database.OpenCollection(database.Client, "subjects")
	validate                             = validator.New("validate", validator.BakedInValidators)
)

func init() {
	err := models.CreateUserIndexes(UsersCollection)
	if err != nil {
		log.Println("Could not create indexes:", err)
	}
	return
}
