package controllers

import (
	// "notos/database"
	// "notos/models"
	"context"
	"log"
	"net/http"

	"notos/internal/helpers"
	"notos/internal/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
		log.Println(err)
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
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if _, err := UsersCollection.InsertOne(ctx, user); err != nil {
		log.Println(err)
		if mongo.IsDuplicateKeyError(err) {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Duplicate Data. Can't signup!",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error Occured while creating the user. Try again!",
		})
		return
	}
	accessToken, refreshToken, err := helpers.GenerateTokens(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error Occured while generating the tokens. Try again!",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data": map[string]string{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}

func Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind user data",
		})
		return
	}
	var foundUser models.User
	if err := UsersCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User doesn't exist. Create new account!",
		})
		return
	}
	if !VerifyPassword(user.Password, foundUser.Password) {
		c.JSON(http.StatusForbidden, gin.H{
			"Error!": "Wrong password. Try again!",
		})
		return 
	}
	accessToken, refreshToken, err := helpers.GenerateTokens(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error Occured while generating the tokens. Try again!",
		})
		return
	}
	var refreshmodel models.RefreshToken
	refreshmodel.ID = primitive.NewObjectID()
	refreshmodel.Token = refreshToken
	refreshmodel.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	refreshmodel.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	refreshmodel.UserID = foundUser.ID
	refreshmodel.Validated = true
	refreshmodel.ExpiresAt, _ = time.Parse(time.RFC3339, time.Now().Add(time.Hour*24*7).Format(time.RFC3339))
	if _, err := RefreshTokenCollection.InsertOne(ctx, refreshmodel); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error Occured while saving the refresh token. Try again!",
		})
		return
	}
	c.SetCookie("accessToken", accessToken, 3600*24*30, "/", "localhost", false, true)
	c.SetCookie("refreshToken", refreshToken, 3600*24*30, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "User signedin successfully",
		"data": map[string]string{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}
func VerifyPassword(userPassword string, providedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword)); err != nil {
		return false
	}
	return true
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
