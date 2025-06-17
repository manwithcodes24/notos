package middlewares

import (
	// "errors"
	// "net/http"
	// "time"

	"github.com/gin-gonic/gin"

)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}