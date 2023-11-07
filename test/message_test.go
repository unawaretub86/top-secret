package topsecret_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/unawaretub86/top-secret/internal/config/errors"
	"github.com/unawaretub86/top-secret/internal/domain/usecase"
)

func TestCalculateMessageOk(t *testing.T) {
	messageUseCase := usecase.NewMessageUseCase()

	requestID := "test123"

	messages := [][]string{
		{"este", "", "", "mensaje", ""},
		{"", "es", "", "", "secreto"},
		{"este", "", "un", "", ""},
	}

	message, err := messageUseCase.GetMessage(requestID, messages[0], messages[1], messages[2])

	msg := "este es un mensaje secreto"

	assert.Equal(t, message, msg)
	assert.Equal(t, err, nil)
}

func TestCalculateMessageErrInvalidNumberSatellites(t *testing.T) {
	messageUseCase := usecase.NewMessageUseCase()

	requestID := "test123"

	message, err := messageUseCase.GetMessage(requestID)

	assert.Equal(t, err, errors.ErrInvalidSatellites)
	assert.Equal(t, message, "")
}
