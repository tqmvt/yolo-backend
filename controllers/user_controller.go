package controllers

import (
	"net/http"
	"saas/models"
	"saas/repository"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	var user []*models.User
	repository.Get(&user)
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
