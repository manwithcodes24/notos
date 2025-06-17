package controllers

import (
	"context"
	"time"

	"net/http"
	"notos/database"
	"notos/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// collection variables
var (
	NotesCollection *mongo.Collection = database.OpenCollection(database.Client, "notes")
	SubjectsCollection *mongo.Collection = database.OpenCollection(database.Client, "subjects")
)

func GetNotes(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var notes []models.Note
	cursor, err := NotesCollection.Find(ctx, bson.M{})

	if err != nil { 
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while listing notes"})
		return
	}	
	err = cursor.All(ctx,&notes)
	if err != nil { 
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while listing notes"})
		return
	}
	c.JSON(http.StatusOK, notes)
}

func GetNoteById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Note by ID working perfectly",
	})
}

func UploadNotes(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Upload Notes working perfectly",
	})
}

func UpdateNotes(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update Notes working perfectly",
	})
}

func DeleteNotes(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Notes working perfectly",
	})
}

func GetNotesBySubjectId(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Notes by Subject ID working perfectly",
	})
}