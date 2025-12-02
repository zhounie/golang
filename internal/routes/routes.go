package routers

import (
	"myapp/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userCtrl := controllers.NewUserController()
	orderCtrl := controllers.NewOrderController()

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", userCtrl.CreateUser)
			users.GET("/:id", userCtrl.GetUserByID)
		}
		orders := v1.Group("/orders")
		{
			orders.POST("/", orderCtrl.CreateOrder)
			orders.GET("/:id", orderCtrl.GetOrderByID)
		}
	}

}