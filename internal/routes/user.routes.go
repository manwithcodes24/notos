package routes

import (
	"notos/internal/controllers"
	"notos/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers)
	incomingRoutes.POST("/users/signup", controllers.Signup)
	incomingRoutes.POST("/users/login", controllers.Login)
	incomingRoutes.GET("/users/:id", controllers.GetUserByID)
	adminRoutes := incomingRoutes.Group("/admin")
	adminRoutes.Use(middlewares.Authentication())
	adminRoutes.Use(middlewares.AdminAuthentication())
	{
		adminRoutes.GET("/users", controllers.GetUsers)
		adminRoutes.PUT("/users/:userId", controllers.UpdateUser)
		adminRoutes.DELETE("/users/:userId", controllers.DeleteUser)
	}
}
