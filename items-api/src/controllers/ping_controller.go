package controllers

import (
	"net/http"

	"github.com/sijanstha/common-utils/src/utils/date_utils"
	rest_response "github.com/sijanstha/items-api/src/utils/http"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(http.ResponseWriter, *http.Request)
}

type pingController struct{}

func (p *pingController) Ping(w http.ResponseWriter, r *http.Request) {

	rest_response.OkWithJsonObject(w, http.StatusOK, createResponse())
}

func createResponse() map[string]string {
	response := make(map[string]string)
	response["timestamp"] = date_utils.GetTodayDateInString()
	response["message"] = "pong"
	return response
}
