package main

import (
	"gorm/internal/config"
	"gorm/internal/models"
	"gorm/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r:=gin.Default()


	config.ConnectDB()
	config.DB.AutoMigrate(&models.Contact{})


	routes.Routes(r)

	r.Run(":8080")

}