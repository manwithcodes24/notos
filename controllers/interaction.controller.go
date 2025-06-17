package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetLikes(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Likes working perfectly",
	})
}

func GetComments(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Comments working perfectly",
	})
}

func LikeNote(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Like Note working perfectly",
	})
}

func CommentNote(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Comment Note working perfectly",
	})
}

func DeleteComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Comment working perfectly",
	})
}

func DeleteLike(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Like working perfectly",
	})
}