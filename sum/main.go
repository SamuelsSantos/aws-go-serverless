package main

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
type Response events.APIGatewayProxyResponse

// handler is our lambda handler invoked by the `lambda.Start` function call
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {

	var a int
	var b int
	var err error

	if a, err = getValueInt(request.QueryStringParameters["a"]); err != nil {
		return responseError("Parâmetro 'b' invalid.")
	}

	if b, err = getValueInt(request.QueryStringParameters["b"]); err != nil {
		return responseError("Parameter 'b' invalid.")
	}

	return responseOk(a + b)
}

func getValueInt(value string) (int, error) {
	return strconv.Atoi(value)
}

func responseError(msg string) (Response, error) {

	var buf bytes.Buffer

	body, _ := json.Marshal(map[string]interface{}{
		"message": msg,
	})

	json.HTMLEscape(&buf, body)

	return Response{
		StatusCode:      500,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil

}

func responseOk(result int) (Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"resultado": result,
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}

	json.HTMLEscape(&buf, body)

	return Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
