package main

import (
	"bytes"
	"encoding/json"
	"net"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer
	nameServers := make([]string, 0)

	domain := request.PathParameters["domain"]

	servers, err := net.LookupNS(domain)
	if err != nil {
		nameServers = append(nameServers, "")
	}

	for _, server := range servers {
		nameServers = append(nameServers, server.Host)
	}

	body, err := json.Marshal(map[string]interface{}{
		"domain":      domain,
		"nameServers": nameServers,
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"X-MyCompany-Func-Reply":           "lookupNS-handler",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
