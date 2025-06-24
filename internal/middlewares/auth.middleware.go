package middlewares

import (
	// "errors"
	// "net/http"
	// "time"
	"context"
	"net/http"
	"notos/internal/controllers"
	"notos/internal/helpers"
	"notos/internal/models"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, "Authorization header is missing!")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		authToken := strings.Split(auth, " ")
		if len(authToken) != 2 || authToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, "Invalid Token Format!")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authToken[1]
		token, err := helpers.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Invalid Token!")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// check expiration time
		exp, ok := claims["exp"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token expiration!"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if int64(exp) < time.Now().Unix() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// check user id
		userID, err := primitive.ObjectIDFromHex(claims["id"].(string))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token User ID"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var user models.User
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = controllers.UsersCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Can't find user. Try again!"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func AdminAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if *user.(models.User).Role != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
