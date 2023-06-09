package constants

import "errors"

// errors
var (
	ErrBadRequest     = errors.New("bad_request")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrNotFound       = errors.New("not_found")
	ErrInternalServer = errors.New("internal_server_error")
	ErrExisted        = errors.New("existed")
	ErrValidation     = errors.New("validate_fail")
)
