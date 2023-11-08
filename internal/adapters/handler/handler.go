package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/unawaretub86/top-secret/internal/domain/services"
	"github.com/unawaretub86/top-secret/internal/domain/usecase"
)

func TopSecretHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	triangulationPort := usecase.NewTriangulationUseCase()
	messagePort := usecase.NewMessageUseCase()

	topService := services.NewTopSecretService(triangulationPort, messagePort)

	lc, _ := lambdacontext.FromContext(ctx)
	requestID := lc.AwsRequestID
	body := request.Body

	// llamamos el port de topsecret para acceder al service
	bodyResponse, err := topService.GetLocationAndMessage(body, requestID)
	if err != nil {
		return handleError(err)
	}

	responseBody, err := json.Marshal(bodyResponse)
	if err != nil {
		return handleError(err)
	}

	return createResponse(200, string(responseBody)), nil
}

// handleError maneja los errores y crea una respuesta de error
func handleError(err error) (*events.APIGatewayProxyResponse, error) {
	return createResponse(404, fmt.Sprintf("Error: %s", err.Error())), nil
}

// createResponse crea una respuesta con el código de estado y el cuerpo proporcionados
func createResponse(statusCode int, body string) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
}
