package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sijanstha/common-utils/src/oauth"
	"github.com/sijanstha/common-utils/src/utils/errors"
	"github.com/sijanstha/items-api/src/domain/items"
	"github.com/sijanstha/items-api/src/services"
	rest_response "github.com/sijanstha/items-api/src/utils/http"
)

var (
	ItemController itemsControllerInterface = &itemController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemController struct{}

func (i *itemController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		rest_response.Error(w, *err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid request body")
		rest_response.Error(w, *restErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		restErr := errors.NewBadRequestError("Invalid item json body")
		rest_response.Error(w, *restErr)
		return
	}

	itemRequest.Seller = oauth.GetClientId(r)

	result, restErr := services.NewItemService().Create(itemRequest)
	if restErr != nil {
		rest_response.Error(w, *restErr)
		return
	}

	rest_response.OkWithJsonObject(w, http.StatusCreated, result)
}

func (i *itemController) Get(w http.ResponseWriter, r *http.Request) {
	rest_response.Ok(w, http.StatusNotImplemented)
}
