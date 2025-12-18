package routes

import (
	"myapp/internal/controllers"
	"myapp/internal/middleware"
	"myapp/pkg/response"

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
		auth.POST("/login", response.Wrapper(authCtrl.Login))
	}
	v1.Use(middleware.AuthMiddleware())
	{
		users := v1.Group("/users")
		users.POST("/", response.Wrapper(userCtrl.CreateUser))
		users.GET("/:id", response.Wrapper(userCtrl.GetUserByID))
		users.DELETE("/", response.Wrapper(userCtrl.DeleteUser))
		// orders := v1.Group("/orders")
		// {
		// 	orders.POST("/", orderCtrl.CreateOrder)
		// 	orders.GET("/:id", orderCtrl.GetOrderByID)
		// }
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
