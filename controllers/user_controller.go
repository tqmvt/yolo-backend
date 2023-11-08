package controllers

import (
	"net/http"
	"saas/models"
	"saas/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	var user []*models.User
	repository.Get(&user)
	ctx.JSON(http.StatusOK, &user)
}

func GetUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	var user *models.User

	// Convert the user ID to an integer
	id, err := strconv.Atoi(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Retrieve the user by ID
	err = repository.GetUserByID(&user, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, &user)
}

func CreateUser(ctx *gin.Context) {
	user := new(models.User)

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Input correct user information!",
		})
		return
	}

	repository.Save(&user)
	ctx.JSON(http.StatusOK, &user)
}
