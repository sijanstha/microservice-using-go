package http

import (
	"github.com/gin-gonic/gin"
	atdomain "github.com/sijanstha/oauth-api/src/domain/access_token"
	"github.com/sijanstha/oauth-api/src/services/access_token"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atdomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}

	token, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, token)

}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}
