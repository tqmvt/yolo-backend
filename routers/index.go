package routers

import (
	"net/http"

	"saas/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {
	// route not exist
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})

	// health check
	route.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })

	// users routes
	route.GET("/v1/users", controllers.GetUsers)
	route.POST("/v1/users", controllers.CreateUser)
}
