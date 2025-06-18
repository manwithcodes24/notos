package routes

import (
	"notos/internal/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers)
	incomingRoutes.POST("/users/signup", controllers.Signup)
	incomingRoutes.POST("/users/login", controllers.Login)
	incomingRoutes.GET("/users/:id", controllers.GetUserByID)
}
