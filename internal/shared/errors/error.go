package errors

type ErrCode string //	@name	ErrCode

type ClientError struct {
	Code ErrCode `json:"error" example:"SERVER_NOT_RESPONDING"`
} //	@name	ClientError
type AppError struct {
	ClientError
	Err error
}

func E(code ErrCode, err ...error) *AppError {
	if len(err) == 0 {
		return &AppError{ClientError: ClientError{Code: code}}
	}
	return &AppError{ClientError: ClientError{Code: code}, Err: err[0]}
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return string(e.Code)
}

func (e *AppError) Client() ClientError {
	return e.ClientError
}
