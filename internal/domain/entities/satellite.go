package entities

import "github.com/unawaretub86/top-secret/internal/config/errors"

type (
	Satellite struct {
		Name     string   `json:"name" binding:"required"`
		Distance float32  `json:"distance" binding:"required"`
		Message  []string `json:"message" binding:"required"`
	}

	Satellites struct {
		Satellites []Satellite `json:"satellites"`
	}
)

// Validate se encarga de validar que los parametros sean correctos
func (satellites *Satellites) ValidateSatellites() error {
	if len(satellites.Satellites) != 3 || satellites.Satellites == nil {
		return errors.ErrInvalidSatellites
	}

	return nil
}
