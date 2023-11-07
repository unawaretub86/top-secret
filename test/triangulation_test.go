package topsecret_test

import (
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/unawaretub86/top-secret/internal/config/errors"
	"github.com/unawaretub86/top-secret/internal/domain/usecase"
)

func TestTriangulatePositionOk(t *testing.T) {
	trilaterationUseCase := usecase.NewTriangulationUseCase()

	requestID := "test123"

	var d1 float32 = 600.0
	var d2 float32 = 500.0
	var d3 float32 = 716.80

	var r1 float32 = -185.12361
	var r2 float32 = 310.74167

	x, y, err := trilaterationUseCase.GetLocation(requestID, d1, d2, d3)

	assert.Equal(t, err, nil)
	assert.Equal(t, x, r1)
	assert.Equal(t, y, r2)
}

func TestTriangulatePositionErrInvalidNumberSatellites(t *testing.T) {
	triangulationUseCase := usecase.NewTriangulationUseCase()

	requestID := "test123"

	_, _, err := triangulationUseCase.GetLocation(requestID)

	assert.Equal(t, err, errors.ErrInvalidSatellites)
}
