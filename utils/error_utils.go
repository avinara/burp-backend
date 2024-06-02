package utils

import (
	"github.com/burp-backend/errors"
	"github.com/burp-backend/model"
)

func FormErrorMessage(err errors.ErrorInterface) model.ErrorResponse {
	return model.ErrorResponse{
		Code:         -1,
		ErrorCode:    err.Code(),
		ErrorMessage: err.Error(),
	}
}
