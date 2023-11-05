package usecase

import (
	"strings"

	"github.com/unawaretub86/top-secret/internal/domain/request"
)

type MessageUseCase struct{}

func NewMessageUseCase() *MessageUseCase {
	return &MessageUseCase{}
}

// GetMessage es una función que recupera y procesa un mensaje a partir de mensajes recibidos desde diferentes satélites.
// Devuelve un mensaje combinado y procesado como una cadena de texto.
func GetMessage(requestId string, messages ...[]string) (msg string) {
	var kenobi, skywalker, sato, fixMsg []string

	// Itera a través de las primeras 5 palabras de los mensajes de los satélites.
	for i := 0; i != 5; i++ {
		// Comprueba si la palabra del satélite Kenobi no es vacio.
		if kenobi[i] != "" {
			// Verifica si la palabra ya se encuentra en el mensaje procesado y, si no, la agrega.
			if !request.Contains(fixMsg, kenobi[i]) {
				fixMsg = append(fixMsg, kenobi[i])
			}
		}
		if skywalker[i] != "" {
			if !request.Contains(fixMsg, skywalker[i]) {
				fixMsg = append(fixMsg, skywalker[i])
			}
		}
		if sato[i] != "" {
			if !request.Contains(fixMsg, sato[i]) {
				fixMsg = append(fixMsg, sato[i])
			}
		}
	}

	// Combina las palabras procesadas en un solo mensaje separado por espacios y devuelve el resultado.
	return strings.Join(fixMsg, " ")
}
