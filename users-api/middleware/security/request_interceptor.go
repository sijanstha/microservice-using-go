package security

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sijanstha/common-utils/src/logger"
	"github.com/sijanstha/common-utils/src/oauth"
	"github.com/sijanstha/common-utils/src/utils/errors"
)

func SecurityInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Method == http.MethodOptions {
			return
		}

		logger.Debug("Inside Security Interceptor :: " + c.Request.RequestURI + " :: " + c.Request.Method)

		if !oauth.IsPublic(c.Request) {
			if !IsAllowedResource(c.Request) {
				logger.Debug("========== Checking token ==========")
				at := c.Request.Header.Get("X-CSRF")
				logger.Debug(fmt.Sprintf("Token is: %s", at))

				err := oauth.AuthenticateRequest(c.Request)
				if err != nil {
					c.JSON(http.StatusUnauthorized, err)
					c.Abort()
				}

			}
		} else {
			if c.Request.Method != http.MethodGet {
				restErr := errors.RestErr{
					Code:    http.StatusUnauthorized,
					Message: "Please login",
					Error:   "unauthorized",
				}
				c.JSON(restErr.Code, restErr)
				c.Abort()
			}
		}

		c.Next()
	}
}
