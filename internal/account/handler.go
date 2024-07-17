package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

type accountHandler struct {
	service Service
}

func NewHandler(service Service) *accountHandler {
	return &accountHandler{service: service}
}

// GetListA godoc
// @Summary Get List
// @Description do ping
// @Success 200 {string} Account
// @Router /auth/accounts [get]
func (h *accountHandler) ListAccounts(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 100
	}

	accounts, err := h.service.ListAccounts(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

// GetListA godoc
// @Summary Get List
// @Description do ping
// @Success 200 {string} Account
// @Router /auth/accounts/:id [get]
func (h *accountHandler) GetAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	account, err := h.service.GetAccount(int(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": account})
}
