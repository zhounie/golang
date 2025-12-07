package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"myapp/internal/models"
	"myapp/internal/services"
	"net/http"
	"strconv"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController() *UserController {
	return &UserController{Service: services.NewUserService()}
}

// @Summary      创建用户
// @Description  创建用户
// @param       name query    string  true  "name"
// @Tags         users
// @Produce      json
// @Router       /api/v1/users [get]
// @Success 200 {object} models.User
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// @Summary      获取用户
// @Description  根据ID获取用户
// @param       id query    string  true  "id"
// @Tags         users
// @Produce      json
// @Router       /api/v1/users/:id [post]
// @Success 200 {object} models.User
func (ctrl *UserController) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invaild user ID"})
		return
	}

	user, err := ctrl.Service.GetUserByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database err"})
		return
	}

	c.JSON(http.StatusOK, user)

}
