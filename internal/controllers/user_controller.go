package controllers

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"myapp/internal/models"
	"myapp/internal/services"
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
func (ctrl *UserController) CreateUser(c *gin.Context) (interface{}, error) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return nil, err
	}

	if err := ctrl.Service.CreateUser(&user); err != nil {
		return nil, err
	}
	return user, nil
}

// @Summary      获取用户
// @Description  根据ID获取用户
// @param       id query    string  true  "id"
// @Tags         users
// @Produce      json
// @Router       /api/v1/users/:id [post]
// @Success 200 {object} models.User
func (ctrl *UserController) GetUserByID(c *gin.Context) (interface{}, error) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		return nil, errors.New("Invaild user ID")
	}

	user, err := ctrl.Service.GetUserByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("User not found")
		}
		return nil, errors.New("Database err")
	}
	return user, nil
}

// @Summary      删除用户
// @Description  根据ID删除用户
// @param       ids query    string  true  "ids"
// @Tags         users
// @Produce      json
// @Router       /api/v1/users [delete]
// @Success 200 {object} models.User
func (ctrl *UserController) DeleteUser(c *gin.Context) (interface{}, error) {
	var req models.DeleteRequest
	log.Printf(c.Param("ids"))
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	if err := ctrl.Service.DeleteUser(req.IDS); err != nil {
		return nil, err
	}
	return "删除成功", nil
}
