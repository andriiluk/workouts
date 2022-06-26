package internal

import "errors"

var (
	ErrBadRequest            = errors.New("bad request")
	ErrBadResponse           = errors.New("bad response")
	ErrInternalService       = errors.New("internal service error")
	ErrDecodeRequest         = errors.New("error decoding request")
	ErrMissingRequiredParams = errors.New("missing required parameters")
)
