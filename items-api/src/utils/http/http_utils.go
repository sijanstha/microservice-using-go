package rest_response

import (
	"encoding/json"
	"net/http"

	"github.com/sijanstha/common-utils/src/utils/errors"
)

func OkWithJsonObject(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func Error(w http.ResponseWriter, err errors.RestErr) {
	OkWithJsonObject(w, err.Code, err)
}

func Ok(w http.ResponseWriter, statusCode int) {
	OkWithJsonObject(w, statusCode, nil)
}
