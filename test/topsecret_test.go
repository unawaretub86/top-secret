package topsecret_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"

	"github.com/unawaretub86/top-secret/internal/domain/entities"
	"github.com/unawaretub86/top-secret/internal/domain/services"
	mocks_test "github.com/unawaretub86/top-secret/test/mocks"
)

type mocks struct {
	triangulationUseCase *mocks_test.MockTriangulationPort
	messageUseCase       *mocks_test.MockMessagePort
}

func TestTopSecretOK(t *testing.T) {
	m := mocks{
		triangulationUseCase: mocks_test.NewMockTriangulationPort(gomock.NewController(t)),
		messageUseCase:       mocks_test.NewMockMessagePort(gomock.NewController(t)),
	}

	topSecretService := services.NewTopSecretService(m.triangulationUseCase, m.messageUseCase)

	requestID := "test123"

	messages := [][]string{
		{"este", "", "", "mensaje", ""},
		{"", "es", "", "", "secreto"},
		{"este", "", "un", "", ""},
	}

	var d1 float32 = 600.0
	var d2 float32 = 500.0
	var d3 float32 = 716.80

	body := `{
		"satellites": [
			{
			"distance": 600.0,
			"message": ["este", "", "", "mensaje", ""]
			},
			{
			"distance": 500.0,
			"message": ["", "es", "", "", "secreto"]
			},
			{
			"distance": 716.80,
			"message": ["este", "", "un", "", ""]
			}
		]
	}`

	var r1 float32 = -185.12361
	var r2 float32 = 310.74167
	var r3 error = nil
	msg := "este es un mensaje secreto"

	m.triangulationUseCase.EXPECT().GetLocation(requestID, d1, d2, d3).Return(&r1, &r2, r3)
	m.messageUseCase.EXPECT().GetMessage(requestID, messages[0], messages[1], messages[2]).Return(msg, r3)

	response, err := topSecretService.GetLocationAndMessage(body, requestID)

	assert.Equal(t, err, nil)
	assert.Equal(t, response, &entities.LocationMessage{
		X:       &r1,
		Y:       &r2,
		Message: &msg,
	})
}
