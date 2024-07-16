package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type productHandler struct {
	service service
}

func NewHandler(service service) *productHandler {
	return &productHandler{service: service}
}

func (h productHandler) CreateProduct(c *gin.Context) {
	var product *Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error json": err.Error()})
	}

	err = h.service.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error message": err.Error()})
	}

	c.JSON(http.StatusCreated, "Successfully created product")
}
