package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jay-SP/restuarant-management-backend/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.Signup())
	incomingRoutes.POST("/users/login", controller.Login())
}
