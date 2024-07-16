package auth

import (
	"myproject/internal/account"
	"myproject/model/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service Service
}

func NewHandler(service Service) *authHandler {
	return &authHandler{service: service}
}

func (h *authHandler) Register(c *gin.Context) {
	var account *account.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error json": err.Error()})
		return
	}

	err := h.service.CreateAccount(account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Successfully created account")
}

func (h *authHandler) Login(c *gin.Context) {
	var signin *request.SignIn
	if err := c.ShouldBindJSON(&signin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(signin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": token})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})

}
