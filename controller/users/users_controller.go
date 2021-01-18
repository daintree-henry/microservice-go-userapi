package users

import (
	"net/http"

	"github.com/daintree-henry/microservice-go-userapi/domain/users"
	"github.com/daintree-henry/microservice-go-userapi/services"
	"github.com/daintree-henry/microservice-go-userapi/utils/utils_errors"
	"github.com/daintree-henry/microservice-go-userapi/utils/utils_logger"
	"github.com/gin-gonic/gin"
)

func CtrCreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils_logger.Error("json parsing error", err)
		c.JSON(http.StatusBadRequest, utils_errors.NewRestError("json parsing error", http.StatusBadRequest, err.Error()))
		return
	}

	result, err := services.UsersService.SvcCreateUser(user)
	if err != nil {
		c.JSON(err.GetStatusCode(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func CtrGetUserById(c *gin.Context) {
	result, err := services.UsersService.SvcGetUserById(c.Param("id"))
	if err != nil {
		c.JSON(err.GetStatusCode(), err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func CtrUpdateUser(c *gin.Context) {
	isPartial := c.Request.Method == http.MethodPatch

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils_logger.Error("json parsing error", err)
		c.JSON(http.StatusBadRequest, utils_errors.NewRestError("json parsing error", http.StatusBadRequest, err.Error()))
		return
	}

	result, err := services.UsersService.SvcUpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.GetStatusCode(), err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func CtrDeleteUserById(c *gin.Context) {
	if err := services.UsersService.SvcDeleteUserById(c.Param("id")); err != nil {
		c.JSON(err.GetStatusCode(), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

func CtrLoginUserByIdAndPw(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils_logger.Error("json parsing error", err)
		c.JSON(http.StatusBadRequest, utils_errors.NewRestError("json parsing error", http.StatusBadRequest, err.Error()))
		return
	}

	result, err := services.UsersService.SvcLoginUserByIdAndPw(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func CtrFindIdByPhoneNumber(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils_logger.Error("json parsing error", err)
		c.JSON(http.StatusBadRequest, utils_errors.NewRestError("json parsing error", http.StatusBadRequest, err.Error()))
		return
	}

	result, err := services.UsersService.SvcFindIdByPhoneNumber(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func CtrGetUsersByStatus(c *gin.Context) {
	result, err := services.UsersService.SvcGetUsersByStatus(c.Param("status"))

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
