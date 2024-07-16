package router

import (
	"github.com/gin-gonic/gin"
	"go-web/controller"
)

func InitRoutes(r *gin.Engine) {
	// User routes
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controller.GetUsers)
		userRoutes.GET(":id", controller.GetUserByID)
		userRoutes.POST("/", controller.CreateUser)
	}
}
