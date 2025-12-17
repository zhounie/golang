package controllers

import (
	"myapp/internal/models"
	"myapp/internal/services"
	"myapp/pkg/jwt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{Service: services.NewAuthService()}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var loginUser models.LoginRequest
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.Service.Login(&loginUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := jwt.GenerateToken(strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
}
