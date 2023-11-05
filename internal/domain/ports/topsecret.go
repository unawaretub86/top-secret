package ports

import "github.com/unawaretub86/top-secret/internal/domain/entities"

type TopSecretPort interface {
	GetLocationAndMessage(string, string) (*entities.LocationMessage, error)
}
