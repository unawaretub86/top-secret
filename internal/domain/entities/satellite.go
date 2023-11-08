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

func (satellites *Satellites) Validate() error {
	for _, satellite := range satellites.Satellites {
		if satellite.Distance == 0 || len(satellite.Message) == 0 {
			return errors.ErrInvalidSatellites
		}
	}

	return nil
}
