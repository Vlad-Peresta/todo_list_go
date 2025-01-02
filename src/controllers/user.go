package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Vlad-Peresta/todo_list_go/src/models"
	"github.com/Vlad-Peresta/todo_list_go/src/schemas"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser godoc
//
//	@Summary		Create User record
//	@Description	Create User record
//	@Tags			auth
//	@Produce		json
//	@Param			Request Body 	body		schemas.AuthInputData  	true	"Authentication Data"
//	@Success		200	{object}	schemas.AuthInputData
//	@Failure		400	{object}	error
//	@Router			/auth/signup [POST]
//
// CreateUser creates User record in the database
func CreateUser(context *gin.Context) {
	var authData schemas.AuthInputData
	var user models.User

	// Binding JSON request body to AuthInputData struct
	if err := context.ShouldBindJSON(&authData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Finding User record with provided Username
	models.GetUserByUsername(&user, authData.Username)
	if user.ID != 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User with provided Username is already used."})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authData.Password), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user = models.User{
		Username: authData.Username,
		Password: string(passwordHash),
	}
	if err := models.CreateRecord(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Create HTTP response
	context.JSON(http.StatusOK, gin.H{"data": user})
}

// Login godoc
//
//	@Summary		Login User
//	@Description	Login User
//	@Tags			auth
//	@Produce		json
//	@Param			Request Body 	body		schemas.AuthInputData  	true	"Authentication Data"
//	@Success		200	{object}	schemas.AuthInputData
//	@Failure		400	{object}	error
//	@Router			/auth/login [POST]
//
// Login authenticates User
func Login(context *gin.Context) {
	var authData schemas.AuthInputData
	var user models.User

	if err := context.ShouldBindJSON(&authData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.GetUserByUsername(&user, authData.Username); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authData.Password)); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password."})
		return
	}

	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generatedToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate token."})
		return
	}

	context.JSON(200, gin.H{"token": token})
}
