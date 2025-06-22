package routes

import (
	"notos/internal/controllers"
	"notos/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func NotesRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/notes", controllers.GetNotes)
	protectedRoutes := incomingRoutes.Group("/notes")
	protectedRoutes.Use(middlewares.Authentication())
	{
		protectedRoutes.GET("/:noteId", controllers.GetNoteById)
		protectedRoutes.GET("/subject/:subjectId", controllers.GetNotesBySubjectId)
		protectedRoutes.POST("/", controllers.UploadNotes)
		protectedRoutes.PATCH("/:noteId", controllers.UpdateNotes)
		protectedRoutes.DELETE("/:noteId", controllers.DeleteNotes)
	}
}

