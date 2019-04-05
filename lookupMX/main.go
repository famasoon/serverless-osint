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
	records := make([]string, 0)

	domain := request.PathParameters["domain"]

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		records = append(records, "")
	}

	for _, record := range mxRecords {
		// Trim last dot of record
		records = append(records, strings.TrimRight(record.Host, "."))
	}

	body, err := json.Marshal(map[string]interface{}{
		"domain":    domain,
		"mxrecords": records,
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
			"X-MyCompany-Func-Reply":           "lookupMX-handler",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
