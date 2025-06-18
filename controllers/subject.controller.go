package controllers

import (
	"context"
	"log"
	"net/http"
	"notos/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)



func GetSubjects(c *gin.Context) {
	var Subjects []models.Subject
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	cursor, err := SubjectsCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while finding subjects"})
		return
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &Subjects); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while fetching subjects"})
		return
	}
	c.JSON(http.StatusOK, Subjects)
}

func GetSubjectById(c *gin.Context) {
	subjectId, err := primitive.ObjectIDFromHex(c.Param("subjectId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while parsing subjectId"})
		return
	}
	var Subject models.Subject
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err = SubjectsCollection.FindOne(ctx, bson.M{"_id": subjectId}).Decode(&Subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while finding subject"})
		return
	}
	c.JSON(http.StatusOK, Subject)
}

func CreateSubject(c *gin.Context) {
	var subject models.Subject
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	if err := c.BindJSON(&subject); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while binding JSON"})
		return
	}
	valErr := validate.Struct(subject)
	if valErr != nil {
		log.Fatal(valErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while validating JSON"})
		return
	}
	subject.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	subject.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	subject.ID = primitive.NewObjectID()
	_, err := SubjectsCollection.InsertOne(ctx, subject)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while inserting subject"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Subject created successfully"})
}

func UpdateSubject(c *gin.Context) {
	var subject models.Subject
	subjectId, err := primitive.ObjectIDFromHex(c.Param("subjectId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while parsing subjectId"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	if err := c.BindJSON(&subject); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while binding JSON"})
		return
	}
	valErr := validate.Struct(subject)
	if valErr != nil {
		log.Fatal(valErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while validating JSON"})
		return
	}
	subject.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	err = SubjectsCollection.FindOneAndUpdate(ctx, bson.M{"_id": subjectId}, bson.M{"$set": subject}).Decode(&subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while updating subject"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subject updated successfully"})
}

func DeleteSubject(c *gin.Context) {
	subjectId, err := primitive.ObjectIDFromHex(c.Param("subjectId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while parsing subjectId"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	_, err = SubjectsCollection.DeleteOne(ctx, bson.M{"_id": subjectId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while deleting subject"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subject deleted successfully"})
}

