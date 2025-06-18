package helper

import (
	"time"
	"notos/internal/models"
	"os"
	"log"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokens(user models.User) (string, string, error) {
	secret := os.Getenv("JWT_SECRET")
	accessClaims := jwt.MapClaims{
		"id": user.ID,
		"email": user.Email,
		"username": user.Username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	refreshClaims := jwt.MapClaims{
		"id": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}
