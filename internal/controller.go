package internal

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"myproject/pkg/security"

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
	//Swagger
	c.route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	jwt := security.JWT{}
	router := Route{}

	// Auth routes
	router.authRoute(c, jwt)

	router.accountRoute(c, c.AuthRequired)
	router.productRoute(c, c.AuthRequired)
	router.cartRoute(c, c.AuthRequired)
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
