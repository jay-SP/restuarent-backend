package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jay-SP/restuarant-management-backend/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:food_id", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.PATCH("/menus/:food_id", controller.UpdateMenu())
}
