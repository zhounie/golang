package main

import (
	"myapp/internal/database"
	"myapp/internal/models"
	"myapp/internal/routes"

	"github.com/gin-gonic/gin"
)

// @title           IKunCoffee API DOCS
// @version         1.0
// @description     IKunCoffee 接口文档

// @contact.name   Zoey
// @contact.url    http://localhost:8080
// @contact.email  iszhounie@gmail.com

// host localhost:8080
// @BasePath  /api/v1
func main() {
	database.InitDB(&models.User{}, &models.Product{})

	router := gin.Default()

	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		panic("Failed to run server: " + err.Error())
	}

}
