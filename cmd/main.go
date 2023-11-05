package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/unawaretub86/top-secret/internal/handler"
)

func main() {
	lambda.Start(handler.Handler.HandleRequest)
}
