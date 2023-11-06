package usecase

import (
	"fmt"
	"strings"

	"github.com/unawaretub86/top-secret/internal/config/errors"
)

type MessageUseCase struct{}

func NewMessageUseCase() *MessageUseCase {
	return &MessageUseCase{}
}

// GetMessage toma una solicitud y una matriz de mensajes y devuelve el mensaje concatenado.
func (useCase *MessageUseCase) GetMessage(requestID string, messages ...[]string) (string, error) {
	if len(messages) < 3 {
		fmt.Printf("[RequestId: %s][%v]", requestID, errors.ErrInvalidSatellites)
		return "", errors.ErrInvalidSatellites
	}

	// Inicializa el mensaje resultante con el primer mensaje.
	message := messages[0]

	for _, m := range messages[1:] {
		message = useCase.combineMessages(message, m)
	}

	response := useCase.cleanMessage(message)

	return response, nil
}

// combineMessages combina dos mensajes en una sola matriz.
func (useCase *MessageUseCase) combineMessages(message1, message2 []string) (message []string) {
	maxLen := len(message1)

	// Encuentra la longitud máxima entre los dos mensajes.
	if len(message2) > len(message1) {
		maxLen = len(message2)
	}

	// Combina las partes no vacías de messageA y messageB en un nuevo mensaje.
	for i := 0; i < maxLen; i++ {
		if i < len(message1) && message1[i] != "" {
			message = append(message, message1[i])
		} else if i < len(message2) && message2[i] != "" {
			message = append(message, message2[i])
		} else {
			message = append(message, "")
		}
	}

	return
}

// cleanMessage limpia el mensaje y lo convierte en un string.
func (useCase *MessageUseCase) cleanMessage(message []string) string {
	msg := strings.Builder{}
	for idx, v := range message {
		if v == "" {
			continue
		}

		// Verifica si el elemento actual es diferente al siguiente para evitar duplicados.
		if idx < len(message)-1 && message[idx] != message[idx+1] {
			msg.WriteString(v)
			msg.WriteString(" ")
		}
	}

	// Agrega el último elemento si no es vacío.
	if message[len(message)-1] != "" {
		msg.WriteString(message[len(message)-1])
	}

	return msg.String()
}
