package routes

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/slimreaper35/konflux-test/models"
	"github.com/slimreaper35/konflux-test/utils"
)

func SignUpHandler(context *gin.Context) {
	var user models.User

	if context.ShouldBindJSON(&user) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	_, err := models.GetUserBy(user.Email)
	if err == nil {
		context.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		return
	}

	if !isValid(user.Email) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}

	if user.Create() != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	context.JSON(http.StatusCreated, user)
}

func LoginHandler(context *gin.Context) {
	var user models.User

	if context.ShouldBindJSON(&user) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	databaseUser, err := models.GetUserBy(user.Email)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if !utils.ComparePasswords(databaseUser.Password, user.Password) {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(databaseUser.Email, databaseUser.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}

func isValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
