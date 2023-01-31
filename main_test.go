package main_test

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	main "github.com/billzajac/tao-tetragram"
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request events.APIGatewayProxyRequest
		expect	string
		err			error
	} {
		{
			// Test name has value
			request: events.APIGatewayProxyRequest{QueryStringParameters: map[string]string{"name": "Billy"}},
			expect:  " ",
			err:		 nil,
		},
		{
			// Test name is null
			request: events.APIGatewayProxyRequest{},
			expect:  " ",
			err:		 nil,
		},
	}

	for i, test := range tests {
		response, err := main.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Contains(t, test.expect, response.Body)
		log.Printf("Test %d: %s", i, response.Body)
	}
}
