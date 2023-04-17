package middlewares

import (
	"net/http"

	"github.com/alvinmdj/mygram-api/helpers"
	"github.com/alvinmdj/mygram-api/models"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "UNAUTHENTICATED",
				Message: err.Error(),
			})
			return
		}

		// store token claims in request data
		c.Set("userData", verifyToken)
		c.Next()
	}
}
