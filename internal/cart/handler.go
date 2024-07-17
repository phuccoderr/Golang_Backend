package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type cartHandler struct {
	service Service
}

func NewCartHandler(service Service) *cartHandler {
	return &cartHandler{service: service}
}

func (h cartHandler) AddToCart(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account ID not found"})
		return
	}
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID not found"})
		return
	}

	quantity, err := strconv.Atoi(c.Query("quantity"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity not found"})
		return
	}

	err = h.service.AddToCart(accountId, productId, quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "CART successfully added"})
}
