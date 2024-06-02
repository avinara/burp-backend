package errors

import (
	"github.com/burp-backend/constants"
)

type ErrorInterface interface {
	Error() string
	Code() uint32
}

type baseError struct {
	code    uint32
	message string
}

func (err baseError) Error() string {
	return err.message
}

func (err baseError) Code() uint32 {
	return err.code
}

func New(code uint32, message string) ErrorInterface {
	return &baseError{
		code:    code,
		message: message,
	}
}

func InternalServerError() ErrorInterface {
	return New(constants.InternalServerErrorCode, constants.InternalServerError)
}

func LoadingConfigurationFileError() ErrorInterface {
	return New(constants.LoadingConfigurationFileErrorCode, constants.LoadingConfigurationFileError)
}

func QueryParamUnavailableError() ErrorInterface {
	return New(constants.QueryParamUnavailableErrorCode, constants.QueryParamUnavailableError)
}

func CookNotFoundError() ErrorInterface {
	return New(constants.CookNotFoundErrorCode, constants.CookNotFoundError)
}

func InvalidRequestError() ErrorInterface {
	return New(constants.InvalidRequestErrorCode, constants.InvalidRequestError)
}

func UserNotFoundError() ErrorInterface {
	return New(constants.UserNotFoundErrorCode, constants.UserNotFoundError)
}

// Database related errors

func DatabaseInitError() ErrorInterface {
	return New(constants.DatabaseInitErrorCode, constants.DatabaseInitError)
}

func DatabaseDeletionError() ErrorInterface {
	return New(constants.DatabaseDeletionErrorCode, constants.DatabaseDeletionError)
}

func DatabaseUpdationError() ErrorInterface {
	return New(constants.DatabaseUpdationErrorCode, constants.DatabaseUpdationError)
}

func DatabaseInsertionError() ErrorInterface {
	return New(constants.DatabaseInsertionErrorCode, constants.DatabaseInsertionError)
}

func ScanningRowsError() ErrorInterface {
	return New(constants.ScanningRowsErrorCode, constants.ScanningRowsError)
}

func DatabaseQueryError() ErrorInterface {
	return New(constants.DatabaseQueryErrorCode, constants.DatabaseQueryError)
}
