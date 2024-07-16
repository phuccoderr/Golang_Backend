package internal

import (
	"myproject/internal/account"
	"myproject/internal/auth"
	"myproject/internal/product"
	"myproject/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	route *gin.Engine
	db    *gorm.DB
}

func NewController(r *gin.Engine, db *gorm.DB) *Controller {
	return &Controller{route: r, db: db}
}

func (c *Controller) SetupRouter() {
	jwt := security.JWT{}

	// Account routes
	accountRepo := account.NewRepository(c.db)
	accountService := account.NewService(accountRepo)
	accountHandler := account.NewHandler(accountService)

	private := c.route.Group("/auth")
	private.Use(c.AuthRequired)
	{
		private.GET("/accounts", accountHandler.ListAccounts)
		private.GET("/accounts/:id", accountHandler.GetAccount)
	}

	// Product routes
	productRepo := product.NewRepository(c.db)
	productService := product.NewService(productRepo)
	productHandler := product.NewHandler(productService)
	{
		c.route.POST("/products", productHandler.CreateProduct)
	}

	// Auth routes
	authService := auth.NewService(accountRepo, jwt)
	authHanlder := auth.NewHandler(authService)

	c.route.POST("/register", authHanlder.Register)
	c.route.POST("/login", authHanlder.Login)
}

func (c *Controller) AuthRequired(g *gin.Context) {

	jwt := security.JWT{}

	tokenString := g.GetHeader("Authorization")
	if tokenString == "" {
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	tokenString = tokenString[7:]

	calims, err := jwt.ParseToken(tokenString)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err = jwt.AuthRoles(calims.Roles)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	g.Next()
}
