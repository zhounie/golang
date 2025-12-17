package routes

import (
	"myapp/internal/controllers"
	"myapp/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "myapp/docs"
)

func SetupRoutes(router *gin.Engine) {
	userCtrl := controllers.NewUserController()
	authCtrl := controllers.NewAuthController()
	// orderCtrl := controllers.NewOrderController()

	v1 := router.Group("/api/v1")
	auth := v1.Group("/auth")
	{
		auth.POST("/login", authCtrl.Login)
	}
	v1.Use(middleware.AuthMiddleware())
	{
		users := v1.Group("/users")
		users.POST("/", userCtrl.CreateUser)
		users.GET("/:id", userCtrl.GetUserByID)
		// orders := v1.Group("/orders")
		// {
		// 	orders.POST("/", orderCtrl.CreateOrder)
		// 	orders.GET("/:id", orderCtrl.GetOrderByID)
		// }
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
