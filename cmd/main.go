package main

import (
	"gorm/internal/config"
	"gorm/internal/models"
	"gorm/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Contact{})

	r := gin.Default()
	routes.Routes(r)

	r.Run(":8080")
}