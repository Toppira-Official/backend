package errors

import "net/http"

var httpStatusByCode = map[ErrCode]int{
	// Server Errors
	ErrServerInternalError: http.StatusInternalServerError,
	ErrServerNotResponding: http.StatusServiceUnavailable,

	// User Errors
	ErrUserInvalidData:   http.StatusBadRequest,
	ErrUserAlreadyExists: http.StatusConflict,
	ErrUserNotFound:      http.StatusNotFound,
}

func HTTPStatus(code ErrCode) int {
	status, ok := httpStatusByCode[code]
	if !ok {
		status = http.StatusInternalServerError
	}

	return status
}
