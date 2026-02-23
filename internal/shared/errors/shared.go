package errors

const (
	ErrServerInternalError           ErrCode = "SERVER_INTERNAL_ERROR"
	ErrServerNotResponding           ErrCode = "SERVER_NOT_RESPONDING"
	ErrTooManyRequests               ErrCode = "TOO_MANY_REQUESTS"
	ErrServiceTemporarilyUnavailable ErrCode = "SERVICE_TEMPORARILY_UNAVAILABLE"
)
