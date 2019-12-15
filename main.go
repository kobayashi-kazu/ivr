package main

import (
	"encoding/xml"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ret := ""
	switch request.Path {
	case "/step-first":
		ret = firstStep(request)
	case "/step-second":
		ret = secondStep(request)
	}
	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/xml",
		},
		StatusCode: http.StatusOK,
		Body:    xml.Header + ret,
	}, nil
}

func main() {
	lambda.Start(handler)
}
