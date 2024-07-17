package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	docs "myproject/cmd/server/docs" // swagger docs

	"myproject/internal"
	"myproject/internal/account"
	"myproject/internal/cart"
	"myproject/internal/product"
	"myproject/pkg/config"
	"myproject/pkg/database"
)

// @title My Project API
// @version 1.0
// @description This is a sample server for a my project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	loadConfig := config.LoadConfig()

	db := database.Connect(loadConfig.DatabaseURL)
	// Create Table
	db.AutoMigrate(&account.Account{})
	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&cart.Cart{})

	server := gin.Default()
	server.Use(gin.Logger())

	docs.SwaggerInfo.BasePath = "/api/v1"
	controller := internal.NewController(server, db)
	controller.SetupRouter()

	server.Run(":8080")
}
