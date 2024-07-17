package internal

import (
	"github.com/gin-gonic/gin"
	"myproject/internal/account"
	"myproject/internal/auth"
	"myproject/internal/cart"
	"myproject/internal/product"
	"myproject/pkg/security"
)

type Route struct{}

func (r Route) accountRoute(c *Controller, AuthRequired func(ctx *gin.Context)) {
	// Account routes
	accountRepo := account.NewRepository(c.db)
	accountService := account.NewService(accountRepo)
	accountHandler := account.NewHandler(accountService)

	private := c.route.Group("/auth")
	private.Use(AuthRequired)
	{
		private.GET("/accounts", accountHandler.ListAccounts)
		private.GET("/accounts/:id", accountHandler.GetAccount)
	}
}

func (r Route) productRoute(c *Controller, AuthRequired func(ctx *gin.Context)) {
	productRepo := product.NewRepository(c.db)
	productService := product.NewService(productRepo)
	productHandler := product.NewHandler(productService)
	{
		c.route.POST("/products", productHandler.CreateProduct)
		c.route.GET("/products", productHandler.GetAllProduct)
	}
}

func (r Route) cartRoute(c *Controller, AuthRequired func(ctx *gin.Context)) {
	cartRepo := cart.NewRepository(c.db)
	cartService := cart.NewService(cartRepo)
	cartHandler := cart.NewCartHandler(cartService)
	{
		c.route.POST("/cart/productId/:productId/accountId/:accountId", cartHandler.AddToCart)
	}
}

func (r Route) authRoute(c *Controller, jwt security.JWT) {
	accountRepo := account.NewRepository(c.db)
	authService := auth.NewService(accountRepo, jwt)
	authHanlder := auth.NewHandler(authService)

	c.route.POST("/register", authHanlder.Register)
	c.route.POST("/login", authHanlder.Login)
}
