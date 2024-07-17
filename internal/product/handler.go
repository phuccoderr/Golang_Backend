package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		return
	}

	err = h.service.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Successfully created product")
}

func (h productHandler) GetAllProduct(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 100
	}

	products, err := h.service.GetListProducts(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error json": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
