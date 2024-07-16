package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type accountHandler struct {
	service Service
}

func NewHandler(service Service) *accountHandler {
	return &accountHandler{service: service}
}

func (h *accountHandler) ListAccounts(c *gin.Context) {
	accounts, err := h.service.ListAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

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
