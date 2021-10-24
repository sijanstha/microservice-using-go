package security

import (
	"net/http"
	"strings"
)

var (
	allowedResources = []string{"ping", "auth", "public", "login", "register"}
)

func IsAllowedResource(request *http.Request) bool {
	for _, str := range allowedResources {
		if strings.Contains(request.RequestURI, str) {
			return true
		}
	}

	return false
}
