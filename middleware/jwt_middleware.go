package middleware

import (
	"project/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GenerateToken(userId int, name string) (string, error) {
	// payload data
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	//header (metode algoritma)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.KEY_JWT))
}

func ExtractUserIdFromJWT(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return -1
}
