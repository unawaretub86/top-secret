package request

import (
	"encoding/json"
	"fmt"

	"github.com/unawaretub86/top-secret/internal/domain/entities"
)

// La función Decode toma un JSON y lo convierte en una estructura tipo Satellites.
func Decode(body string, requestId string) (*entities.Satellites, error) {
	satellites := &entities.Satellites{}

	err := json.Unmarshal([]byte(body), &satellites)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error marshaling API Gateway request: %v]", body, err)
		return nil, err
	}

	// Devuelve los satellites.
	return satellites, nil
}
