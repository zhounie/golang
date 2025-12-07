package routes

import (
	"myapp/internal/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "myapp/docs"
)

func SetupRoutes(router *gin.Engine) {
	userCtrl := controllers.NewUserController()
	// orderCtrl := controllers.NewOrderController()

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", userCtrl.CreateUser)
			users.GET("/:id", userCtrl.GetUserByID)
		}
		// orders := v1.Group("/orders")
		// {
		// 	orders.POST("/", orderCtrl.CreateOrder)
		// 	orders.GET("/:id", orderCtrl.GetOrderByID)
		// }
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
