package routes

import (
	"notos/controllers"
	"github.com/gin-gonic/gin"
)

func NotesRoutes (incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/notes", controllers.GetNotes)
	incomingRoutes.GET("/notes/:noteId", controllers.GetNoteById)
	incomingRoutes.GET("/notes/subject/:subjectId", controllers.GetNotesBySubjectId)
	incomingRoutes.POST("/notes", controllers.UploadNotes)
	incomingRoutes.PATCH("/notes/:noteId", controllers.UpdateNotes)
	incomingRoutes.DELETE("/notes/:noteId", controllers.DeleteNotes)
}
