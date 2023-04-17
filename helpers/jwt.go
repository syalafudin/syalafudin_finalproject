package helpers

import (
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("JWT_SECRET")

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	// creates a new token with the specified signing method and claims.
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// creates and returns a complete, signed JWT.
	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	// init error message
	errResponse := errors.New("sign in to proceed")
	// get Authorization header value
	headerToken := c.Request.Header.Get("Authorization")
	// check if Authorization header contains Bearer as a suffix
	if bearer := strings.HasPrefix(headerToken, "Bearer"); !bearer {
		return nil, errResponse
	}

	// headerToken: Bearer <token-here>
	// get the <token-here> value after splitting inside index 1
	stringToken := strings.Split(headerToken, " ")[1]

	// parse token into a pointer of struct jwt.Token
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		// check if signing method is HS256 by casting the method into pointer of struct jwt.SigningMethodHMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	// check if token still valid after casting into type of jwt.MapClaims
	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, errResponse
	}

	// return claims (contains id & email of the successfully logged in user),
	return token.Claims.(jwt.MapClaims), nil
}
