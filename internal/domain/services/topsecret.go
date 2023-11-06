package services

import (
	"fmt"

	"github.com/unawaretub86/top-secret/internal/config/errors"
	"github.com/unawaretub86/top-secret/internal/domain/entities"
	"github.com/unawaretub86/top-secret/internal/domain/ports"
	"github.com/unawaretub86/top-secret/internal/domain/request"
)

type topSecretService struct {
	triangulationUseCase ports.TriangulationPort
	messageUseCase       ports.MessagePort
}

func NewTopSecretService(triangulation ports.TriangulationPort, message ports.MessagePort) *topSecretService {
	return &topSecretService{
		triangulationUseCase: triangulation,
		messageUseCase:       message,
	}
}

// GetLocationAndMessage se en carga de llamar a la logica de calcular mensaje, y calcular posición
func (topSecretService *topSecretService) GetLocationAndMessage(body string, requestID string) (*entities.LocationMessage, error) {
	// llamamos a la funcion alojada en request para convertir el string que recibimos desde el el evento a la entidad satellites
	satellites, err := request.Decode(body, requestID)
	if err != nil {
		return nil, err
	}

	// obtenemos la informacion particular de cada satellite
	satelite := satellites.Satellites

	if len(satellites.Satellites) != 3 {
		fmt.Printf("[RequestId: %s][%v]", requestID, errors.ErrInvalidSatellites)
		return nil, errors.ErrInvalidSatellites
	}

	// Llama a la función GetLocation del caso de uso 'triangulationUseCase' para determinar la ubicación (x, y)
	// a partir de las distancias recibidas de tres satelites: satelite[0], satelite[1] y satelite[2].
	x, y := topSecretService.triangulationUseCase.GetLocation(requestID, satelite[0].Distance, satelite[1].Distance, satelite[2].Distance)

	// Llama a la función GetMessage del caso de uso 'messageUseCase' para obtener el mensaje secreto.
	// Utiliza el 'requestID' para identificar el mensaje y las palabras recibidas de los satelites satelite[0], satelite[1] y satelite[2].
	message := topSecretService.messageUseCase.GetMessage(requestID, satelite[0].Message, satelite[1].Message, satelite[2].Message)

	return entities.NewLocationMessage(x, y, message, requestID)
}
