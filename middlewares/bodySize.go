package middlewares

import (
	"net/http"

	"github.com/alvinmdj/mygram-api/models"
	"github.com/gin-gonic/gin"
)

func BodySizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		maxBodyBytes := int64(2 << 20) // 2 MiB

		var w http.ResponseWriter = c.Writer
		c.Request.Body = http.MaxBytesReader(w, c.Request.Body, maxBodyBytes)

		if c.Request.ContentLength > maxBodyBytes {
			c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, models.ErrorResponse{
				Error:   "REQUEST ENTITY TOO LARGE",
				Message: "request body (file uploaded) too large",
			})
			return
		}

		c.Next()
	}
}
