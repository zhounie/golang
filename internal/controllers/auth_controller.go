package controllers

import (
	"myapp/internal/models"
	"myapp/internal/services"
	"myapp/pkg/jwt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{Service: services.NewAuthService()}
}

func (ctrl *AuthController) Login(c *gin.Context) (interface{}, error) {
	var loginUser models.LoginRequest
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		return nil, err
	}

	user, err := ctrl.Service.Login(&loginUser)
	if err != nil {
		return nil, err
	}
	token, err := jwt.GenerateToken(strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		return nil, err
	}
	return token, nil
}
