package routes

import (
	"notos/controllers"
	"github.com/gin-gonic/gin"
)


func SubjectRoutes (incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/subjects", controllers.GetSubjects)
	incomingRoutes.GET("/subject/:subjectId", controllers.GetSubjectById)
	incomingRoutes.POST("/subjects", controllers.CreateSubject)
	incomingRoutes.PATCH("/subject/:subjectId", controllers.UpdateSubject)
	incomingRoutes.DELETE("/subject/:subjectId", controllers.DeleteSubject)
}
