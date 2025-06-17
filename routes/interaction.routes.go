package routes

import (
	"notos/controllers"
	"github.com/gin-gonic/gin"
)

func InteractionRoutes (incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/likes/:notesId", controllers.GetLikes)
	incomingRoutes.GET("/comments/:notesId", controllers.GetComments)
	incomingRoutes.POST("/like/:notesId", controllers.LikeNote)
	incomingRoutes.POST("/comment/:notesId", controllers.CommentNote)
	incomingRoutes.DELETE("/comment/:notesId/:commentId", controllers.DeleteComment)
	incomingRoutes.DELETE("/like/:likeId", controllers.DeleteLike)
}

