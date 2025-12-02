package main

import (
	"myapp/database"
	"myapp/models"
	"myapp/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB(&models.User{}, &models.Order{})

	router := gin.Default()

	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		panic("Failed to run server: ", err.Error())
	}

}