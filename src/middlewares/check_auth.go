package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	config "github.com/Vlad-Peresta/todo_list_go/src/conf"
	"github.com/Vlad-Peresta/todo_list_go/src/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CheckAuth(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing."})
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token format."})
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token."})
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		context.Abort()
		return
	}
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired."})
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	if err := config.DB.Find(&user, "id = ?", claims["id"]).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.ID == 0 {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	context.Set("CurrentUser", user)

	context.Next()
}
