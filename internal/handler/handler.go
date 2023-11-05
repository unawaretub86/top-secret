package handler

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/unawaretub86/top-secret/internal/domain/entities"
	"github.com/unawaretub86/top-secret/internal/domain/ports"
)

type Handler interface {
	HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)
}

type handler struct {
	topSecretPort ports.TopSecretPort
}

func NewHandler(service ports.TopSecretPort) *handler {
	return &handler{
		topSecretPort: service,
	}
}

func (handler *handler) TopSecretHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	lc, _ := lambdacontext.FromContext(ctx)

	requestID := lc.AwsRequestID

	body := request.Body

	bodyResponse, err := handler.topSecretPort.GetLocationAndMessage(body, requestID)
	if err != nil {
		return handler.handleError(err)
	}

	responseBody := handler.formatResponseBody(bodyResponse)

	return handler.createResponse(200, responseBody), nil
}

func (handler *handler) handleError(err error) (*events.APIGatewayProxyResponse, error) {
	return handler.createResponse(404, fmt.Sprintf("Error: %s", err.Error())), nil
}

func (handler *handler) createResponse(statusCode int, body string) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
}

func (handler *handler) formatResponseBody(response *entities.LocationMessage) string {
	return fmt.Sprintf("{\"X\": \"%v\", \"Y\": \"%v\", \"message\": \"%s\"}", response.X, response.Y, *response.Message)
}
