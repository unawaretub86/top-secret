package errors

import "errors"

var (
	ErrInvalidLocationMessage = errors.New("invalid locationMessage")
	ErrInvalidSatellites      = errors.New("satellites must contain 3 satellites")
	ErrInvalidSatellite       = errors.New("all fields must be full")
)
