package entities

import (
	"fmt"

	"github.com/unawaretub86/top-secret/internal/config/errors"
)

type LocationMessage struct {
	X       *float32
	Y       *float32
	Message *string
}

// NewLocationMessage crea una nueva entidad de locationMessage
func NewLocationMessage(x, y float32, message, requestID string) (*LocationMessage, error) {
	locationMessage := &LocationMessage{
		X:       &x,
		Y:       &y,
		Message: &message,
	}

	if err := locationMessage.Validate(); err != nil {
		fmt.Printf("[RequestId: %s][%v]", requestID, errors.ErrInvalidLocationMessage)
		return nil, err
	}

	return locationMessage, nil
}

// Validate se encarga de validar que los parametros sean correctos
func (locationMessage *LocationMessage) Validate() error {
	if locationMessage.X == nil || locationMessage.Y == nil || locationMessage.Message == nil {
		return errors.ErrInvalidLocationMessage
	}

	return nil
}
