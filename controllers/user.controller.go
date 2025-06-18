package controllers

import (
	// "notos/internal/database"
	// "notos/internal/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Get Users working perfectly",
	})
}

func Signup(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Signup working perfectly",
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login working perfectly",
	})
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("login or passoword is incorrect")
		check = false
	}
	return check, msg
}

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get User working perfectly",
	})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Get User by ID working perfectly id: " + id,
	})
}
