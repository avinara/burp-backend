package utils

import (
	"encoding/json"
	"net/http"

	"github.com/burp-backend/model"
)

func WriteJSON(w http.ResponseWriter, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func WriteErrorWithMessage(w http.ResponseWriter, errorResponse model.ErrorResponse) error {
	statusCode := GetHttpStatus(errorResponse.ErrorCode)
	w.WriteHeader(statusCode)
	return WriteJSON(w, errorResponse)
}

func GetHttpStatus(code uint32) (status int) {
	firstThreeDigits := code / 100
	switch firstThreeDigits {
	case 400:
		status = http.StatusBadRequest
	case 401:
		status = http.StatusUnauthorized
	case 403:
		status = http.StatusForbidden
	case 404:
		status = http.StatusNotFound
	case 405:
		status = http.StatusMethodNotAllowed
	case 406:
		status = http.StatusNotAcceptable
	case 408:
		status = http.StatusRequestTimeout
	case 424:
		status = http.StatusFailedDependency
	default:
		status = http.StatusInternalServerError
	}
	return
}
