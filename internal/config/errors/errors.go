package errors

import "errors"

var (
	// Entity errores
	ErrInvalidLocationMessage = errors.New("invalid locationMessage")
	ErrInvalidSatellites      = errors.New("satellites must contain 3 satellites")
)
