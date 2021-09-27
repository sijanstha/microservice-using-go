package users

import (
	"github.com/gin-gonic/gin"
	"github.com/sijanstha/domain/users"
	"github.com/sijanstha/services"
	"github.com/sijanstha/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}

	result, err := services.UserService.CreateUser(user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-PUBLIC") == "true"))
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Code, err)
		return
	}

	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-PUBLIC") == "true"))
}

func DeleteUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Code, err)
		return
	}

	deleteErr := services.UserService.DeleteUser(userId)
	if deleteErr != nil {
		c.JSON(deleteErr.Code, deleteErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "Deleted"})
}

func UpdateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UserService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-PUBLIC") == "true"))
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UserService.SearchUser(status)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-PUBLIC") == "true"))
}

func Login(c *gin.Context) {
	var req users.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}
	user, err := services.UserService.FindUserForAuthentication(&req)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-PUBLIC") == "true"))
}
