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

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Code, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "WIP!")
}

func DeleteUser(c *gin.Context){
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Code, err)
		return
	}

	deleteErr := services.DeleteUser(userId)
	if deleteErr != nil {
		c.JSON(deleteErr.Code, deleteErr)
		return
	}
	c.JSON(http.StatusOK, "Deleted Successfully")
}

func UpdateUser(c *gin.Context)  {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
