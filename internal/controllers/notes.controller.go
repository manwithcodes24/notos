package controllers

import (
	"context"
	"log"
	"time"

	"net/http"
	"notos/internal/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// collection variables

func GetNotes(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var notes []models.Note
	cursor, err := NotesCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while listing notes"})
		return
	}
	err = cursor.All(ctx, &notes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while listing notes"})
		return
	}
	c.JSON(http.StatusOK, notes)
}

func GetNoteById(c *gin.Context) {
	noteId, err := primitive.ObjectIDFromHex(c.Param("noteId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while parsing noteId"})
		return
	}
	var Note models.Note
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err = NotesCollection.FindOne(ctx, bson.M{"_id": noteId}).Decode(&Note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while finding note"})
		return
	}
	c.JSON(200, Note)
}

func UploadNotes(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var Note models.Note
	var Subject models.Subject
	var user models.User
	if err := c.BindJSON(&Note); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while binding JSON"})
		return
	}
	valErr := validate.Struct(Note)
	if valErr != nil {
		log.Println(valErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while validating JSON"})
		return
	}
	if err := UsersCollection.FindOne(ctx, bson.M{"_id": Note.UserID}).Decode(&user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error Occured while finding user"})
		return
	}
	log.Println(Note.SubjectID)
	if err := SubjectsCollection.FindOne(ctx, bson.M{"_id": Note.SubjectID}).Decode(&Subject); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error Occured while finding subject"})
		return
	}
	Note.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	Note.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	Note.ID = primitive.NewObjectID()
	_, err := NotesCollection.InsertOne(ctx, Note)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while inserting note"})
		return
	}
	c.JSON(http.StatusOK, Note)

}

func UpdateNotes(c *gin.Context) {
	var Note models.Note
	noteId, err := primitive.ObjectIDFromHex(c.Param("noteId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while parsing noteId"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	if err := c.BindJSON(&Note); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while binding JSON"})
		return
	}
	valErr := validate.Struct(Note)
	if valErr != nil {
		log.Println(valErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while validating JSON"})
		return
	}
	Note.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	err = NotesCollection.FindOneAndUpdate(ctx, bson.M{"_id": noteId}, bson.M{"$set": Note}).Decode(&Note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while updating note"})
		return
	}
}

func DeleteNotes(c *gin.Context) {
	noteId, err := primitive.ObjectIDFromHex(c.Param("noteId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while parsing noteId"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	_, err = NotesCollection.DeleteOne(ctx, bson.M{"_id": noteId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while deleting note"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}

func GetNotesBySubjectId(c *gin.Context) {
	subjectId, err := primitive.ObjectIDFromHex(c.Param("subjectId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while parsing subjectId"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var subject models.Subject
	err = SubjectsCollection.FindOne(ctx, bson.M{"_id": subjectId}).Decode(&subject)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error Occured while finding subject"})
		return
	}
	var notes []models.Note
	cursor, err := NotesCollection.Find(ctx, bson.M{"subjectId": subjectId})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error Occured while listing notes"})
		return
	}
	err = cursor.All(ctx, &notes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while listing notes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notes": notes, "subject": subject})
}
