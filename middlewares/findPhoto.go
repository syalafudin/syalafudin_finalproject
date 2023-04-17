package middlewares

import (
	"net/http"
	"strconv"

	"github.com/alvinmdj/mygram-api/database"
	"github.com/alvinmdj/mygram-api/models"
	"github.com/gin-gonic/gin"
)

func FindPhoto() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photo := models.Photo{}

		// get route param "photoId"
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
				Error:   "BAD REQUEST",
				Message: "invalid parameter",
			})
			return
		}

		// check if photo exists
		err = db.Debug().First(&photo, photoId).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
				Error:   "NOT FOUND",
				Message: "photo data doesn't exist",
			})
			return
		}
	}
}
