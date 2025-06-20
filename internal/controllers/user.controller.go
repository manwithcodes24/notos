package controllers

import (
	// "notos/database"
	// "notos/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"notos/internal/database"
	"notos/internal/models"
	"time"
    "go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	UsersCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
)



func GetUsers(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Get Users working perfectly",
	})
}

func Signup(c *gin.Context) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	if err := c.BindJSON(&user); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind user data",
		})
		return
	}
	user.ID = primitive.NewObjectID()
	user.Password = HashPassword(user.Password)
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.TotalNotes = 0
	if err := validate.Struct(user); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if _, err := UsersCollection.InsertOne(ctx, user); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error Occured while creating the user",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
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
