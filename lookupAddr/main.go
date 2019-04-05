package main

import (
	"bytes"
	"encoding/json"
	"net"
	"strings"

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
	hosts := make([]string, 0)

	addr := request.PathParameters["addr"]

	names, err := net.LookupAddr(addr)
	if err != nil {
		hosts = append(hosts, "")
	}

	for _, name := range names {
		// Trim last dot of record
		hosts = append(hosts, strings.TrimRight(name, "."))
	}

	body, err := json.Marshal(map[string]interface{}{
		"addr":  addr,
		"hosts": hosts,
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
			"X-MyCompany-Func-Reply":           "lookupAddr-handler",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
