package main

import (
	"github.com/gin-gonic/gin"
	"myproject/internal"
	"myproject/internal/account"
	"myproject/internal/product"
	"myproject/pkg/config"
	"myproject/pkg/database"
)

func main() {
	loadConfig := config.LoadConfig()

	db := database.Connect(loadConfig.DatabaseURL)
	// Create Table
	db.AutoMigrate(&account.Account{})
	db.AutoMigrate(&product.Product{})

	server := gin.Default()
	server.Use(gin.Logger())

	controller := internal.NewController(server, db)
	controller.SetupRouter()

	server.Run(":8080")
}
