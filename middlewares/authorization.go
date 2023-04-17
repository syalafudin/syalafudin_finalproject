package middlewares

import (
	"net/http"
	"strconv"

	"github.com/alvinmdj/mygram-api/database"
	"github.com/alvinmdj/mygram-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		// get route param "socialMediaId"
		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
				Error:   "BAD REQUEST",
				Message: "invalid parameter",
			})
			return
		}

		// get token claims, which is set in authentication middleware
		userData := c.MustGet("userData").(jwt.MapClaims)

		// get user id from token claims
		userId := uint(userData["id"].(float64))
		socialMedia := models.SocialMedia{}

		// get user_id column from social media table with the associated social media id
		err = db.Debug().Select("user_id").First(&socialMedia, socialMediaId).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
				Error:   "NOT FOUND",
				Message: "data doesn't exist",
			})
			return
		}

		// check if user id from db of the associated social media == user id from token claims
		if socialMedia.UserID != userId {
			c.AbortWithStatusJSON(http.StatusForbidden, models.ErrorResponse{
				Error:   "FORBIDDEN",
				Message: "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		// get route param "photoId"
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
				Error:   "BAD REQUEST",
				Message: "invalid parameter",
			})
			return
		}

		// get token claims, which is set in authentication middleware
		userData := c.MustGet("userData").(jwt.MapClaims)

		// get user id from token claims
		userId := uint(userData["id"].(float64))
		photo := models.Photo{}

		// get user_id column from photo table with the associated photo id
		err = db.Debug().Select("user_id").First(&photo, photoId).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
				Error:   "NOT FOUND",
				Message: "data doesn't exist",
			})
			return
		}

		// check if user id from db of the associated photo == user id from token claims
		if photo.UserID != userId {
			c.AbortWithStatusJSON(http.StatusForbidden, models.ErrorResponse{
				Error:   "FORBIDDEN",
				Message: "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		// get route param "commentId"
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
				Error:   "BAD REQUEST",
				Message: "invalid parameter",
			})
			return
		}

		// get token claims, which is set in authentication middleware
		userData := c.MustGet("userData").(jwt.MapClaims)

		// get user id from token claims
		userId := uint(userData["id"].(float64))
		comment := models.Comment{}

		// get user_id column from comment table with the associated comment id
		err = db.Debug().Select("user_id").First(&comment, commentId).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
				Error:   "NOT FOUND",
				Message: "data doesn't exist",
			})
			return
		}

		// check if user id from db of the associated comment == user id from token claims
		if comment.UserID != userId {
			c.AbortWithStatusJSON(http.StatusForbidden, models.ErrorResponse{
				Error:   "FORBIDDEN",
				Message: "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}
