package main

import (
	"myapp/internal/database"
	"myapp/internal/models"
	"myapp/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB(&models.User{})

	router := gin.Default()

	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		panic("Failed to run server: " + err.Error())
	}

}
