package main

import (
	"log"
	"notos/middlewares"
	"notos/routes"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middlewares.Authentication())
	routes.NotesRoutes(router)
	routes.SubjectRoutes(router)
	routes.InteractionRoutes(router)
	router.Run(":" + port)
}
