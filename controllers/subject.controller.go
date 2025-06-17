package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetSubjects(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Subjects working perfectly",
	})
}

func GetSubjectById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Subject by ID working perfectly",
	})
}

func CreateSubject(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create Subject working perfectly",
	})
}

func UpdateSubject(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update Subject working perfectly",
	})
}

func DeleteSubject(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Subject working perfectly",
	})
}

